package redis

import (
	"context"
	"os"
	"time"

	"github.com/dongquinn/tech_news_back_go/types"
	"github.com/redis/go-redis/v9"
)

func RedisClient() *redis.Client {
	host := os.Getenv("REDIS_HOST")
	userName := os.Getenv("REDIS_USER")
	password := os.Getenv("REDIS_PASS")

	client := redis.NewClient(&redis.Options{
		Addr: host,
		Username: userName,
		Password: password,
	})

	return client
}

func SetItem(client *redis.Client, email string, uuid string, minuteDuration int) error {
	ctx := context.Background()

	item := types.AccountItem{
		Uuid: uuid,
	}

	expireDuration := minuteDuration * 60 * 1000
	err := client.Set(ctx, email, item, time.Duration(expireDuration)).Err()

	if err != nil {
		return err
	}

	return nil
}

func GetItem(client *redis.Client, email string) (string, error) {
	ctx := context.Background()

	uuid, err := client.Get(ctx, email).Result()

	if err != nil {
		return "", err
	}

	return uuid, nil
}

func DeleteItem(client *redis.Client, email string) error {
	ctx := context.Background()

	err := client.Del(ctx, email).Err()

	if err != nil {
		return err
	}

	return nil
}