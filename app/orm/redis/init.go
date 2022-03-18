package redis 

import (
	"test_sql/infra"
)

func InitRedisInstances() {
	Token = &jwtToken{infra.Client}
}