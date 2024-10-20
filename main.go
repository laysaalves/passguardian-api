package main

import (
    "passguardian-api/config"
    "passguardian-api/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    config.ConnectToMongo()
    router := gin.Default()
    routes.InitializeRoutes(router)
    router.Run("localhost:3000")
}
