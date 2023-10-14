package database

import "os"

var (
	PetsTable PetsInterface
)

func InitDB() error {
	err := InitSqlDB(SqlConnInfo{
		User:     os.Getenv("user"),
		Password: os.Getenv("password"),
		Host:     os.Getenv("host"),
		Port:     os.Getenv("port"),
		Dbname:   os.Getenv("dbname"),
	})
	return err
}
