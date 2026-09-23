# Go-lang Blog Apps API

A modular, production-ready RESTful Blog API backend built with **Go**, **Echo v5**, **GORM**, and **PostgreSQL**. Features JWT-based authentication, user management, full CRUD operations for posts with automatic slug generation, and hierarchical (threaded) comments.

---

## Features

- **Authentication & Security**: Secure user signup, login, password hashing with bcrypt, and stateless JWT token authentication middleware.
- **Posts Management**: Full CRUD operations for blog posts with auto-generated URL-friendly slugs and author associations.
- **Hierarchical Comments**: Create, read, update, and delete comments on posts, with support for nested/threaded replies (`parent_id`).
- **Clean Architecture**: Organized separation of concerns across Handlers, Services, Repositories, and Data Models.
- **Database & Migrations**: GORM with PostgreSQL driver, automated schema migrations, and UUID primary keys.

---

## Project Folder Architecture

The project follows standard Go project structure standards and clean layered architecture:

```text
Go-lang-Blog-Apps/
├── cmd/
│   └── server/
│       └── main.go                 # Application entry point
├── config/
│   └── config.go                   # Environment configuration loader
├── internal/
│   ├── app/
│   │   ├── middleware.go           # JWT authentication middleware
│   │   ├── routes.go               # Echo route definitions and dependency injection
│   │   └── server.go               # Server initialization and startup
│   ├── auth/
│   │   ├── jwt.go                  # JWT token generation and claims parsing
│   │   └── password.go             # Bcrypt password hashing and verification
│   ├── database/
│   │   └── db.go                   # PostgreSQL connection & GORM AutoMigrate
│   ├── handlers/
│   │   ├── auth_handller.go        # HTTP handlers for user signup and login
│   │   ├── comment_handller.go     # HTTP handlers for comment CRUD
│   │   └── post_handler.go         # HTTP handlers for post CRUD
│   ├── models/
│   │   ├── comment.go              # Comment model & database schema
│   │   ├── post.go                 # Post model & database schema
│   │   └── user.go                 # User model & database schema
│   ├── repositories/
│   │   ├── comment_repository.go   # Comment database queries
│   │   ├── post_repository.go      # Post database queries
│   │   └── user_repository.go      # User database queries
│   ├── services/
│   │   ├── comment_service.go      # Comment business logic
│   │   ├── post_service.go         # Post business logic (slug generation, etc.)
│   │   └── user_service.go         # User business logic
│   └── utils/
│       ├── response.go             # Standardized JSON response helpers
│       └── slug.go                 # Utility to generate URL-safe slugs
├── .env                            # Environment variables (local)
├── .env.example                    # Sample environment variables template
├── .gitignore                      # Git ignore rules
├── go.mod                          # Go module dependencies
└── README.md                       # Project documentation
```

### Architectural Layers
- **Handlers (`internal/handlers`)**: Parse HTTP requests, validate input, bind payloads, and return JSON responses.
- **Services (`internal/services`)**: Encapsulate business logic, validations, and orchestrate repository calls.
- **Repositories (`internal/repositories`)**: Direct database access and GORM queries.
- **Models (`internal/models`)**: Data definitions, GORM database tags, and JSON serialization tags.

---

## Getting Started & Setup

