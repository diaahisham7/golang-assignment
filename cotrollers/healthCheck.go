package cotrollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(context *gin.Context) {
	context.JSON(
		http.StatusOK,
		gin.H{"pong": "pong"},
	)
}

// Echo ...
func Echo(context *gin.Context) {
	var receivedObject interface{}
	if err := context.ShouldBindJSON(&receivedObject); err != nil {
		message := "Invalid Body. Error: " + err.Error()
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": message},
		)
		return
	}
	context.JSON(200, receivedObject)

}
