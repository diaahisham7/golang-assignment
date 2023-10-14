package main

import (
	"golang-assignment/database"
	"golang-assignment/models"
	"testing"
)

var (
	testSqlDb = database.SqlConnInfo{
		User:     "sql11653192",
		Password: "l4nvArV7QH",
		Host:     "sql11.freesqldatabase.com",
		Port:     "3306",
		Dbname:   "sql11653192",
	}
)

func TestSQLDBConnection(t *testing.T) {
	err := database.InitSqlDB(testSqlDb)
	if err != nil {
		t.Fatalf("connection db error:%s ", err.Error())
	}
}

func TestSqlDbCrud(t *testing.T) {
	err := database.InitSqlDB(testSqlDb)
	if err != nil {
		t.Fatalf("connection db error:%s ", err.Error())
	}
	pet, err := database.PetsTable.Create(models.PetModel{
		Name:  "testPet",
		Age:   5,
		Breed: "cat",
	})
	if err != nil {
		t.Fatalf("error creating pets table: %s", err.Error())
		return
	}
	if pet.Age != 5 || pet.Name != "testPet" || pet.Breed != "cat" {
		_ = database.PetsTable.DeleteById(pet.Id)
		t.Fatalf("error creating pets table: %s", "wrong created data")
		return
	}

	fetchedPet, err := database.PetsTable.GetById(pet.Id)
	if err != nil {
		t.Fatalf("error get by id from pets table: %s", err.Error())
		return
	}
	if fetchedPet.Age != 5 || fetchedPet.Name != "testPet" || fetchedPet.Breed != "cat" {
		_ = database.PetsTable.DeleteById(pet.Id)
		t.Fatalf("error get by id pets table: %s", "wrong created data")
		return
	}

	updatedPet, err := database.PetsTable.UpdateById(pet.Id, models.PetModel{
		Name:  "bella",
		Age:   3,
		Breed: "tiger",
	})
	if err != nil {
		t.Fatalf("error update by id from pets table: %s", err.Error())
		return
	}
	if updatedPet.Age != 3 || updatedPet.Name != "bella" || updatedPet.Breed != "tiger" {
		_ = database.PetsTable.DeleteById(pet.Id)
		t.Fatalf("error update by id pets table: %s", "wrong updated data")
		return
	}

	fetchedPet, err = database.PetsTable.GetById(pet.Id)
	if err != nil {
		t.Fatalf("error get by id from pets table: %s", err.Error())
		return
	}
	if fetchedPet.Age != 3 || fetchedPet.Name != "bella" || fetchedPet.Breed != "tiger" {
		_ = database.PetsTable.DeleteById(pet.Id)
		t.Fatalf("error get by id pets table: %s", "wrong updated data")
		return
	}

	err = database.PetsTable.DeleteById(pet.Id)
	if err != nil {
		t.Fatalf("error delete by id from pets table: %s", err.Error())
		return
	}

}
