package application

import (
	"github.com/gin-gonic/gin"
	"golang-assignment/cotrollers"
)

func mapUrl(server *gin.Engine) {
	healthCheckGroup := server.Group("", CORSMiddleware())
	healthCheckGroup.POST("/ping", cotrollers.Ping)
	healthCheckGroup.POST("/echo", cotrollers.Echo)

	petsGroup := server.Group("/pets", CORSMiddleware())

	petsGroup.GET("/:id", cotrollers.GetPetByID)
	petsGroup.GET("", cotrollers.GetPets)
	petsGroup.POST("", cotrollers.AddPet)
	petsGroup.PUT("/:id", cotrollers.EditPet)
	petsGroup.DELETE("/:id", cotrollers.DeletePets)

}
