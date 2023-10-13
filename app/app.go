package application

import (
	"golang-assignment/database"
)

// StartApplication ... start the server and mapurl
func StartApplication() {

	err := database.InitDB()
	if err != nil {

	}

}
