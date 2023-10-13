package application

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang-assignment/database"
)

// StartApplication ... start the server and mapurl
func StartApplication() {
	server := gin.Default()
	mapUrl(server)

	err := database.InitDB()
	if err != nil {
		log.Fatal("error init database: ", err.Error())
		return
	}
	_ = server.Run(":5050")
}
