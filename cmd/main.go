package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"nta-blog/config"
	"nta-blog/db"

	"github.com/gofiber/fiber/v2"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	errsChan := make(chan error, 1)
	app := fiber.New()
	mongoDB := db.ConnectMongo(config.MongoURI)
	redisDB := db.ConnectRedis(config.RedisHost, config.RedisPort, config.RedisPass)

	address := fmt.Sprintf("%s:%s", config.AppHost, config.AppPort)

	defer db.DisconnectMongo(mongoDB)
	defer db.DisconnectRedis(redisDB)
	defer cancel()

	startServer(app, address, errsChan)
	handleShutdown(ctx, app, errsChan)
}

func startServer(app *fiber.App, address string, errsChan chan<- error) {
	if err := app.Listen(address); err != nil {
		errsChan <- fmt.Errorf("Something went wrong when starting the server! %w", err)
	}
}

func handleShutdown(ctx context.Context, app *fiber.App, errsChan <-chan error) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errsChan:
		fmt.Printf("Server error: %v", err)

	// lắng nghe tín hiệu tắt server từ người dùng như ctrl + C
	case <-quit:
		fmt.Printf("Shutting down server...")
		if err := app.Shutdown(); err != nil {
			fmt.Printf("Somethings wrong when you shutting down server!(%v)", err)
		}
	case <-errsChan:
		fmt.Printf("Server is error!")
	case <-ctx.Done():
		fmt.Printf("Server exiting by context cancellation!")
	}
}
