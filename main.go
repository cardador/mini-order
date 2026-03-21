package main

import (
	"context"
	"fmt"
	"interview/order/api"
	eventbus "interview/order/event-bus"
	"interview/order/logger"
	"interview/order/store"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func main() {
	srv := &http.Server{Addr: ":8080", Handler: nil}

	go eventbus.ProcessOrder()

	go func() {
		fmt.Println("Server starting on :8080...")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	connStr := "postgres://postgres:password@localhost:5432/postgres?sslmode=disable"
	db, err := store.NewPostgresStore(connStr)
	if err != nil {
		log.Fatalf("Problem to initialize DB, %s", err)
	}
	// DynamoDB
	ctx := context.TODO()

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("dummy", "dummy", "")),
	)
	if err != nil {
		panic(err)
	}
	dynamoClient := dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		o.BaseEndpoint = aws.String("http://localhost:8000")
	})
	repo := store.NewDynamoStore(dynamoClient, "Orders")

	http.HandleFunc("/order", logger.Logger(api.HandleOrder(db)))
	http.HandleFunc("/order/get/{id}", logger.Logger(api.GetOrder(db)))

	http.HandleFunc("/dynamo/order", logger.Logger(api.HandleOrder(repo)))
	http.HandleFunc("/dynamo/order/get/{id}", logger.Logger(api.GetOrder(repo)))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}
