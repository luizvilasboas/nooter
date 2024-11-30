# nooter

nooter is a web application for creating and managing notes. This project consists of a REST API built with Go, using the Gin framework, Swagger for API documentation, and GORM for database integration.

## Technologies Used

- **[Gin](https://gin-gonic.com/):** Web framework for Go.
- **[GORM](https://gorm.io/):** Object-Relational Mapping (ORM) library for Go.
- **[Swagger](https://swagger.io/):** Tool for documenting APIs.
- **Database:** Supports multiple databases (MySQL, PostgreSQL, SQLite, etc.), configurable as needed.

## Features

- CRUD (Create, Read, Update, Delete) operations for **notes**.
- Notes organization by **categories**.
- Routes grouped under the **v1** version of the API.
- Default category assigned to notes if no category is specified during creation.
- Automatic timestamps (`created_at`, `updated_at`) to track changes.
- Interactive API documentation available via Swagger.

## Installation and Setup

### Prerequisites

- [Go](https://golang.org/) (version 1.20 or higher)
- Configured database (MySQL, PostgreSQL, SQLite, or other supported by GORM)

### Steps

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/nooter.git
   cd nooter
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Start the server:

   ```bash
   go run main.go
   ```

4. Access the interactive Swagger documentation at: `http://localhost:8080/swagger/index.html`.

### Using docker

You can pull the image from [Docker Hub](https://hub.docker.com/repository/docker/olooeez/nooter/general).

## Testing

Run tests using the following command:

```bash
go test ./...
```

## Main Routes

### Categories
- `GET /api/v1/categories`: List all categories.
- `POST /api/v1/categories`: Create a new category.
- `GET /api/v1/categories/:id`: Retrieve details of a category.
- `PUT /api/v1/categories/:id`: Update a category.
- `DELETE /api/v1/categories/:id`: Delete a category.

### Notes
- `GET /api/v1/notes`: List all notes.
- `POST /api/v1/notes`: Create a new note.
- `GET /api/v1/notes/:id`: Retrieve details of a note.
- `PUT /api/v1/notes/:id`: Update a note.
- `DELETE /api/v1/notes/:id`: Delete a note.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to improve the project.

## License

This project is licensed under the [MIT License](https://gitlab.com/olooeez/nooter/-/blob/main/LICENSE).
