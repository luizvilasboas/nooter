package storage

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlite3"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/mattn/go-sqlite3"
	"gitlab.com/olooeez/nooter/models"
)

type Storage interface {
	InitDB(filepath string)
	GetTodos() ([]models.Todo, error)
	AddTodo(todo models.Todo) (models.Todo, error)
	GetTodoByID(id int) (models.Todo, error)
	UpdateTodo(id int, updatedTodo models.Todo) (models.Todo, error)
	DeleteTodoByID(id int) error
}

type DefaultStorage struct {
	db *sql.DB
}

func (ds *DefaultStorage) InitDB(filepath string) {
	var err error

	ds.db, err = sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	if err := ds.db.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Database connection established")

	ds.runMigrations()
}

func (ds *DefaultStorage) runMigrations() {
	if ds.db == nil {
		log.Fatal("Database is not initialized")
	}

	driver, err := sqlite3.WithInstance(ds.db, &sqlite3.Config{})
	if err != nil {
		log.Fatalf("Could not start SQL driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"sqlite3", driver)

	if err != nil {
		log.Fatalf("Migration setup failed: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Could not run up migrations: %v", err)
	}

	fmt.Println("Migrations ran successfully")
}

func (ds *DefaultStorage) GetTodos() ([]models.Todo, error) {
	rows, err := ds.db.Query("SELECT id, title, details, done FROM todos")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch todos: %v", err)
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Details, &todo.Done); err != nil {
			return nil, fmt.Errorf("failed to scan todo: %v", err)
		}
		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return todos, nil
}

func (ds *DefaultStorage) AddTodo(todo models.Todo) (models.Todo, error) {
	result, err := ds.db.Exec("INSERT INTO todos (title, details, done) VALUES (?, ?, ?)", todo.Title, todo.Details, todo.Done)
	if err != nil {
		return models.Todo{}, fmt.Errorf("failed to insert todo: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Todo{}, fmt.Errorf("failed to get last insert id: %v", err)
	}

	todo.ID = int(id)
	return todo, nil
}

func (ds *DefaultStorage) GetTodoByID(id int) (models.Todo, error) {
	row := ds.db.QueryRow("SELECT id, title, details, done FROM todos WHERE id = ?", id)

	var todo models.Todo
	if err := row.Scan(&todo.ID, &todo.Title, &todo.Details, &todo.Done); err != nil {
		if err == sql.ErrNoRows {
			return models.Todo{}, fmt.Errorf("todo not found with id: %d", id)
		}
		return models.Todo{}, fmt.Errorf("failed to fetch todo: %v", err)
	}

	return todo, nil
}

func (ds *DefaultStorage) UpdateTodo(id int, updatedTodo models.Todo) (models.Todo, error) {
	_, err := ds.db.Exec("UPDATE todos SET title = ?, details = ?, done = ? WHERE id = ?", updatedTodo.Title, updatedTodo.Details, updatedTodo.Done, id)
	if err != nil {
		return models.Todo{}, fmt.Errorf("failed to update todo: %v", err)
	}

	updatedTodo.ID = id
	return updatedTodo, nil
}

func (ds *DefaultStorage) DeleteTodoByID(id int) error {
	result, err := ds.db.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed to delete todo: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("todo not found with id: %d", id)
	}

	return nil
}
