package database

import "golang-assignment/models"

type PetsInterface interface {
	GetById(id int) (models.PetModel, error)
	GetAll(pageSize int, pageNo int) ([]models.PetModel, error)
	Create(pet models.PetModel) (models.PetModel, error)
	UpdateById(id int, pet models.PetModel) (models.PetModel, error)
	DeleteById(id int) error
}
