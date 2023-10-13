package database

import (
	"errors"
	"fmt"
	"golang-assignment/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type petsSqlModel struct {
	gorm.Model
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Breed string `json:"breed"`
}

func initSqlDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("user"),
		os.Getenv("password"),
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("dbname"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.New(fmt.Sprintf("error connecting to database: %s", err.Error()))
	}
	db.AutoMigrate(&petsSqlModel{})
	PetsTable = petsSqlTable{
		DB: db,
	}
	return nil
}

//region sql models

//endregion

// region petsSqlTable
type petsSqlTable struct {
	DB *gorm.DB
}

func (p petsSqlTable) GetById(id int) (models.PetModel, error) {
	var result petsSqlModel
	e := p.DB.First(&result, id)
	if e.Error != nil {
		return models.PetModel{}, e.Error
	}
	return models.PetModel{
		Id:    int(result.ID),
		Name:  result.Name,
		Age:   result.Age,
		Breed: result.Breed,
	}, nil
}

func (p petsSqlTable) GetAll(pageSize int, pageNo int) ([]models.PetModel, error) {
	offset := (pageNo - 1) * pageSize
	limit := pageNo * pageSize
	result := make([]petsSqlModel, 0)
	e := p.DB.Model(petsSqlModel{}).Limit(limit).Offset(offset).Find(&result)
	if e.Error != nil {
		return []models.PetModel{}, e.Error
	}
	petsResult := make([]models.PetModel, 0)
	for _, pet := range result {
		petsResult = append(petsResult, models.PetModel{
			Id:    int(pet.ID),
			Name:  pet.Name,
			Age:   pet.Age,
			Breed: pet.Breed,
		})
	}
	return petsResult, nil
}

func (p petsSqlTable) Create(pet models.PetModel) (models.PetModel, error) {
	petCreated := petsSqlModel{
		Name:  pet.Name,
		Age:   pet.Age,
		Breed: pet.Breed,
	}

	e := p.DB.Create(&petCreated)
	if e.Error != nil {
		return models.PetModel{}, e.Error
	}
	pet.Id = int(petCreated.ID)
	return pet, nil
}

func (p petsSqlTable) UpdateById(id int, pet models.PetModel) (models.PetModel, error) {
	var fetchedPet petsSqlModel
	e := p.DB.First(&fetchedPet, id)
	if e.Error != nil {
		return models.PetModel{}, e.Error
	}
	fetchedPet.Name = pet.Name
	fetchedPet.Age = pet.Age
	fetchedPet.Breed = pet.Breed
	e = p.DB.Save(&fetchedPet)
	if e.Error != nil {
		return models.PetModel{}, e.Error
	}
	pet.Id = int(fetchedPet.ID)
	return pet, nil
}

func (p petsSqlTable) DeleteById(id int) error {
	var fetchedPet petsSqlModel
	e := p.DB.First(&fetchedPet, id)
	if e.Error != nil {
		return e.Error
	}
	e = p.DB.Delete(&fetchedPet)
	if e.Error != nil {
		return e.Error
	}
	return nil
}

//endregion
