package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	pb "github.com/oliveiramarcio/dogphotoapp/proto/generated"
	"github.com/stretchr/testify/assert"
)

func TestGetDogPhoto(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := &pb.DogPhotoResponse{
			Status:  "success",
			Message: "http://localhost/image.jpg",
		}

		responseBytes, _ := json.Marshal(response)
		w.WriteHeader(http.StatusOK)
		w.Write(responseBytes)
	}))
	defer mockServer.Close()

	server := &Server{
		baseUrl: mockServer.URL,
		cache: &InMemoryCache{
			duration: time.Millisecond * 3,
		},
	}

	req := &pb.DogPhotoRequest{
		Breed: "bulldog",
	}

	// First hit (should have no cache)
	res, err := server.GetDogPhoto(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, "success", res.Status)
	assert.Equal(t, "http://localhost/image.jpg", res.Message)

	// Second hit (should have cache)
	res, err = server.GetDogPhoto(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, "success (cached)", res.Status)
	assert.Equal(t, "http://localhost/image.jpg", res.Message)

	// Third hit after waiting for the cache to expire should not have the result cached
	time.Sleep(3 * time.Millisecond)
	res, err = server.GetDogPhoto(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, "success", res.Status)
	assert.Equal(t, "http://localhost/image.jpg", res.Message)
}
