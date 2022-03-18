package infra 


import (
	"fmt"

	"github.com/go-redis/redis/v7"
	"test_sql/config"
)

var Client *redis.Client

func InitRedis(config config.Config) {
	Client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port), //redis port
	})

	_, err := Client.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func CloseRedis() {
	if Client != nil {
		if err := Client.Close(); err != nil {
			fmt.Println("[ERROR] Cannot close Redis connection, err:", err)
		}
	}
}