package store

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

type StorageService struct {
	redisClient *redis.Client
}

var storeService = &StorageService{}

const CacheDuration = 6 * time.Hour

func InitializeStore() *StorageService {

	godotenv.Load(".env")
	REDIS_ADDR := os.Getenv("REDIS_ADDR")
	REDIS_PASSWORD := os.Getenv("REDIS_PASSWORD")
	REDIS_DB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	redisClient := redis.NewClient(&redis.Options{
		Addr:     REDIS_ADDR,
		Password: REDIS_PASSWORD,
		DB:       REDIS_DB,
	})

	pong, err := redisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))

	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}

func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
	err := storeService.redisClient.Set(shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}
}

func GetOriginalUrl(shortUrl string) string {
	result, err := storeService.redisClient.Get(shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed GetOriginalUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return result
}
