package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client

func init() {
	// Connect to Redis.
	// IMPORTANT: Use the container name 'my-full-stack-redis' as the hostname!
	rdb = redis.NewClient(&redis.Options{
		Addr:     "my-full-stack-redis:6379", // Redis default port is 6379
		Password: "",                         // No password by default
		DB:       0,                          // Default DB
	})

	// Ping Redis to ensure connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	log.Println("Successfully connected to Redis!")
}

// corsMiddleware adds CORS headers
func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*") // Allow any origin for simplicity in demo
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	// Increment a counter in Redis
	val, err := rdb.Incr(ctx, "page_views").Result()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error incrementing counter: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with JSON
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"count": %d}`, val)
}

func main() {
	// Apply CORS middleware to our handler
	http.Handle("/api/count", corsMiddleware(http.HandlerFunc(countHandler)))
	
	port := "8080"
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}