package main

import (
    "passguardian-api/config"
    "passguardian-api/routes"
    "github.com/gin-gonic/gin"
    "passguardian-api/utils"
    "os"
    "log"
)

func main() {
    key := os.Getenv("CRYPTO_KEY")
    if key == "" {
        var err error
        key, err = utils.GenerateKey()
        if err != nil {
            log.Fatal("Failed to generate encryption key:", err)
        }
        log.Println("Generated encryption key:", key)
    }
    config.ConnectToMongo()
    router := gin.Default()
    routes.InitializeRoutes(router)
    router.Run("0.0.0.0:3000")
}
