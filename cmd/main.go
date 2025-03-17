package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"nta-blog/internal/infrastructure/config"
	"nta-blog/internal/infrastructure/db"
	"nta-blog/internal/lib/appctx"
	"nta-blog/internal/lib/logger"
	"nta-blog/internal/middleware"
	routes "nta-blog/internal/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	errsChan := make(chan error, 1)
	app := fiber.New(config.FiberConfig(config.Env.AppENV))
	app.Use(cors.New(cors.Config{
		AllowOrigins:     config.Env.AllowOrigin,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS", // Đảm bảo có OPTIONS
		AllowHeaders:     "Content-Type, Authorization",            // Liệt kê rõ ràng các header
		AllowCredentials: true,                                     // Hỗ trợ cookie/token
		ExposeHeaders:    "Content-Length",                         // Nếu cần
	}))
	logger := logger.NewLogger()

	if config.Env.AppENV == "development" {
		app.Use(middleware.LoggerConfigFiber())
	}

	redisDB := db.ConnectRedis(config.Env.RedisURL)

	mongoClient, err := db.ConnectMongo(config.Env.MongoURI)
	if err != nil {
		logger.Debug().Msg(err.Error())
		panic(err)
	}
	mongoDB := mongoClient.Database("blog")
	db.SetupUserCollection(mongoDB)
	db.SetupBlogCollection(mongoDB)
	db.SetupTagCollection(mongoDB)
	db.SetupImageCollection(mongoDB)
	db.SetupGuestBookCollection(mongoDB)

	cld, err := db.NewCld(config.Env.CloudinaryName, config.Env.CloudinaryAPIKey, config.Env.CloudinaryAPISecret)
	if err != nil {
		logger.Debug().Msg(err.Error())
	}

	appContext := appctx.NewAppContext(mongoDB, redisDB, cld, logger)

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
