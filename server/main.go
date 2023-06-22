package main

import (
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/joho/godotenv"
	pb "github.com/oliveiramarcio/dogphotoapp/proto/generated"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	serverAddress := os.Getenv("SERVER_ADDRESS")
	if strings.TrimSpace(serverAddress) == "" {
		log.Fatal("server address not set")
	}

	dogApiBaseUrl := os.Getenv("DOG_API_BASE_URL")
	if strings.TrimSpace(dogApiBaseUrl) == "" {
		log.Fatal("dog api base url not set")
	}

	secondsToExpire := os.Getenv("CACHE_EXPIRATION_SECONDS")
	if strings.TrimSpace(secondsToExpire) == "" {
		log.Fatal("cache expiration seconds not set")
	}

	cacheExpirationSeconds, err := strconv.Atoi(secondsToExpire)
	if err != nil {
		log.Fatalf("failed to parse cache expiration seconds: %v", err)
	}

	lis, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterDogPhotoServiceServer(s, &Server{
		baseUrl: dogApiBaseUrl,
		cache: &InMemoryCache{
			duration: time.Second * time.Duration(cacheExpirationSeconds),
		},
	})
	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
