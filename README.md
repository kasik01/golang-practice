# Todo App API (Golang, Gin, GORM, JWT, PostgreSQL)

## Introduction

This is a RESTful API project for a Todo management application using:
- **Golang** with the **Gin** framework
- **GORM** as the ORM
- **PostgreSQL** as the database
- **JWT** for user authentication
- **Docker** and **docker-compose** for easy deployment

## Features

- User registration and login with JWT authentication
- CRUD operations for tasks (create, read, update, delete)
- Get tasks by user
- Export tasks to Excel file
- Protect routes with JWT authentication middleware

## Project Structure

```
.
├── Dockerfile
├── docker-compose.yml
├── main.go
├── go.mod
├── pkg
│   ├── config        // Database connection
│   ├── controllers   // Request handlers
│   ├── middleware    // JWT middleware
│   ├── models        // Model definitions and DB queries
│   ├── routes        // Route definitions
│   └── utils         // Utilities (password hashing, JWT, ...)
└── ...
```

## Running with Docker

1. **Clone the project and navigate to the directory**
2. **Run:**
   ```sh
   docker-compose up --build
   ```
   - PostgreSQL will run on port `5432`
   - The API will run on port `8080`

3. **Check environment variables in `docker-compose.yml`** (change user, password, secret if needed)

## API Endpoints

### Auth

- `POST /signup`  
  Register a new user  
  ```json
  {
    "username": "yourname",
    "password": "yourpassword"
  }
  ```

- `POST /signin`  
  Login, returns JWT  
  ```json
  {
    "username": "yourname",
    "password": "yourpassword"
  }
  ```

### Task (requires Authorization: Bearer <token> header)

- `POST /tasks`  
  Create a new task  
  ```json
  {
    "title": "Task 1",
    "description": "Details",
    "due_date": "2024-06-01T12:00:00Z"
  }
  ```

- `GET /tasks`  
  Get user's tasks

- `PUT /tasks/:id`  
  Update a task

- `DELETE /tasks/:id`  
  Delete a task

- `GET /tasks/export`  
  Export tasks to Excel file

## Notes

- All `/tasks` routes require a valid JWT.
- To debug in VS Code, use the configuration in `.vscode/launch.json`.
