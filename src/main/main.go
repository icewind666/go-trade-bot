package main

import (
	"crypto/tls"
	"github.com/joho/godotenv"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"os"
	"time"
	"trading-bot/src/main/investapi"
	"trading-bot/src/main/models"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	log.Println("Starting up")
	apiToken, exists := os.LookupEnv("API_TOKEN")

	if !exists {
		log.Fatalln("no token")
	}

	apiAddress, exists := os.LookupEnv("API_ADDRESS")

	if !exists {
		log.Fatalln("no address")
	}

	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	cred := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	})

	opts = append(opts, grpc.WithTransportCredentials(cred))
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	conn, err := grpc.Dial(apiAddress, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	// Add token to gRPC Request.
	ctx = metadata.AppendToOutgoingContext(ctx, "Authorization", "Bearer "+apiToken)
	ctx = metadata.AppendToOutgoingContext(ctx, "x-app-name", "golang-sdk-sample")

	// Send the request.
	client := investapi.NewUsersServiceClient(conn)
	request, err := client.GetInfo(ctx, &investapi.GetInfoRequest{})
	userInfo := models.MapFrom(request)

	if err != nil {
		log.Println(err)
	}

	log.Println(userInfo.CanWorkWithForeignShares)
}
