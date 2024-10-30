# Golang API for credentials storage: "passguardian"

![gopher](misc/gopher-dance.gif)

## Tech Stack
- Go
- Gin
- MongoDB
- Postman

## Introduction

The purpose of this API is to learn how synchronous cryptography works in practice, taking into account the Information Security and Introduction to Computer Programming classes at my Computer Science university, which cover these topics using the C language, which is similar to Golang.

## API Endpoints
| route               | description                                          
|----------------------|-----------------------------------------------------
| <kbd>GET /login-credentials</kbd>     | retrieves user credentials see [response details](#get-credentials-detail)
| <kbd>POST /login-credentials</kbd>     | create user credentials see [request details](#post-credentials-detail)
| <kbd>GET /credentials/:id</kbd> | retrieves credential by id see [response details](#get-credential-by-id)
| <kbd>DELETE /credentials/:id</kbd> | deletes credential by id see [response details](#delete-credential-by-id)

<h3 id="get-credentials-detail">GET /login-credentials</h3>

**RESPONSE**
```json
{
  "serviceName": "GitHub",
  "user": "Golango",
  "password": "mypassword432 <- (decrypted)"
}
```
```json
{
  "serviceName": "Instagram",
  "user": "Pythoneiro",
  "password": "senhona426 <- (decrypted)"
}
```

<h3 id="post-credentials-detail">POST /login-credentials</h3>

**REQUEST**
```json
{
  "serviceName": "Twitch",
  "user": "Java Scripto",
  "password": "mypassword123 <- (encrypted)"
}
```

**RESPONSE**
```json
{
  "message": "Credential saved successfully!"
}
```

<h3 id="get-credential-by-id">GET /credentials/:id</h3>

**REQUEST**
```json
{
  "id": "6716baed4f1c6829e2ab9cbe"
}
```

**RESPONSE**
```json
{
  "serviceName": "Instagram",
  "user": "Pythoneiro",
  "password": "senhona426 <- (decrypted)"
}
```

<h3 id="delete-credential-by-id">DELETE /credentials/:id</h3>

**REQUEST**
```json
{
  "id": "6716baed4f1c6829e2ab9cbe"
}
```

**RESPONSE**
```json
{
  "message": "Credential deleted successfully!"
}