# nooter

> A REST API for creating and managing notes, built with Go, Gin, and GORM.

## About the Project

`nooter` is a web application for creating and managing notes. This repository contains the REST API backend, which features CRUD operations for notes and categories, interactive API documentation via Swagger, and flexible database support through GORM.

Key features include:
*   CRUD operations for notes and categories.
*   Automatic timestamps to track creation and updates.
*   Default category assignment for new notes.
*   Interactive API documentation with Swagger.

## Tech Stack

*   [Go](https://golang.org/)
*   [Gin](https://gin-gonic.com/) (Web Framework)
*   [GORM](https://gorm.io/) (ORM)
*   [Swagger](https://swagger.io/) (API Documentation)

## Usage

Below are the instructions for you to set up and run the project locally.

### Prerequisites

You need to have the following software installed:

*   [Go](https://golang.org/dl/) (version 1.20 or higher)
*   A configured database supported by GORM (e.g., MySQL, PostgreSQL, SQLite)

### Installation and Setup

Follow the steps below:

1.  **Clone the repository**
    ```bash
    git clone https://github.com/luizvilasboas/nooter.git
    ```

2.  **Navigate to the project directory**
    ```bash
    cd nooter
    ```

3.  **Install dependencies**
    ```bash
    go mod tidy
    ```

### Workflow

1.  **Run the server**
    ```bash
    go run main.go
    ```
    The server will start on `http://localhost:8080`.

2.  **Access API Documentation**

    Interactive API documentation is available via Swagger at:
    **[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)**

### Docker

You can also run the application using Docker. A pre-built image is available on [Docker Hub](https://hub.docker.com/repository/docker/olooeez/nooter/general).

### Testing

To run the project's tests, execute the following command:
```bash
go test ./...
```

## API Routes

### Categories
- `GET /api/v1/categories`: List all categories.
- `POST /api/v1/categories`: Create a new category.
- `GET /api/v1/categories/:id`: Retrieve a specific category.
- `PUT /api/v1/categories/:id`: Update a category.
- `DELETE /api/v1/categories/:id`: Delete a category.

### Notes
- `GET /api/v1/notes`: List all notes.
- `POST /api/v1/notes`: Create a new note.
- `GET /api/v1/notes/:id`: Retrieve a specific note.
- `PUT /api/v1/notes/:id`: Update a note.
- `DELETE /api/v1/notes/:id`: Delete a note.

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
