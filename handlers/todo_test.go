package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"gitlab.com/olooeez/nooter/mocks"
	"gitlab.com/olooeez/nooter/models"
)

// setupTest initializes the mock controller, mock storage, and handlers, and returns them
func setupTest(t *testing.T) (*gomock.Controller, *Handlers, *mocks.MockStorage) {
	ctrl := gomock.NewController(t)
	mockStorage := mocks.NewMockStorage(ctrl)
	handlers := &Handlers{Storage: mockStorage}
	return ctrl, handlers, mockStorage
}

func TestGetTodos(t *testing.T) {
	ctrl, handlers, mockStorage := setupTest(t)
	defer ctrl.Finish()

	mockStorage.EXPECT().GetTodos().Return([]models.Todo{}, nil)

	req, _ := http.NewRequest(http.MethodGet, "/todos", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.GetTodos)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var todos []models.Todo
	if err := json.NewDecoder(rr.Body).Decode(&todos); err != nil {
		t.Errorf("could not decode response: %v", err)
	}

	if len(todos) != 0 {
		t.Errorf("expected no todos, got %v", todos)
	}
}

func TestCreateTodo(t *testing.T) {
	ctrl, handlers, mockStorage := setupTest(t)
	defer ctrl.Finish()

	newTodo := models.TodoRequest{
		Title:   "Test Title",
		Details: "Test Details",
		Done:    false,
	}

	expectedTodo := models.Todo{
		ID:      1,
		Title:   "Test Title",
		Details: "Test Details",
		Done:    false,
	}

	mockStorage.EXPECT().AddTodo(gomock.Any()).Return(expectedTodo, nil)

	body, _ := json.Marshal(newTodo)
	req, _ := http.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.CreateTodo)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var todo models.Todo
	if err := json.NewDecoder(rr.Body).Decode(&todo); err != nil {
		t.Errorf("could not decode response: %v", err)
	}

	if todo != expectedTodo {
		t.Errorf("expected %v, got %v", expectedTodo, todo)
	}
}

func TestGetTodoByID(t *testing.T) {
	ctrl, handlers, mockStorage := setupTest(t)
	defer ctrl.Finish()

	expectedTodo := models.Todo{
		ID:      1,
		Title:   "Test Title",
		Details: "Test Details",
		Done:    false,
	}

	mockStorage.EXPECT().GetTodoByID(1).Return(expectedTodo, nil)

	req, _ := http.NewRequest(http.MethodGet, "/todos/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.GetTodoByID)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var todo models.Todo
	if err := json.NewDecoder(rr.Body).Decode(&todo); err != nil {
		t.Errorf("could not decode response: %v", err)
	}

	if todo != expectedTodo {
		t.Errorf("expected %v, got %v", expectedTodo, todo)
	}
}

func TestUpdateTodo(t *testing.T) {
	ctrl, handlers, mockStorage := setupTest(t)
	defer ctrl.Finish()

	updatedTodo := models.Todo{
		ID:      1,
		Title:   "Updated Title",
		Details: "Updated Details",
		Done:    true,
	}

	mockStorage.EXPECT().UpdateTodo(1, gomock.Any()).Return(updatedTodo, nil)

	reqBody, _ := json.Marshal(models.TodoRequest{
		Title:   "Updated Title",
		Details: "Updated Details",
		Done:    true,
	})

	req, _ := http.NewRequest(http.MethodPut, "/todos/1", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.UpdateTodo)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var todo models.Todo
	if err := json.NewDecoder(rr.Body).Decode(&todo); err != nil {
		t.Errorf("could not decode response: %v", err)
	}

	if todo != updatedTodo {
		t.Errorf("expected %v, got %v", updatedTodo, todo)
	}
}

func TestDeleteTodoByID(t *testing.T) {
	ctrl, handlers, mockStorage := setupTest(t)
	defer ctrl.Finish()

	mockStorage.EXPECT().DeleteTodoByID(1).Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, "/todos/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.DeleteTodoByID)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}
}
