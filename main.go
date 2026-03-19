package main

import (
	"context"
	"fmt"
	"interview/order/api"
	eventbus "interview/order/event-bus"
	"interview/order/logger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
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

	http.HandleFunc("/order", logger.Logger(api.HandleOrder))

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
