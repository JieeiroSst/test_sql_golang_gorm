package main

import (
	"os"
	"test_sql/app/handler"
	"test_sql/app/orm"
	"test_sql/app/orm/redis"
	"test_sql/config"
	"test_sql/infra"
)

func main() {
	defer close()

	engine := handler.InitEngine()

	configs, err := config.LoadConfig("config")
	if err != nil { // Handle errors reading the config file
		panic(err)
	}

	infra.InitPostgreSQL(configs)
	infra.InitRedis(configs)
	orm.InitOrmInstances()
	redis.InitRedisInstances()

	infra.InitLogging()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	engine.Run(":" + port)
}

func close() {
	infra.ClosePostgreSql()
	infra.CloseRedis()
}
