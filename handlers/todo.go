package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gitlab.com/olooeez/nooter/models"
	"gitlab.com/olooeez/nooter/storage"
)

// GetTodos godoc
// @Summary Lista todas as tarefas
// @Description Retorna uma lista de todas as tarefas cadastradas
// @Tags todos
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Todo
// @Router /todos [get]
func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := storage.GetTodos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// CreateTodo godoc
// @Summary Cria uma nova tarefa
// @Description Adiciona uma nova tarefa ao sistema
// @Tags todos
// @Accept  json
// @Produce  json
// @Param todo body models.Todo true "Nova Tarefa"
// @Success 201 {object} models.Todo
// @Router /todos [post]
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	createdTodo, err := storage.AddTodo(todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdTodo)
}

// GetTodoByID godoc
// @Summary Retorna uma tarefa específica
// @Description Busca uma tarefa pelo ID
// @Tags todos
// @Accept  json
// @Produce  json
// @Param id path int true "ID da Tarefa"
// @Success 200 {object} models.Todo
// @Failure 404 {object} map[string]string
// @Router /todos/{id} [get]
func GetTodoByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	todo, err := storage.GetTodoByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

// UpdateTodo godoc
// @Summary Atualiza uma tarefa existente
// @Description Atualiza os detalhes de uma tarefa pelo ID
// @Tags todos
// @Accept  json
// @Produce  json
// @Param id path int true "ID da Tarefa"
// @Param todo body models.Todo true "Atualização da Tarefa"
// @Success 200 {object} models.Todo
// @Failure 404 {object} map[string]string
// @Router /todos/{id} [put]
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedTodo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	todo, err := storage.UpdateTodo(id, updatedTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

// DeleteTodoByID godoc
// @Summary Deleta uma tarefa
// @Description Remove uma tarefa pelo ID fornecido
// @Tags todos
// @Accept  json
// @Produce  json
// @Param id path int true "ID da Tarefa"
// @Success 204
// @Failure 404 {object} map[string]string
// @Router /todos/{id} [delete]
func DeleteTodoByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = storage.DeleteTodoByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
