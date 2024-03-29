package configs

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Configure postgresql connection
func ConfigurePostgreSQLConnection(host, user, password, database, port string) *gorm.DB {
	dsn := fmt.Sprintf("host=%+v user=%+v password=%+v dbname=%+v port=%+v", host, user, password, database, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		panic(err)
	}

	return db
}
