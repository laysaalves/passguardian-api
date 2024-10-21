package routes

import (
    "passguardian-api/controllers"
    "github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
    router.POST("/app/account/login-credentials", controllers.SaveCredentials)
    router.GET("/app/account/login-credentials", controllers.GetCredentials)
    router.GET("/credentials/:id", controllers.GetCredentialByID)
    router.DELETE("/credentials/:id", controllers.DeleteCredentials)
}
