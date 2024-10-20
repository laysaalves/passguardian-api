package controllers

import (
    "passguardian-api/models"
    "passguardian-api/services"
    "github.com/gin-gonic/gin"
    "net/http"
)

func SaveCredentials(c *gin.Context) {
    var cred models.Credential

    if err := c.BindJSON(&cred); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := services.SaveCredential(cred)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Save credentials error"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Credentials saved successfully!"})
}

func GetCredentials(c *gin.Context) {
    credentials, err := services.GetAllCredentials()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Get credentials error"})
        return
    }

    c.JSON(http.StatusOK, credentials)
}
