# Blog API with Middleware

A professional, robust RESTful API for a blogging platform built with **Go**. This project demonstrates the implementation of custom middleware, secure authentication, and a clean architecture following **Domain-Driven Design (DDD)** principles.

## 🚀 Features

- **Secure Authentication**: JWT-based authentication with password hashing using `bcrypt`.
- **Custom Middleware Stack**:
  - **Rate Limiting**: Protects the API from brute-force attacks and abuse using a token bucket algorithm.
  - **JWT Authentication**: Ensures secure access to protected resources.
  - **Logging**: Structured request/response logging for better observability.
  - **CORS & Preflight**: Full support for Cross-Origin Resource Sharing.
- **Graceful Shutdown**: Ensures all in-flight requests are completed before the server shuts down, preventing data loss and enhancing reliability.
- **Comprehensive Post Management**:
  - Full CRUD operations for blog posts.
  - **Author-only Permissions**: Secure post modifications where only the original author can update or delete their content.
  - **Slug Generation**: Automatic SEO-friendly URL slugs for posts.
- **Advanced Querying**:
  - Search functionality across posts.
  - Filtering by categories and tags.
  - Pagination support for large datasets.
- **Database Excellence**:
  - PostgreSQL integration using `sqlx` for efficient data mapping.
  - Automated database migrations with `sql-migrate`.

## 🛠️ Tech Stack

- **Language**: [Go](https://golang.org/) (v1.25+)
- **Database**: [PostgreSQL](https://www.postgresql.org/)
- **Libraries**:
  - `sqlx`: General purpose extensions to `database/sql`.
  - `godotenv`: Environment variable management.
  - `slug`: SEO-friendly slug generation.
  - `bcrypt`: Secure password hashing.
  - `sql-migrate`: Database schema migrations.
  - `x/time/rate`: High-performance rate limiting.

## 📋 Prerequisites

- Go 1.25 or higher
- PostgreSQL instance
- `sql-migrate` tool (optional, handled internally by the app)

## ⚙️ Installation & Setup

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/mhbhuiyan99/Blog-API-with-Middleware.git
   cd Blog-API-with-Middleware
   ```

2. **Configure Environment Variables**:
   Create a `.env` file in the root directory and populate it with your configuration:
   ```env
   VERSION=1.0.0
   SERVICE_NAME=BlogAPI
   HTTP_PORT=4000
   JWT_SECRET_KEY=your_secret_key_here

   DB_HOST=localhost
   DB_PORT=5432
   DB_NAME=blogAPI
   DB_USER=postgres
   DB_PASSWORD=your_password
   DB_ENABLE_SSL_MODE=false

   RATE_LIMITER_ENABLED=true
   RATE_LIMITER_RPS=2
   RATE_LIMITER_BURST=4
   ```

3. **Install Dependencies**:
   ```bash
   go mod download
   ```

4. **Run the Application**:
   The application will automatically handle database migrations on startup.
   ```bash
   go run main.go
   ```

## 🛣️ API Endpoints

### User Authentication
| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| POST | `/users` | Register a new user | No |
| POST | `/users/login` | Login and receive JWT | No |

### Post Management
| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| GET | `/posts` | Get all published posts | No |
| GET | `/posts/{id}` | Get a specific post by ID | No |
| POST | `/posts` | Create a new blog post | **Yes** |
| PUT | `/posts/{id}` | Update an existing post | **Yes** (Author only) |
| DELETE | `/posts/{id}` | Delete a post | **Yes** (Author only) |
| GET | `/posts/search` | Search posts by keyword | No |
| GET | `/posts/category/{category}` | Filter posts by category | No |
| GET | `/posts/tag/{tag}` | Filter posts by tag | No |

### User Specific
| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| GET | `/my-posts/drafts` | Get current user's drafts | **Yes** |
| GET | `/my-posts/published` | Get current user's published posts | **Yes** |
| GET | `/users/{userId}/posts` | Get posts by a specific user | No |

## 🧪 Testing

### Rate Limiting

You can test the rate limiting middleware by sending rapid requests:
```bash
for i in {1..8}; do curl -s -w "Request $i: %{http_code}\n" http://localhost:4000/posts; done
```

### Graceful Shutdown

To test the graceful shutdown functionality, you can simulate long-running requests and then send an interrupt signal to the server.

**Simulate a slow response and interrupt the server**:
    In one terminal, start a slow request:
    ```bash
    curl --limit-rate 1 http://localhost:4000/posts
    ```
    Immediately press `Ctrl+C` in the terminal running your server. If graceful shutdown works, the `curl` command will continue to receive data until the response finishes (or your 5s timeout kicks in).


## 🏗️ Project Structure

```text
├── cmd/            # Application entry point and server setup
├── config/         # Configuration management
├── domain/         # Core domain models and logic
├── infra/          # Infrastructure (database connections)
├── migrations/     # SQL migration files
├── post/           # Post domain logic (DDD)
├── user/           # User domain logic (DDD)
├── repo/           # Repository layer for data access
├── rest/           # REST API handlers and middlewares
└── util/           # Utility functions
```

---
*Developed as a learning project for mastering Middleware and Authentication in Go.*
