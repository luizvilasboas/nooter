package handlers

import (
	"encoding/json"
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
	todos, err := h.Storage.GetTodos()
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	// Retorna uma lista vazia se não houver tarefas
	if todos == nil {
		todos = []models.Todo{}
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

	createdTodo, err := h.Storage.AddTodo(todo)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
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

	todo, err := h.Storage.GetTodoByID(id)
	if err != nil {
		writeJSONError(w, http.StatusNotFound, "Not Found", "Todo not found")
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
		Title:   req.Title,
		Details: req.Details,
		Done:    req.Done,
	}

	todo, err := h.Storage.UpdateTodo(id, updatedTodo)
	if err != nil {
		writeJSONError(w, http.StatusNotFound, "Not Found", "Todo not found")
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

	err = h.Storage.DeleteTodoByID(id)
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
