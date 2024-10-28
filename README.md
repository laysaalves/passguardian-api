# Golang API for credentials storage: "passguardian"

![gopher](misc/gopher-dance.gif)

## Tech Stack
- Go
- Gin
- MongoDB
- Docker
- Postman

## Introduction

This is an API with the sole purpose of learning how to use non-relational databases from the Golang language and how to build the security of an application, in this case using dummy data for encryption, as well as serving in a Docker container.

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
  "message": "Credentials saved successfully!"
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
  "password": "senhona426"
}
```

<h3 id="delete-credential-by-id">DELETE /credentials/:id</h3>

**RESPONSE**
```json
{
  "message": "Credentials deleted successfully!"
}