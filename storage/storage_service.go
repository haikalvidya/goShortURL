// wrapper for redis
package storage

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// declaration for the storage service and redis
var (
	storageService = &StorageService{}
	ctx = context.Background()
)

// define struct of wrapper for raw redis client
type StorageService struct {
	redisClient *redis.Client
}

// set timer for cache redis duration
const CacheDuration = 6 * time.Hour

// init the storage service and return a store pointer
func InitializeStorage() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:		"localhost:6379",
		Password:	"",
		DB:			0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("Redis started successfully: ping message = {%s}", pong)
	storageService.redisClient = redisClient
	return storageService
}


// maping for original URL and generated short url
func SavedURLMapping(shortUrl string, originalUrl string, userId string){
	err := storageService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("failed to saving key url, the error: %v - short url: %s - original url: %s\n", err, shortUrl, originalUrl))
	}

	fmt.Printf("Saved shortUrl: %s - originalUrl: %s\n", shortUrl, originalUrl)
}

// for retrueve the initial long URL once the short url is provided
// or when user calling the short url to retrieve the original url
func RetrieveInitialURL(shortUrl string) string {
	result, err := storageService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("failed to retrive url | error: %v - short url: %s\n", err, shortUrl))
	}
	return result
}