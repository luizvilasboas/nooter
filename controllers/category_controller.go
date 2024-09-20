package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/olooeez/nooter/models"
)

// CreateCategory godoc
// @Summary Criar Categoria
// @Description Cria uma nova categoria
// @Tags categories
// @Accept json
// @Produce json
// @Param category body models.CategoryCreate true "Categoria"
// @Success 201 {object} models.Category
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /categories [post]
func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Dados inválidos"})
		return
	}

	category.Notes = []models.Note{}

	if err := DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Erro ao criar categoria"})
		return
	}

	c.JSON(http.StatusCreated, category)
}

// GetAllCategories godoc
// @Summary Obter Todas as Categorias
// @Description Retorna todas as categorias
// @Tags categories
// @Produce json
// @Success 200 {array} models.Category
// @Failure 500 {object} ErrorResponse
// @Router /categories [get]
func GetAllCategories(c *gin.Context) {
	var categories []models.Category
	if err := DB.Preload("Notes").Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Erro ao buscar categorias"})
		return
	}

	c.JSON(http.StatusOK, categories)
}

// GetCategoryByID godoc
// @Summary Obter Categoria por ID
// @Description Retorna uma categoria específica pelo ID
// @Tags categories
// @Produce json
// @Param id path int true "ID da Categoria"
// @Success 200 {object} models.Category
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /categories/{id} [get]
func GetCategoryByID(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	if err := DB.Preload("Notes").First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "Categoria não encontrada"})
		return
	}

	c.JSON(http.StatusOK, category)
}

// UpdateCategory godoc
// @Summary Atualizar Categoria
// @Description Atualiza uma categoria existente
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "ID da Categoria"
// @Param category body models.CategoryCreate true "Categoria"
// @Success 200 {object} models.Category
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /categories/{id} [put]
func UpdateCategory(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	if err := DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "Categoria não encontrada"})
		return
	}

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Dados inválidos"})
		return
	}

	if err := DB.Save(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Não foi possível atualizar a categoria"})
		return
	}

	c.JSON(http.StatusOK, category)
}

// DeleteCategory godoc
// @Summary Deletar Categoria
// @Description Deleta uma categoria pelo ID
// @Tags categories
// @Param id path int true "ID da Categoria"
// @Success 200 {object} SuccessResponse
// @Failure 404 {object} ErrorResponse
// @Router /categories/{id} [delete]
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if id == 1 {
		c.JSON(http.StatusForbidden, ErrorResponse{Message: "Não pode deletar categoria default"})
	}

	var category models.Category
	if err := DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "Categoria não encontrada"})
		return
	}

	if err := DB.Delete(category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "Não foi possível deletar a categoria"})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "Categoria deletada com sucesso"})
}
