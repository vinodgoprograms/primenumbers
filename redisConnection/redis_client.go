package redisConnection

import (
	"github.com/go-redis/redis"
)

var PrimesKey string = "primes"
var PrimesMax string = "max"

func GetRedisClient() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       13})
	return client

}
