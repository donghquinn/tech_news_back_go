package database

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/dongquinn/tech_news_back_go/configs"
	types "github.com/dongquinn/tech_news_back_go/types/user"
	redis "github.com/redis/go-redis/v9"
)

var redisConfig = configs.RedisConfig

func RedisInstance() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr: redisConfig.Addr,
		Username: redisConfig.User,
		Password: redisConfig.Password,
	})

	return redisClient
}

func RedisSet(redis *redis.Client, key string, value string) error {
	var ctx = context.Background()

	err := redis.Set(ctx, key, value, time.Hour * 3).Err()
 
    if err != nil {
		log.Printf("[REDIS] Set Value Error: %v", err)
        return err
    }	

	return nil
}

func RedisLoginSet(redis *redis.Client, sessionId string,  email string, name string, userId string) error {
	sessionInfo := types.LoginRedisStruct {
		Email: email,
		Name: name,
		UserId: userId}

	var ctx = context.Background()

	err := redis.Set(ctx, sessionId, sessionInfo, time.Hour * 3).Err()
 
    if err != nil {
		log.Printf("[REDIS] Set Value Error: %v", err)
        return err
    }	

	return nil
}

func RedisLoginGet(redis *redis.Client, key string) (types.LoginRedisStruct, error) {
	var ctx = context.Background()

	var loginData types.LoginRedisStruct

	value, getErr := redis.Get(ctx, key).Result()

    if getErr != nil {
		log.Printf("[REDIS] Get Value Error: %v", getErr)
        return types.LoginRedisStruct{}, getErr
    }

	json.Unmarshal([]byte(value), &loginData)
	
	return loginData, nil
}

func RedisGet(redis *redis.Client, key string) (string, error) {
	var ctx = context.Background()

	value, getErr := redis.Get(ctx, key).Result()

    if getErr != nil {
		log.Printf("[REDIS] Get Value Error: %v", getErr)
        return "", getErr
    }	

	return value, nil
}