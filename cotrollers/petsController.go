package cotrollers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang-assignment/domain/pets"
	"golang-assignment/models"
	"net/http"
	"strconv"
)

func GetPetByID(context *gin.Context) {
	id, err := strconv.ParseFloat(context.Param("id"), 64)
	if err != nil {
		log.Error("id param error: ", err.Error())
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	result, err := pets.GetById(int(id))
	if err != nil {
		log.Error(err.Error())
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	context.JSON(http.StatusOK, result)
}

func GetPets(context *gin.Context) {
	pageNo, err := strconv.ParseFloat(context.Query("pageno"), 64)
	if err != nil {
		log.Error("pageno Query error: ", err.Error())
		pageNo = 0
	}
	pageSize, err := strconv.ParseFloat(context.Query("pagesize"), 64)
	if err != nil {
		log.Error("pagesize Query error: ", err.Error())
		pageSize = 0
	}

	result, err := pets.GetAll(int(pageSize), int(pageNo))

	if err != nil {
		log.Error(err.Error())
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, result)
}

func AddPet(context *gin.Context) {
	pet := models.PetModel{}
	err := context.ShouldBindJSON(&pet)
	if err != nil {
		log.Error("error in request body: ", err.Error())
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	result, err := pets.Create(pet)
	if err != nil {
		log.Error(err.Error())
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	context.JSON(http.StatusOK, result)

}

func EditPet(context *gin.Context) {
	id, err := strconv.ParseFloat(context.Param("id"), 64)
	if err != nil {
		log.Error("id param error: ", err.Error())
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	pet := models.PetModel{}
	err = context.ShouldBindJSON(&pet)
	if err != nil {
		log.Error("error in request body: ", err.Error())
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	result, err := pets.UpdateById(int(id), pet)
	if err != nil {
		log.Error(err.Error())
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	context.JSON(http.StatusOK, result)
}

func DeletePets(context *gin.Context) {
	id, err := strconv.ParseFloat(context.Param("id"), 64)
	if err != nil {
		log.Error("id param error: ", err.Error())
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err = pets.Delete(int(id))
	if err != nil {
		log.Error(err.Error())
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	context.JSON(http.StatusOK, "Deleted successfully")
}
