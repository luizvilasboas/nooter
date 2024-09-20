package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"gitlab.com/olooeez/nooter/models"
)

func TestCreateNote(t *testing.T) {
	setupRouter()

	note := `{"title": "Test Note", "content": "This is a test note", "category_id": 1}`
	req, _ := http.NewRequest("POST", "/api/v1/notes", bytes.NewReader([]byte(note)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code 201, got %v", w.Code)
	}
}

func TestGetAllNotes(t *testing.T) {
	setupRouter()

	note := models.Note{Title: "Test Note", Content: "This is a test note"}
	db.Create(&note)

	req, _ := http.NewRequest("GET", "/api/v1/notes", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", w.Code)
	}
}

func TestGetNoteByID(t *testing.T) {
	setupRouter()

	note := models.Note{Title: "Test Note", Content: "This is a test note"}
	db.Create(&note)

	req, _ := http.NewRequest("GET", "/api/v1/notes/"+strconv.Itoa(int(note.ID)), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", w.Code)
	}
}

func TestUpdateNote(t *testing.T) {
	setupRouter()

	note := models.Note{Title: "Test Note", Content: "This is a test note"}
	db.Create(&note)

	updatedNote := `{"title": "Updated Note", "content": "This is an updated test note", "category_id": 1}`
	req, _ := http.NewRequest("PUT", "/api/v1/notes/"+strconv.Itoa(int(note.ID)), bytes.NewReader([]byte(updatedNote)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", w.Code)
	}
}

func TestDeleteNote(t *testing.T) {
	setupRouter()

	note := models.Note{Title: "Test Note", Content: "This is a test note"}
	db.Create(&note)

	req, _ := http.NewRequest("DELETE", "/api/v1/notes/"+strconv.Itoa(int(note.ID)), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", w.Code)
	}
}
