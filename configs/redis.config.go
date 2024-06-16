package configs

import "os"

type RedisConfigStruct struct {
	Addr string
	User string
	Password string
}

var RedisConfig RedisConfigStruct

func SetRedisConfig() {
	RedisConfig.Addr = os.Getenv("REDIS_HOST")
	RedisConfig.User = os.Getenv("REDIS_USER")
	RedisConfig.Password = os.Getenv("REDIS_PASSWORD")
}