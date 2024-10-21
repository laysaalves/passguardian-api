package controllers

import (
	"net/http"
	"passguardian-api/config"
	"passguardian-api/models"
	"passguardian-api/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func DeleteCredentials(c *gin.Context) {
    idParam := c.Param("id")

    objectID, err := primitive.ObjectIDFromHex(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    err = services.DeleteCredential(config.DB, objectID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete credentials error"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Credentials deleted successfully!"})
}

func GetCredentials(c *gin.Context) {
    credentials, err := services.GetAllCredentials()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Get credentials error"})
        return
    }

    c.JSON(http.StatusOK, credentials)
}

func GetCredentialByID(c *gin.Context) {
    idParam := c.Param("id")

    objectID, err := primitive.ObjectIDFromHex(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    credential, err := services.GetCredentialByID(config.DB, objectID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Credential not found"})
        return
    }
    c.JSON(http.StatusOK, credential)
}