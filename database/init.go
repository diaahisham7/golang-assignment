package database

var (
	PetsTable PetsInterface
)

func InitDB() error {
	err := initSqlDB()
	return err
}
