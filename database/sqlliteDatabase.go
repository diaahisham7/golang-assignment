package database

import "golang-assignment/models"

func initSqlLiteDB() error {
	/*
		init gorm sqlite
	*/
	PetsTable = petsSqlliteTable{
		/*
			DB : db
		*/
	}

	return nil
}

//region sqllite models

type petsSqlliteModel struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Breed string `json:"breed"`
}

//endregion

//region petsSqlliteTable
type petsSqlliteTable struct {
	/*
		DB database
	*/
}

func (p petsSqlliteTable) GetById(id int) (models.PetModel, error) {
	//TODO implement me
	panic("implement me")
}

func (p petsSqlliteTable) GetAll(pageSize int, pageNo int) ([]models.PetModel, error) {
	//TODO implement me
	panic("implement me")
}

func (p petsSqlliteTable) Create(pet models.PetModel) (models.PetModel, error) {
	//TODO implement me
	panic("implement me")
}

func (p petsSqlliteTable) UpdateById(id int, pet models.PetModel) (models.PetModel, error) {
	//TODO implement me
	panic("implement me")
}

func (p petsSqlliteTable) DeleteById(id int) error {
	//TODO implement me
	panic("implement me")
}

//endregion
