package utils

import (
    "crypto/rand"
    "encoding/base64"
    "fmt"
    "os"
    "github.com/joho/godotenv"
)

func GenerateKey() (string, error) {
    key := make([]byte, 16)
    _, err := rand.Read(key)
    if err != nil {
        return "", err
    }
    encodedKey := base64.StdEncoding.EncodeToString(key)

    if err := storeKeyInEnv(encodedKey); err != nil {
        return "", err
    }

    return encodedKey, nil
}

func storeKeyInEnv(key string) error {
    if err := godotenv.Load(); err != nil {
        return err
    }
    file, err := os.OpenFile(".env", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
    if err != nil {
        return err
    }
    defer file.Close()
    if _, err := file.WriteString(fmt.Sprintf("CRYPTO_KEY=%s\n", key)); err != nil {
        return err
    }

    return nil
}

