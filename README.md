# User Management API (Go + Fiber + SQLC + MySQL + Zap)

This is a simple REST API for managing users built using Go.  
It was created to practice backend fundamentals like clean architecture, database handling, validation, and logging.

The API supports basic CRUD operations along with pagination.

---

## Features

- Create user
- Get user by ID
- Get all users (pagination)
- Update user
- Delete user
- Input validation (name, DOB)
- DOB format + future date validation
- Structured logging using Zap

---

## Tech Stack

- Go
- Fiber
- MySQL
- SQLC
- Zap
- validator/v10
- godotenv

---

## Project Structure

```
/cmd/server/main.go
/config/
/db/migrations/
/db/sqlc/<generated>
/internal/
├── handler/
├── repository/
├── service/
├── routes/
├── middleware/
├── models/
└── logger/
```

---

## Setup & Run Instructions

### 1. Clone the repository

```bash
git clone https://github.com/your-username/user-management.git
cd user-management
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Create environment file

Copy .env.example to .env:

```bash
cp .env.example .env
```

Update DB credentials inside .env:

```bash
APP_PORT=
DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=
```

### 4. Make sure MySQL is running

Create database if it does not exist:

```bash
CREATE DATABASE user_api;
```

### 5. Run the application

```bash
go run cmd/server/main.go
```

### 6. Server will start at

```bash
http://localhost:3000
```

## API Demonstration

Below are example requests and responses for all available endpoints.

---

### Create User

**Request**

```http
POST /users
Content-Type: application/json
{
  "name": "John Doe",
  "dob": "2000-01-01"
}
```

**Response**

```http
{
  "id": 1,
  "name": "John Doe",
  "dob": "2000-01-01",
  "age": 26
}
```

### Get User by ID

**Request**

```http
GET /users/:id
```

**Response**

```http
{
  "id": 1,
  "name": "John Doe",
  "dob": "2000-01-01",
  "age": 26
}
```

### Get All Users (Pagination)

**Request**

```http
GET /users?page=1&limit=10
```

**Response**

```http
{
  "count": 2,
  "data": [
    {
      "id": 1,
      "name": "John Doe",
      "dob": "2000-01-01",
      "age": 26
    },
    {
      "id": 2,
      "name": "Jane Doe",
      "dob": "1998-05-10",
      "age": 27
    }
  ],
  "limit": 10,
  "page": 1
}
```

### Update User

**Request**

```http
PUT /users/1
Content-Type: application/json
{
  "name": "Updated Name",
  "dob": "2000-01-01"
}
```

**Response**

```http
{
  "id": 1,
  "name": "Updated Name",
  "dob": "2000-01-01",
}
```

### Delete User

**Request**

```http
DELETE /users/1
```

**Response**

```http
204 No Content
```
