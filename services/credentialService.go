package services

import (
    "context"
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "errors"
    "time"
    "io"
    "passguardian-api/config"
    "passguardian-api/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func encrypt(plainText, key []byte) (string, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }

    cipherText := make([]byte, aes.BlockSize+len(plainText))
    iv := cipherText[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return "", err
    }

    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

    return base64.StdEncoding.EncodeToString(cipherText), nil
}

func decrypt(cipherText string, key []byte) (string, error) {
    cipherTextBytes, _ := base64.StdEncoding.DecodeString(cipherText)

    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }

    if len(cipherTextBytes) < aes.BlockSize {
        return "", errors.New("ciphertext too short")
    }
    iv := cipherTextBytes[:aes.BlockSize]
    cipherTextBytes = cipherTextBytes[aes.BlockSize:]

    stream := cipher.NewCFBDecrypter(block, iv)
    stream.XORKeyStream(cipherTextBytes, cipherTextBytes)

    return string(cipherTextBytes), nil
}

var key = []byte("your-16-byte-key")

func SaveCredential(cred models.Credential) error {
    cred.CreatedAt = time.Now()

    encryptedPassword, err := encrypt([]byte(cred.Password), key)
    if err != nil {
        return err
    }
    cred.Password = encryptedPassword

    _, err = config.DB.Collection("credentials").InsertOne(context.Background(), cred)
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

        decryptedPassword, err := decrypt(cred.Password, key)
        if err == nil {
            cred.Password = decryptedPassword
        }

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

    decryptedPassword, err := decrypt(credential.Password, key)
    if err == nil {
        credential.Password = decryptedPassword
    }

    return &credential, nil
}
