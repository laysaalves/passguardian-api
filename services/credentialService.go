package services

import (
    "context"
    "passguardian-api/config"
    "passguardian-api/models"
    "time"
    "go.mongodb.org/mongo-driver/bson"
)

func SaveCredential(cred models.Credential) error {
    cred.CreatedAt = time.Now()
    _, err := config.DB.Collection("credentials").InsertOne(context.Background(), cred)
    return err
}

func GetAllCredentials() ([]models.Credential, error) {
    var credentials []models.Credential
    cursor, err := config.DB.Collection("credentials").Find(context.Background(), bson.M{})
    if err != nil {
        return nil, err
    }

    for cursor.Next(context.Background()) {
        var cred models.Credential
        cursor.Decode(&cred)
        credentials = append(credentials, cred)
    }
    return credentials, nil
}
