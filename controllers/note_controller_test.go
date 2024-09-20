package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"gitlab.com/olooeez/nooter/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var router *gin.Engine
var db *gorm.DB

func setupRouter() {
	gin.SetMode(gin.TestMode)
	router = gin.Default()

	var err error
	db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&models.Note{})
	DB = db

	router.POST("/notes", CreateNote)
	router.GET("/notes", GetAllNotes)
	router.GET("/notes/:id", GetNoteByID)
	router.PUT("/notes/:id", UpdateNote)
	router.DELETE("/notes/:id", DeleteNote)
}

func TestCreateNote(t *testing.T) {
	setupRouter()

	note := `{"title": "Test Note", "content": "This is a test note"}`
	req, _ := http.NewRequest("POST", "/notes", bytes.NewReader([]byte(note)))
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

	req, _ := http.NewRequest("GET", "/notes", nil)
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

	req, _ := http.NewRequest("GET", "/notes/"+strconv.Itoa(int(note.ID)), nil)
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

	updatedNote := `{"title": "Updated Note", "content": "This is an updated test note"}`
	req, _ := http.NewRequest("PUT", "/notes/"+strconv.Itoa(int(note.ID)), bytes.NewReader([]byte(updatedNote)))
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

	req, _ := http.NewRequest("DELETE", "/notes/"+strconv.Itoa(int(note.ID)), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", w.Code)
	}
}
