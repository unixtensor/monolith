package main

import (
	"context"
	"log"
	"os"
	"regexp"

	"github.com/redis/go-redis/v9"
	"github.com/unixtensor/monolith/pkg/api"
)

func main() {
	token, token_set := os.LookupEnv("TOKEN")
	port, port_set := os.LookupEnv("PORT")
	_, debug_set := os.LookupEnv("DEBUG")
	if !token_set {
		log.Fatal("Environment variable: TOKEN is not set, STOPPING.")
	}
	if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(token) {
		log.Fatalf("TOKEN must only contain letters and numbers, STOPPING.")
	}
	if !port_set {
		port = "3000"
	}

	bg_ctx := context.Background()
	Redis := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	defer Redis.Close()

	api.V1(bg_ctx, api.Config{
		Port:  port,
		Token: token,
		Debug: debug_set,
		Redis: Redis,
	})
}
