/* This package can be used to connect to a local redis server.
   Port = 6379
   Password = ""
   DB = 13
*/
package redisConnection

import (
	"github.com/go-redis/redis"
)
//redis key for the generated prime numbers
var PrimesKey string = "primes"
//redis key for the integer used to generate the prime numbers
var PrimesMax string = "max"

/* 
   Connects to local redis server and returns a redis.Client object.
   Refer to the package information for connection parameters.

   Input:
      - None
   Output:
      - instance of redis.Client
*/
func GetRedisClient() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       13})
	return client

}
