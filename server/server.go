package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	pb "github.com/oliveiramarcio/dogphotoapp/proto/generated"
)

type Server struct {
	pb.UnimplementedDogPhotoServiceServer
	baseUrl string
	cache   CacheRepository
}

func (s *Server) GetDogPhoto(ctx context.Context, req *pb.DogPhotoRequest) (*pb.DogPhotoResponse, error) {
	breed := req.GetBreed()

	if cachedUrl, ok := s.cache.Get(breed); ok {
		log.Printf("retrieved dog photo for breed '%s' from cache", breed)
		return &pb.DogPhotoResponse{
			Status:  "success (cached)",
			Message: cachedUrl,
		}, nil
	}

	url := fmt.Sprintf("%s/api/breed/%s/images/random", s.baseUrl, breed)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make a request to the dog api: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request to the dog api failed with status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read the response body: %v", err)
	}

	var result pb.DogPhotoResponse
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("can not unmarshal json")
	}

	if strings.EqualFold(result.Status, "success") {
		log.Printf("retrieved dog photo for breed '%s' from %s", breed, url)
		s.cache.Set(breed, result.Message)
	}

	return &result, nil
}
