package database

import (
	"errors"
	"fmt"
	"golang-assignment/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type petsSqlModel struct {
	gorm.Model
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Breed string `json:"breed"`
}

type SqlConnInfo struct {
	User     string
	Password string
	Host     string
	Port     string
	Dbname   string
}

func InitSqlDB(connInfo SqlConnInfo) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		connInfo.User,
		connInfo.Password,
		connInfo.Host,
		connInfo.Port,
		connInfo.Dbname,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.New(fmt.Sprintf("error connecting to database: %s", "code:1"))
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
		return models.PetModel{}, errors.New("error getting data: code:2")
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
		return []models.PetModel{}, errors.New("error getting data: code:3")
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
		return models.PetModel{}, errors.New("error adding data: code:4")
	}
	pet.Id = int(petCreated.ID)
	return pet, nil
}

func (p petsSqlTable) UpdateById(id int, pet models.PetModel) (models.PetModel, error) {
	var fetchedPet petsSqlModel
	e := p.DB.First(&fetchedPet, id)
	if e.Error != nil {
		return models.PetModel{}, errors.New("error getting data: code:5")
	}
	fetchedPet.Name = pet.Name
	fetchedPet.Age = pet.Age
	fetchedPet.Breed = pet.Breed
	e = p.DB.Save(&fetchedPet)
	if e.Error != nil {
		return models.PetModel{}, errors.New("error updating data: code:6")
	}
	pet.Id = int(fetchedPet.ID)
	return pet, nil
}

func (p petsSqlTable) DeleteById(id int) error {
	var fetchedPet petsSqlModel
	e := p.DB.First(&fetchedPet, id)
	if e.Error != nil {
		return errors.New("error getting data: code:7")
	}
	e = p.DB.Delete(&fetchedPet)
	if e.Error != nil {
		return errors.New("error deleting data: code:8")
	}
	return nil
}

//endregion
