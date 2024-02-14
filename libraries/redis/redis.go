package redis

import (
	"context"
	"os"

	"github.com/dongquinn/tech_news_back_go/types"
	"github.com/redis/go-redis/v9"
)

func RedisClient() *redis.Client {
	host := os.Getenv("REDIS_HOST")
	password := os.Getenv("REDIS_PASSWORD")

	client := redis.NewClient(&redis.Options{
		Addr: host,
		Password: password,
	})

	return client
}

func SetItem(client *redis.Client, email string, uuid string) error {
	ctx := context.Background()

	item := types.AccountItem{
		Uuid: uuid,
	}

	err := client.Set(ctx, email, item, 10).Err()

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