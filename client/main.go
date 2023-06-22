package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	pb "github.com/oliveiramarcio/dogphotoapp/proto/generated"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	photoDir := os.Getenv("PHOTO_BASE_DIR")
	if strings.TrimSpace(serverAddress) == "" {
		log.Fatal("photo base directory not set")
	}

	breed := flag.String("breed", "", "dog breed")
	flag.Parse()
	if *breed == "" {
		log.Fatal("missing breed argument")
	}

	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to server: %v", err)
	}
	defer conn.Close()

	c := pb.NewDogPhotoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.GetDogPhoto(ctx, &pb.DogPhotoRequest{Breed: *breed})
	if err != nil {
		log.Fatalf("could not retrieve dog photo: %v", err)
	}

	err = downloadAndSavePhoto(fmt.Sprintf("%s/%s", photoDir, *breed), resp.Message)
	if err != nil {
		log.Fatalf("could not save dog photo: %v", err)
	}
}
