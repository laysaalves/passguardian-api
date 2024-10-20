package config

import (
    "context"
    "log"
    "time"
    "os"
    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectToMongo() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal(".env error:", err)
    }

    mongoURI := os.Getenv("MONGO_URI")
    clientOptions := options.Client().ApplyURI(mongoURI)

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal("MongoDB connection error:", err)
    }

    if err := client.Ping(ctx, nil); err != nil {
        log.Fatal("MongoDB ping error:", err)
    }

    log.Println("MongoDB connected successfully!")
    DB = client.Database("passguardianDB")
}
