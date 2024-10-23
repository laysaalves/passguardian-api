# Golang API for credentials storage: "passguardian"

![gopher](misc/gopher-dance.gif)

## Tech Stack
- Go
- Gin
- MongoDB
- Docker
- Postman

## Introduction

This is an API with the sole purpose of learning how to use non-relational databases from the Golang language and how to build the security of an application, in this case using fictitious data for encryption.

## API Endpoints
​
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
  "password": "mypassword432"
}
```

<h3 id="post-credentials-detail">POST /login-credentials</h3>

**REQUEST**
```json
{
  "serviceName": "Twitch",
  "user": "Java Scripto",
  "password": "mypassword123"
}
```

**RESPONSE**
```json
{
  message: "Credentials saved successfully!"
}
```