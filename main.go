package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/amanraj8848/rssScrapper/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}
func main() {
	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in .env")
	}
	DB_URL := os.Getenv("DB_URL")
	if DB_URL == "" {
		log.Fatal("DB_URL is not found in .env")
	}
	db, err := sql.Open("postgres", DB_URL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")
		apiCfg := apiConfig{
		DB: database.New(db), // Corrected assignment to use *sql.DB
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	router.Mount("/v1", v1Router)

	v1Router.Get("/check", handlerReadiness) // Corrected typo in handler name
	v1Router.Get("/err", handlerErr)        // Corrected typo in handler name
	v1Router.Post("/user", apiCfg.handlerCreateUser)
	v1Router.Get("/user", apiCfg.middlewareAuth(apiCfg.handlerGetUser))


	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)
	err2 := srv.ListenAndServe()
	if err2 != nil {
		log.Fatal(err2)
	}
}

