package services

import (
    "time"
    "context"
    "passguardian-api/config"
    "passguardian-api/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveCredential(cred models.Credential) error {
    cred.CreatedAt = time.Now()
    _, err := config.DB.Collection("credentials").InsertOne(context.Background(), cred)
    return err
}

func DeleteCredential(db *mongo.Database, ID primitive.ObjectID) error {
    filter := bson.M{"_id": ID}
    _, err := db.Collection("credentials").DeleteOne(context.Background(), filter)
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

func GetCredentialByID(db *mongo.Database, ID primitive.ObjectID) (*models.Credential, error) {
    var credential models.Credential

    filter := bson.M{"_id": ID}
    err := db.Collection("credentials").FindOne(context.Background(), filter).Decode(&credential)

    if err != nil {
        return nil, err
    }

    return &credential, nil
}