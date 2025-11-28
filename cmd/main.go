package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/partickle/avito-pr-review-service/internal/handler/common"
	flag "github.com/spf13/pflag"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}

	dbPool := mustInitDB()
	startServer(dbPool, resolvePort())

}

func mustInitDB() *pgxpool.Pool {
	var dbPool *pgxpool.Pool

	config, err := pgxpool.ParseConfig(getConnectionString())
	if err != nil {
		log.Fatalf("Unable to parse connection string: %v\n", err)
	}

	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = time.Minute * 30

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbPool, err = pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}

	pingAttemptsLimit := 3
	var pingErr error

	for i := 1; i <= pingAttemptsLimit; i++ {
		pingCtx, pingCancel := context.WithTimeout(context.Background(), 5*time.Second)
		pingErr = dbPool.Ping(pingCtx)
		pingCancel()
		if pingErr == nil {
			break
		}
		log.Printf("dp ping attempt %d failed: %v", i, pingErr)
		if i < pingAttemptsLimit {
			time.Sleep(500 * time.Millisecond)
		}
	}

	err = dbPool.Ping(ctx)
	if pingErr != nil {
		log.Fatalf("Unable to ping database")
	}

	log.Println("Database connection pool established")
	return dbPool
}

func getConnectionString() string {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, dbname)
}

func startServer(
	dbPool *pgxpool.Pool,
	port string,
) {
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: initRouter(),
	}

	log.Printf("Server started on %s\n", srv.Addr)
	serverErr := make(chan error, 1)
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
		}
	}()

	waitGracefulShutdown(srv, dbPool, serverErr)

	log.Println("Shutting down service-pr-review")
}

func resolvePort() string {
	port := os.Getenv("PORT")

	var portFlag = flag.String("port", "", "укажите порт")
	flag.Parse()

	if portFlag != nil && *portFlag != "" {
		port = *portFlag
	}

	if port == "" {
		port = "8080"
	}

	return port
}

func initRouter(courier *handlers.CourierController) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/ping", common.Ping)
	r.Head("/healthcheck", common.HealthCheck)

	r.Route("/team", func(r chi.Router) {
		r.Post("/add", courier.Get)
		r.Get("/get", courier.Create)
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/setIsActive", courier.Get)
		r.Get("/getReview", courier.Create)
	})

	r.Route("/pullRequest", func(r chi.Router) {
		r.Post("/create", courier.Get)
		r.Post("/merge", courier.Create)
		r.Post("/reassign", courier.Create)
	})

	return r
}

func waitGracefulShutdown(srv *http.Server, dbPool *pgxpool.Pool, serverErr <-chan error) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var reason string
	select {
	case <-ctx.Done():
		reason = ctx.Err().Error()
	case err := <-serverErr:
		reason = "server error: " + err.Error()
	}
	log.Printf("Shutdown initiared (%s)", reason)

	log.Println("Shutdown down HTTP server...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("HTTP server shutdown failed: %v\n", err)
	} else {
		log.Println("HTTP server stopped")
	}

	log.Println("Closing DB connection...")
	dbPool.Close()
	log.Println("DB pool closed")
}
