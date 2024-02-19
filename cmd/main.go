package main

import (
	"log"
	"os"

	"totoconfigapi/internal/api"
	"totoconfigapi/internal/db"

	"github.com/gin-gonic/gin"
)

func main() {

	mongoDBURI := os.Getenv("MONGO_URI")
	mongoDBDatabase := os.Getenv("MONGO_DATABASE")
	mongoDBCollection := os.Getenv("MONGO_COLLECTION")

	mongoDB, err := db.NewMongoDB(mongoDBURI, mongoDBDatabase, mongoDBCollection)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer mongoDB.Close()

	router := gin.Default()

	apiHandler := api.NewHandler(mongoDB)

	apiHandler.RegisterRoutes(router)

	if err := router.Run(); err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}
