package infra_test

import (
	"test_sql/config"
	"test_sql/infra"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitDB(t *testing.T) {
	postgres := config.PostgreSQL{
		Username: "postgres",
		Password: "postgres",
		Host:     "localhost",
		Name:     "postgres",
		Port:     "5432",
		Debug:    false,
		SSL:      "disable",
	}
	configs := config.Config{
		PostgreSQL: postgres,
	}

	err := infra.InitPostgreSQL(configs)
	assert.NoError(t, err)
}
