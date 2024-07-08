package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/hiahir357/rssagg/internal/database"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	// Loading env variables
	godotenv.Load(".env")

	// Getting env varibles
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}
	dbURLString := os.Getenv("DB_URL")
	if dbURLString == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	// Database connection
	conn, err := sql.Open("postgres", dbURLString)
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}
	queries := database.New(conn)
	apiCfg := apiConfig{
		DB: queries,
	}

	// Scraper routine
	go startScraping(queries, 10, time.Minute)

	// Api router
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Api router enpoints v1
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)

	v1Router.Post("/users", apiCfg.handleCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handleGetUser))

	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handleCreateFeed))
	v1Router.Get("/feeds", apiCfg.handleGetFeeds)

	v1Router.Get("/posts", apiCfg.middlewareAuth(apiCfg.handleGetPostsForUser))

	v1Router.Post("/feeds_follows", apiCfg.middlewareAuth(apiCfg.handleCreateFeedFollow))
	v1Router.Get("/feeds_follows", apiCfg.middlewareAuth(apiCfg.handleGetFeedFollows))
	v1Router.Delete("/feeds_follows/{feedFollowId}", apiCfg.middlewareAuth(apiCfg.handleDeleteFeedFollow))

	// Mounting router
	router.Mount("/v1", v1Router)

	// Server
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	log.Printf("Server starting on port %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
