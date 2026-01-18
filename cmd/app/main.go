package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/esmrzv/go_shop/internal/auth"
	"github.com/esmrzv/go_shop/internal/config"
	"github.com/esmrzv/go_shop/internal/db"
	"github.com/esmrzv/go_shop/internal/http/handler"
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
	authService := service.NewAuthService(userRepo, cfg.JWTSecret)
	authHandler := handler.NewAuthHandler(authService)
	productRepo := repository.NewProductPostgres(database)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	mux := http.NewServeMux()
	mux.HandleFunc("/register", authHandler.Register)
	mux.HandleFunc("/login", authHandler.Login)
	mux.Handle("/products", auth.JWTMiddleware(cfg.JWTSecret, http.HandlerFunc(productHandler.Create)))
	mux.Handle("/products/list", auth.JWTMiddleware(cfg.JWTSecret,
		http.HandlerFunc(productHandler.List),
	))
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
