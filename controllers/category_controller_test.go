package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"gitlab.com/olooeez/nooter/models"
)

func TestCreateCategory(t *testing.T) {
	setupRouter()

	category := models.CategoryCreate{Name: "Nova Categoria"}
	jsonValue, _ := json.Marshal(category)

	req, _ := http.NewRequest("POST", "/api/v1/categories", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}
}

func TestGetAllCategories(t *testing.T) {
	setupRouter()

	req, _ := http.NewRequest("GET", "/api/v1/categories", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestGetCategoryByID(t *testing.T) {
	setupRouter()

	category := models.Category{Name: "Categoria de Teste"}
	DB.Create(&category)

	req, _ := http.NewRequest("GET", "/api/v1/categories/"+strconv.Itoa(int(category.ID)), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestUpdateCategory(t *testing.T) {
	setupRouter()

	category := models.Category{Name: "Categoria para Atualizar"}
	DB.Create(&category)

	updatedCategory := models.CategoryCreate{Name: "Categoria Atualizada"}
	jsonValue, _ := json.Marshal(updatedCategory)

	req, _ := http.NewRequest("PUT", "/api/v1/categories/"+strconv.Itoa(int(category.ID)), bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestDeleteCategory(t *testing.T) {
	setupRouter()

	category := models.Category{Name: "Categoria para Deletar"}
	DB.Create(&category)

	req, _ := http.NewRequest("DELETE", "/api/v1/categories/"+strconv.Itoa(int(category.ID)), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	req, _ = http.NewRequest("DELETE", "/api/v1/categories/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusForbidden {
		t.Errorf("Expected status code %d, got %d", http.StatusForbidden, w.Code)
	}
}
