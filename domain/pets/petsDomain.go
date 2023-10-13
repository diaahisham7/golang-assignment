package pets

import (
	"golang-assignment/database"
	"golang-assignment/models"
)

func GetById(id int) (models.PetModel, error) {
	return database.PetsTable.GetById(id)

}

func GetAll(pageSize int, pageNo int) ([]models.PetModel, error) {
	if pageNo < 1 {
		pageNo = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	return database.PetsTable.GetAll(pageSize, pageNo)
}

func Create(pet models.PetModel) (models.PetModel, error) {
	return database.PetsTable.Create(pet)
}

func UpdateById(id int, pet models.PetModel) (models.PetModel, error) {
	return database.PetsTable.UpdateById(id, pet)
}

func Delete(id int) error {
	return database.PetsTable.DeleteById(id)
}
