package database

import (
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

type redisDatabase struct {
	client *redis.Client
}

func createRedisDatabase() (Database, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:32768",
		Password: "",
		DB:       0,
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, &CreateDatabaseError{}
	}

	return &redisDatabase{client: client}, nil
}

func (r *redisDatabase) Set(key, value string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err := r.client.Set(ctx, key, value, 0).Err()
	if err != nil {
		return generateError("set", err)
	}
	return key, nil
}

func (r *redisDatabase) Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	value, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return generateError("get", err)
	}

	return value, nil
}

func (r *redisDatabase) Delete(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err := r.client.Del(ctx, key).Err()
	if err != nil {
		return generateError("delete", err)
	}
	return key, nil
}

func generateError(operation string, err error) (string, error) {
	log.Print(err)
	if err == redis.Nil {
		return "", &OperationError{operation}
	}
	return "", &DownError{}
}
