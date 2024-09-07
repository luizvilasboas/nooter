package storage

import (
	"database/sql"
	"errors"

	"gitlab.com/olooeez/nooter/models"
)

func GetTodos() ([]models.Todo, error) {
	rows, err := db.Query("SELECT id, title, details, done FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Details, &todo.Done)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func AddTodo(todo models.Todo) (models.Todo, error) {
	result, err := db.Exec("INSERT INTO todos (title, details, done) VALUES (?, ?, ?)", todo.Title, todo.Details, todo.Done)
	if err != nil {
		return models.Todo{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Todo{}, err
	}

	todo.ID = int(id)
	return todo, nil
}

func GetTodoByID(id int) (models.Todo, error) {
	row := db.QueryRow("SELECT id, title, details, done FROM todos WHERE id = ?", id)

	var todo models.Todo
	err := row.Scan(&todo.ID, &todo.Title, &todo.Details, &todo.Done)
	if err == sql.ErrNoRows {
		return models.Todo{}, errors.New("todo not found")
	} else if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

func UpdateTodo(id int, updatedTodo models.Todo) (models.Todo, error) {
	_, err := db.Exec("UPDATE todos SET title = ?, details = ?, done = ? WHERE id = ?", updatedTodo.Title, updatedTodo.Details, updatedTodo.Done, id)
	if err != nil {
		return models.Todo{}, err
	}

	updatedTodo.ID = id
	return updatedTodo, nil
}

func DeleteTodoByID(id int) error {
	_, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		return errors.New("todo not found")
	}

	return nil
}
