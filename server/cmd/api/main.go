package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/nathanmintegui/edina/internal/database"
	"github.com/nathanmintegui/edina/internal/http"
	"github.com/nathanmintegui/edina/internal/user"

	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	connectionString := os.Getenv("POSTGRES_URI")
	conn, err := database.NewConnection(ctx, connectionString)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	r := gin.Default()
	http.SetRoutes(r)

	user.Configure()
	user.SetRoutes(r)

	r.Run()
}
