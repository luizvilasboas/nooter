package controllers

import (
	"net/http"
	"strconv"

	"gitlab.com/olooeez/nooter/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

// GetAllNotes godoc
// @Summary Lista todas as notas
// @Description Obtém uma lista de todas as notas
// @Tags notes
// @Accept json
// @Produce json
// @Success 200 {object} []models.Note
// @Failure 500 {object} ErrorResponse
// @Router /notes [get]
func GetAllNotes(c *gin.Context) {
	var notes []models.Note
	if err := DB.Find(&notes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Erro ao buscar notas"})
		return
	}

	c.JSON(http.StatusOK, notes)
}

// GetNoteByID godoc
// @Summary Obtém uma nota por ID
// @Description Busca uma nota pelo seu ID
// @Tags notes
// @Accept json
// @Produce json
// @Param id path int true "ID da Nota"
// @Success 200 {object} models.Note
// @Failure 404 {object} ErrorResponse
// @Router /notes/{id} [get]
func GetNoteByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var note models.Note
	if err := DB.First(&note, id).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "Nota não encontrada"})
		return
	}

	c.JSON(http.StatusOK, note)
}

// CreateNote godoc
// @Summary Cria uma nova nota
// @Description Adiciona uma nova nota ao banco de dados
// @Tags notes
// @Accept json
// @Produce json
// @Param note body models.NoteCreate true "Nota a ser criada"
// @Success 201 {object} models.Note
// @Failure 400 {object} ErrorResponse
// @Router /notes [post]
func CreateNote(c *gin.Context) {
	var note models.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Dados inválidos"})
		return
	}

	if err := DB.Create(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Erro ao criar nota"})
		return
	}

	c.JSON(http.StatusCreated, note)
}

// UpdateNote godoc
// @Summary Atualiza uma nota
// @Description Atualiza uma nota existente no banco de dados
// @Tags notes
// @Accept json
// @Produce json
// @Param id path int true "ID da Nota"
// @Param note body models.NoteCreate true "Nota a ser atualizada"
// @Success 200 {object} models.Note
// @Failure 404 {object} ErrorResponse
// @Failure 400 {object} ErrorResponse
// @Router /notes/{id} [put]
func UpdateNote(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var note models.Note

	if err := DB.First(&note, id).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "Nota não encontrada"})
		return
	}

	var updatedNote models.Note
	if err := c.ShouldBindJSON(&updatedNote); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Dados inválidos"})
		return
	}

	note.Title = updatedNote.Title
	note.Content = updatedNote.Content
	note.UpdatedAt = updatedNote.UpdatedAt

	if err := DB.Save(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Dados inválidos"})
		return
	}

	c.JSON(http.StatusOK, note)
}

// DeleteNote godoc
// @Summary Deleta uma nota
// @Description Remove uma nota do banco de dados
// @Tags notes
// @Accept json
// @Produce json
// @Param id path int true "ID da Nota"
// @Success 204 {object} SuccessResponse
// @Failure 404 {object} ErrorResponse
// @Router /notes/{id} [delete]
func DeleteNote(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var note models.Note
	if err := DB.First(&note, id).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "Nota não encontrada"})
		return
	}

	if err := DB.Delete(note, id).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "Não foi possível deletar a nota"})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "Nota deletada com sucesso"})
}
