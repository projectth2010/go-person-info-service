// config/db.go
package config

import (
    "context"
    "fmt"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func getMongoURI() string {
    host := os.Getenv("MONGODB_HOST")
    if host == "" {
        host = "localhost"
    }
    port := os.Getenv("MONGODB_PORT")
    if port == "" {
        port = "27017"
    }
    return fmt.Sprintf("mongodb://%s:%s", host, port)
}

func ConnectDB() error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Get MongoDB URI from environment or use default
    mongoURI := os.Getenv("MONGODB_URI")
    if mongoURI == "" {
        mongoURI = getMongoURI()
    }

    // Set client options
    clientOptions := options.Client().ApplyURI(mongoURI)

    // Connect to MongoDB
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        return fmt.Errorf("failed to connect to MongoDB: %w", err)
    }

    // Check the connection
    err = client.Ping(ctx, nil)
    if err != nil {
        return fmt.Errorf("failed to ping MongoDB: %w", err)
    }

    // Get database name from environment or use default
    dbName := os.Getenv("MONGODB_NAME")
    if dbName == "" {
        dbName = "person_info_service"
    }

    DB = client.Database(dbName)
    return nil
}