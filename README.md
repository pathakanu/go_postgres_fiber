# Go Postgres Fiber

This is a simple REST API built with Go, Fiber, and GORM for managing a collection of books stored in a PostgreSQL database.

## Features

*   Create a new book
*   Get all books
*   Get a book by its ID
*   Delete a book by its ID

## API Endpoints

| Method | Endpoint                | Description          |
|--------|-------------------------|----------------------|
| POST   | `/api/create_books`     | Creates a new book.  |
| GET    | `/api/get_all_books`    | Retrieves all books. |
| GET    | `/api/get_book/:id`     | Retrieves a single book by its ID. |
| DELETE | `/api/delete_book/:id`  | Deletes a book by its ID. |

## Getting Started

### Prerequisites

*   [Go](https://golang.org/doc/install) (version 1.15+ recommended)
*   [PostgreSQL](https://www.postgresql.org/download/)
*   [Docker](https://www.docker.com/get-started) (optional, for running Postgres in a container)

### Installation

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/pathakanu/go_postgres_fiber.git
    cd go_postgres_fiber
    ```

2.  **Install dependencies:**
    ```bash
    go mod tidy
    ```

3.  **Set up environment variables:**
    Create a `.env` file in the root of the project and add the following variables:
    ```env
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=your_postgres_user
    DB_PASSWORD=your_postgres_password
    DB_NAME=your_database_name
    DB_SSLMODE=disable
    ```

### Usage

1.  **Run the application:**
    ```bash
    go run main.go
    ```
    The server will start on `http://localhost:3000`.

2.  **Example cURL requests:**

    *   **Create a book:**
        ```bash
        curl -X POST http://localhost:3000/api/create_books -H "Content-Type: application/json" -d '{"title":"The Go Programming Language","author":"Alan A. A. Donovan","year":2015}'
        ```

    *   **Get all books:**
        ```bash
        curl http://localhost:3000/api/get_all_books
        ```

    *   **Get a book by ID:**
        ```bash
        curl http://localhost:3000/api/get_book/1
        ```

    *   **Delete a book by ID:**
        ```bash
        curl -X DELETE http://localhost:3000/api/delete_book/1
        ```
