package routes

import (
	"gitlab.com/olooeez/nooter/controllers"

	"github.com/gin-gonic/gin"
)

func NoteRoutes(router *gin.Engine) {
	noteGroup := router.Group("/api/notes")
	{
		noteGroup.GET("/", controllers.GetAllNotes)
		noteGroup.GET("/:id", controllers.GetNoteByID)
		noteGroup.POST("/", controllers.CreateNote)
		noteGroup.PUT("/:id", controllers.UpdateNote)
		noteGroup.DELETE("/:id", controllers.DeleteNote)
	}
}