### 1. Prerequisites
- **Go**: Version 1.22 or later installed ([Download Go](https://go.dev/dl/))
- **PostgreSQL**: Version 13 or later running locally or via Docker
- **PostgreSQL Extension**: Ensure the `uuid-ossp` extension is available (enabled automatically by Postgres or run `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

### 2. Clone the Repository
```bash
git clone https://github.com/mamun-jsx/Go-lang-Blog-Apps.git
cd Go-lang-Blog-Apps
```

### 3. Configure Environment Variables
Copy `.env.example` to create your `.env` file:
```bash
cp .env.example .env
```

Edit `.env` to match your local PostgreSQL configuration:
```env
APP_PORT=8080

# Database credentials
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_postgres_password
DB_NAME=my_blog_app

# Secret key for JWT signing
JWT_SECRET=super_secret_jwt_key_change_me
```

### 4. Create the Database
Ensure PostgreSQL is running, then create the database specified in `DB_NAME`:
```bash
# Using PostgreSQL CLI:
createdb my_blog_app -U postgres

# Or inside psql:
psql -U postgres -c "CREATE DATABASE my_blog_app;"
```

### 5. Install Dependencies
```bash
go mod tidy
```

### 6. Run the Application
Start the server:
```bash
go run cmd/server/main.go
```

The database tables (`users`, `posts`, `comments`) will be automatically created on startup via GORM AutoMigrate. You should see:
```text
Successfully connected to postgreSQL
⇨ http server started on [::]:8080
```

---

## API Endpoints Reference

Base URL: `http://localhost:8080/api/v1`

Protected endpoints require the `Authorization` header:
```http
Authorization: Bearer <your_jwt_token>
```

### Summary Table

| Method | Endpoint | Auth Required | Description |
|---|---|---|---|
| `POST` | `/api/v1/signup` | No | Register a new user |
| `POST` | `/api/v1/login` | No | Log in and obtain JWT token |
| `POST` | `/api/v1/posts` | **Yes** | Create a new blog post |
| `GET` | `/api/v1/posts` | **Yes** | Get list of all blog posts |
| `GET` | `/api/v1/posts/:id` | **Yes** | Get a single blog post by ID |
| `PUT` | `/api/v1/posts/:id` | **Yes** | Update a post by ID |
| `DELETE` | `/api/v1/posts/:id` | **Yes** | Delete a post by ID |
| `POST` | `/api/v1/posts/:id/comments` | **Yes** | Add comment/reply to a post |
| `GET` | `/api/v1/posts/:id/comments` | **Yes** | Get all comments for a post |
| `PUT` | `/api/v1/comments/:id` | **Yes** | Update a comment by ID |
| `DELETE` | `/api/v1/comments/:id` | **Yes** | Delete a comment by ID |

---

### Authentication Endpoints

#### 1. User Signup
- **Method**: `POST`
- **Path**: `/api/v1/signup`
- **Request Body**:
```json
{
  "user_name": "johndoe",
  "email": "john@example.com",
  "password": "secretpassword"
}
```
- **Response** (`201 Created`):
```json
{
  "success": true,
  "message": "user_created",
  "data": {
    "id": "e305e94b-4fc8-47fb-a0ef-f15d22cf8be5",
    "user_name": "johndoe",
    "email": "john@example.com",
    "created_at": "2026-09-23T10:00:00Z",
    "updated_at": "2026-09-23T10:00:00Z"
  }
}
```

#### 2. User Login
- **Method**: `POST`
- **Path**: `/api/v1/login`
- **Request Body**:
```json
{
  "email": "john@example.com",
  "password": "secretpassword"
}
```
- **Response** (`200 OK`):
```json
{
  "success": true,
  "message": "login succesfull",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": "e305e94b-4fc8-47fb-a0ef-f15d22cf8be5",
      "user_name": "johndoe",
      "email": "john@example.com",
      "created_at": "2026-09-23T10:00:00Z",
      "updated_at": "2026-09-23T10:00:00Z"
    }
  }
}
```

---

### Posts Endpoints

#### 3. Create Post
- **Method**: `POST`
- **Path**: `/api/v1/posts`
- **Headers**: `Authorization: Bearer <token>`
- **Request Body**:
```json
{
  "title": "My First Blog Post",
  "content": "This is the content of my first post."
}
```
- **Response** (`201 Created`):
```json
{
  "success": true,
  "message": "post created",
  "data": {
    "id": "1198656f-2391-4e78-9e67-d6e04d41e7f0",
    "author_id": "e305e94b-4fc8-47fb-a0ef-f15d22cf8be5",
    "title": "My First Blog Post",
    "slug": "my-first-blog-post",
    "content": "This is the content of my first post.",
    "created_at": "2026-09-23T10:05:00Z",
    "updated_at": "2026-09-23T10:05:00Z"
  }
}
```

#### 4. List All Posts
- **Method**: `GET`
- **Path**: `/api/v1/posts`
- **Headers**: `Authorization: Bearer <token>`
- **Response** (`200 OK`):
```json
{
  "success": true,
  "message": "posts",
  "data": [
    {
      "id": "1198656f-2391-4e78-9e67-d6e04d41e7f0",
      "author_id": "e305e94b-4fc8-47fb-a0ef-f15d22cf8be5",
      "title": "My First Blog Post",
      "slug": "my-first-blog-post",
      "content": "This is the content of my first post.",
      "created_at": "2026-09-23T10:05:00Z",
      "updated_at": "2026-09-23T10:05:00Z"
    }
  ]
}
```

#### 5. Get Post by ID
- **Method**: `GET`
- **Path**: `/api/v1/posts/:id`
- **Headers**: `Authorization: Bearer <token>`
- **Response** (`200 OK`):
```json
{
  "success": true,
  "message": "post",
  "data": {
    "id": "1198656f-2391-4e78-9e67-d6e04d41e7f0",
    "author_id": "e305e94b-4fc8-47fb-a0ef-f15d22cf8be5",
    "title": "My First Blog Post",
    "slug": "my-first-blog-post",
    "content": "This is the content of my first post.",
    "created_at": "2026-09-23T10:05:00Z",
    "updated_at": "2026-09-23T10:05:00Z"
  }
}
```

#### 6. Update Post
- **Method**: `PUT`
- **Path**: `/api/v1/posts/:id`
- **Headers**: `Authorization: Bearer <token>`
- **Request Body**:
```json
{
  "title": "Updated Post Title",
  "content": "Updated content goes here."
}
```
- **Response** (`200 OK`):
```json
{
  "success": true,
  "message": "post updated",
  "data": {
    "id": "1198656f-2391-4e78-9e67-d6e04d41e7f0",
    "author_id": "e305e94b-4fc8-47fb-a0ef-f15d22cf8be5",
    "title": "Updated Post Title",
    "slug": "updated-post-title",
    "content": "Updated content goes here.",
    "created_at": "2026-09-23T10:05:00Z",
    "updated_at": "2026-09-23T10:10:00Z"
  }
}
```

#### 7. Delete Post
- **Method**: `DELETE`
- **Path**: `/api/v1/posts/:id`
- **Headers**: `Authorization: Bearer <token>`
- **Response** (`200 OK`):
```json
{
  "success": true,
  "message": "post deleted",
  "data": null
}
```

---

### Comments Endpoints

#### 8. Add Comment to a Post
- **Method**: `POST`
- **Path**: `/api/v1/posts/:id/comments`
- **Headers**: `Authorization: Bearer <token>`
- **Request Body (Top-level comment)**:
```json
{
  "content": "Great article! Thanks for sharing."
}
```
- **Request Body (Nested reply to an existing comment)**:
```json
{
  "content": "I agree with your comment!",
  "parent_id": "parent-comment-uuid-here"
}
```
- **Response** (`201 Created`):
```json
{
  "success": true,
  "message": "comment added",
  "data": {
    "id": "4286f916-d34e-4fba-bb5f-2c35470d0617",
    "post_id": "1198656f-2391-4e78-9e67-d6e04d41e7f0",
    "user_id": "e305e94b-4fc8-47fb-a0ef-f15d22cf8be5",
    "parent_id": null,
    "content": "Great article! Thanks for sharing.",
    "created_at": "2026-09-23T10:15:00Z",
    "updated_at": "2026-09-23T10:15:00Z"
  }
}
```

#### 9. List Comments of a Post
- **Method**: `GET`
- **Path**: `/api/v1/posts/:id/comments`
- **Headers**: `Authorization: Bearer <token>`
- **Response** (`200 OK`):
```json
{
  "success": true,
  "message": "comments",
  "data": [
    {
      "id": "4286f916-d34e-4fba-bb5f-2c35470d0617",
      "post_id": "1198656f-2391-4e78-9e67-d6e04d41e7f0",
      "user_id": "e305e94b-4fc8-47fb-a0ef-f15d22cf8be5",
      "parent_id": null,
      "content": "Great article! Thanks for sharing.",
      "created_at": "2026-09-23T10:15:00Z",
      "updated_at": "2026-09-23T10:15:00Z"
    }
  ]
}
```

#### 10. Update Comment
- **Method**: `PUT`
- **Path**: `/api/v1/comments/:id`
- **Headers**: `Authorization: Bearer <token>`
- **Request Body**:
```json
{
  "content": "Updated comment text."
}
```
- **Response** (`200 OK`):
```json
{
  "success": true,
  "message": "comment updated",
  "data": {
    "id": "4286f916-d34e-4fba-bb5f-2c35470d0617",
    "post_id": "1198656f-2391-4e78-9e67-d6e04d41e7f0",
    "user_id": "e305e94b-4fc8-47fb-a0ef-f15d22cf8be5",
    "content": "Updated comment text.",
    "created_at": "2026-09-23T10:15:00Z",
    "updated_at": "2026-09-23T10:20:00Z"
  }
}
```

#### 11. Delete Comment
- **Method**: `DELETE`
- **Path**: `/api/v1/comments/:id`
- **Headers**: `Authorization: Bearer <token>`
- **Response** (`200 OK`):
```json
{
  "success": true,
  "message": "comment deleted",
  "data": null
}
```

---

## License

This project is open-source and available under the [MIT License](LICENSE).
