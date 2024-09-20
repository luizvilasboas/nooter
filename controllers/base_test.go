package controllers

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/olooeez/nooter/models"
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
	db.AutoMigrate(&models.Category{})

	db.Create(&models.Category{Name: "Default"})
	DB = db

	router.POST("/api/v1//notes", CreateNote)
	router.GET("/api/v1//notes", GetAllNotes)
	router.GET("/api/v1//notes/:id", GetNoteByID)
	router.PUT("/api/v1//notes/:id", UpdateNote)
	router.DELETE("/api/v1/notes/:id", DeleteNote)

	router.POST("/api/v1/categories", CreateCategory)
	router.GET("/api/v1/categories", GetAllCategories)
	router.GET("/api/v1/categories/:id", GetCategoryByID)
	router.PUT("/api/v1/categories/:id", UpdateCategory)
	router.DELETE("/api/v1/categories/:id", DeleteCategory)
}
