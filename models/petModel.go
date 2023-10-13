package models

type PetModel struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Breed string `json:"breed"`
}
