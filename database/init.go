package database

var (
	PetsTable PetsInterface
)

func InitDB() error {
	err := initSqlLiteDB()
	return err
}
