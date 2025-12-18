package main

import (
	"fmt"
	"log"
	"net/http"

	"guvi-project/db"
	"guvi-project/handlers"
	"guvi-project/middleware"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default environment variables")
	}

	// 1. Initialize Database Connections
	db.Init()

	// 2. Create Router
	mux := http.NewServeMux()

	// 3. Define Routes
	mux.HandleFunc("/api/register", handlers.Register)
	mux.HandleFunc("/api/login", handlers.Login)
	
	// Protected Profile Route using Middleware
	mux.HandleFunc("/api/profile", middleware.Auth(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.GetProfile(w, r)
		} else if r.Method == http.MethodPost {
			handlers.UpdateProfile(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))

	// 4. Setup CORS
	// Allow frontend origin (default Vite port 5173)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173", "http://127.0.0.1:5173"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	handler := c.Handler(mux)

	// 5. Start Server
	fmt.Println("Go Backend Server running on port 8080...")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
