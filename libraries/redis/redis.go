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
	password := os.Getenv("REDIS_PASSWORD")

	client := redis.NewClient(&redis.Options{
		Addr: host,
		Username: userName,
		Password: password,
	})

	return client
}

func SetItem(client *redis.Client, email string, uuid string) error {
	ctx := context.Background()

	item := types.AccountItem{
		Uuid: uuid,
	}

	err := client.Set(ctx, email, item, 60*10*time.Second).Err()

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