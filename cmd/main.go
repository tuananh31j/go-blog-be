package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"nta-blog/config"
	"nta-blog/db"
	"nta-blog/libs/appctx"
	"nta-blog/libs/logger"
	"nta-blog/middleware"
	"nta-blog/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	errsChan := make(chan error, 1)
	app := fiber.New()

	logger := logger.NewLogger()

	redisDB := db.ConnectRedis(config.Env.RedisHost, config.Env.RedisPort, config.Env.RedisPass)

	mongoClient, err := db.ConnectMongo(config.Env.MongoURI)
	mongoDB := mongoClient.Database("blog")
	db.SetupUserCollection(mongoDB)
	db.SetupBlogCollection(mongoDB)
	db.SetupTagCollection(mongoDB)

	appContext := appctx.NewAppContext(mongoDB, redisDB, logger)

	app.Use(middleware.Recover(appContext))

	routes.InitRoutes(app, appContext)
	app.Use(middleware.NotFound)

	if err != nil {
		panic(err)
	}

	address := fmt.Sprintf("%s:%s", config.Env.AppHost, config.Env.AppPort)

	defer db.DisconnectMongo(mongoClient)
	defer db.DisconnectRedis(redisDB)
	defer cancel()

	startServer(app, address, errsChan)
	handleShutdown(ctx, app, errsChan)
}

func startServer(app *fiber.App, address string, errsChan chan<- error) {
	if err := app.Listen(address); err != nil {
		errsChan <- fmt.Errorf("Something went wrong when starting the server! %w", err)
	} else {
		logger.ZeroLog.Info().Msg(fmt.Sprintf("Server is running on: %v", address))
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
