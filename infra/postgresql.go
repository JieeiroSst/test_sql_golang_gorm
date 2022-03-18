package infra

import (
	"fmt"
	"test_sql/config"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PGDB struct {
	*gorm.DB
}

var pgSingleton *PGDB

func InitPostgreSQL(config config.Config) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.PostgreSQL.Host,
		config.PostgreSQL.Username,
		config.PostgreSQL.Password,
		config.PostgreSQL.Name,
		config.PostgreSQL.Port,
		config.PostgreSQL.SSL)

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic(err)
		return err
	}

	db.LogMode(true)

	pgSingleton = &PGDB{db}

	return nil
}

func GetDB() *PGDB {
	if pgSingleton == nil {
		panic("can not connect database")
	}
	return pgSingleton
}

func ClosePostgreSql() error {
	return pgSingleton.Close()
}
