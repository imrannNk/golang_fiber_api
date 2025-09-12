# Go User API

A simple REST API for user management built with Go and Fiber.

## Setup

```bash
go mod tidy
go run web.go
```

## API Endpoints

- `GET /users` - Get all users
- `GET /users/:userId` - Get user by ID- `GET /users/search` - Search users by name or email
- `GET /users/search` - Search users by name or email
- `POST /users` - Create new user
- `PUT /users/:userId` - Update user
- `DELETE /users/:userId` - Delete user

Server runs on port 8080.