package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gitlab.com/olooeez/nooter/models"
	"gitlab.com/olooeez/nooter/storage"
)

type JSONError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

type Handlers struct {
	Storage storage.Storage
}

// GetTodos godoc
// @Summary Lista todas as tarefas
// @Description Retorna uma lista de todas as tarefas cadastradas
// @Tags todos
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Todo
// @Failure 500 {object} JSONError
// @Router /todos [get]
func (h *Handlers) GetTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo
	err := h.Storage.Read("todos", "1=1", &todos)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
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
// @Param todo body models.TodoRequest true "Nova Tarefa"
// @Success 201 {object} models.Todo
// @Failure 400 {object} JSONError
// @Failure 500 {object} JSONError
// @Router /todos [post]
func (h *Handlers) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var req models.TodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, http.StatusBadRequest, "Bad Request", "Invalid JSON payload")
		return
	}

	todo := models.Todo{
		Title:   req.Title,
		Details: req.Details,
		Done:    req.Done,
	}

	id, err := h.Storage.Create("todos", &todo)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	todo.ID = id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

// GetTodoByID godoc
// @Summary Retorna uma tarefa específica
// @Description Busca uma tarefa pelo ID
// @Tags todos
// @Accept  json
// @Produce  json
// @Param id path int true "ID da Tarefa"
// @Success 200 {object} models.Todo
// @Failure 400 {object} JSONError
// @Failure 404 {object} JSONError
// @Failure 500 {object} JSONError
// @Router /todos/{id} [get]
func (h *Handlers) GetTodoByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "Bad Request", "Invalid ID format")
		return
	}

	var todos []models.Todo
	condition := fmt.Sprintf("id = %d", id)
	err = h.Storage.Read("todos", condition, &todos)
	if err != nil || len(todos) == 0 {
		writeJSONError(w, http.StatusNotFound, "Not Found", "Todo not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos[0])
}

// UpdateTodo godoc
// @Summary Atualiza uma tarefa existente
// @Description Atualiza os detalhes de uma tarefa pelo ID
// @Tags todos
// @Accept  json
// @Produce  json
// @Param id path int true "ID da Tarefa"
// @Param todo body models.TodoRequest true "Atualização da Tarefa"
// @Success 200 {object} models.Todo
// @Failure 400 {object} JSONError
// @Failure 404 {object} JSONError
// @Failure 500 {object} JSONError
// @Router /todos/{id} [put]
func (h *Handlers) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "Bad Request", "Invalid ID format")
		return
	}

	var req models.TodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, http.StatusBadRequest, "Bad Request", "Invalid JSON payload")
		return
	}

	updatedTodo := models.Todo{
		ID:      id,
		Title:   req.Title,
		Details: req.Details,
		Done:    req.Done,
	}

	condition := fmt.Sprintf("id = %d", id)
	err = h.Storage.Update("todos", &updatedTodo, condition)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTodo)
}

// DeleteTodoByID godoc
// @Summary Deleta uma tarefa
// @Description Remove uma tarefa pelo ID fornecido
// @Tags todos
// @Accept  json
// @Produce  json
// @Param id path int true "ID da Tarefa"
// @Success 204
// @Failure 400 {object} JSONError
// @Failure 404 {object} JSONError
// @Failure 500 {object} JSONError
// @Router /todos/{id} [delete]
func (h *Handlers) DeleteTodoByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "Bad Request", "Invalid ID format")
		return
	}

	condition := fmt.Sprintf("id = %d", id)
	err = h.Storage.Delete("todos", condition)
	if err != nil {
		writeJSONError(w, http.StatusNotFound, "Not Found", "Todo not found")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func writeJSONError(w http.ResponseWriter, status int, errorTitle, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(JSONError{
		Error:   errorTitle,
		Message: message,
	})
}
