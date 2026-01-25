package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/esmrzv/go_shop/internal/cartworker"
	"github.com/esmrzv/go_shop/internal/config"
	"github.com/esmrzv/go_shop/internal/db"
	"github.com/esmrzv/go_shop/internal/http/handler"
	"github.com/esmrzv/go_shop/internal/http/routers"
	"github.com/esmrzv/go_shop/internal/repository"
	"github.com/esmrzv/go_shop/internal/service"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки .env: %v", err)
	}

	cfg := config.Load()

	database, err := db.NewPostgres(db.Config{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		User:     cfg.DB.User,
		Password: cfg.DB.Password,
		Name:     cfg.DB.Name,
		SSLMode:  cfg.DB.SSLMode,
	})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	defer database.Close()

	userRepo := repository.NewUserPostgres(database)
	productRepo := repository.NewProductPostgres(database)
	categoryRepo := repository.NewCategoryRepo(database)

	authService := service.NewAuthService(userRepo, cfg.JWTSecret)
	productService := service.NewProductService(productRepo)
	categoryService := service.NewCategoryService(categoryRepo)

	authHandler := handler.NewAuthHandler(authService)
	productHandler := handler.NewProductHandler(productService)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	cartWorker := cartworker.NewWorker()
	go cartWorker.Run()

	cartService := service.NewAsyncCartService(cartWorker)
	cartHandler := handler.NewCartHandler(cartService)

	mux := http.NewServeMux()

	routers.RegisterAuthRouter(mux, authHandler)
	routers.RegisterCategoryRouter(mux, categoryHandler, cfg.JWTSecret)
	routers.RegisterProductRouter(mux, productHandler, cfg.JWTSecret)
	routers.RegisterCartRouter(mux, cartHandler, cfg.JWTSecret)

	srv := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: mux,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		log.Printf("server listening on port %s", cfg.AppPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("could not listen on %s: %v\n", cfg.AppPort, err)
		}
	}()

	<-ctx.Done()
	log.Println("shutting down server...")

	shutdonwCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdonwCtx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}
	log.Println("http server shotdown seccessfully")

}
