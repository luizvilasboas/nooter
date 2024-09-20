package routes

import (
	"gitlab.com/olooeez/nooter/controllers"

	"github.com/gin-gonic/gin"
)

func NoteRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		noteGroup := v1.Group("/notes")
		{
			noteGroup.GET("/", controllers.GetAllNotes)
			noteGroup.GET("/:id", controllers.GetNoteByID)
			noteGroup.POST("/", controllers.CreateNote)
			noteGroup.PUT("/:id", controllers.UpdateNote)
			noteGroup.DELETE("/:id", controllers.DeleteNote)
		}

		categories := v1.Group("/categories")
		{
			categories.GET("/", controllers.GetAllCategories)
			categories.GET("/:id", controllers.GetCategoryByID)
			categories.POST("/", controllers.CreateCategory)
			categories.PUT("/:id", controllers.UpdateCategory)
			categories.DELETE("/:id", controllers.DeleteCategory)
		}
	}
}
