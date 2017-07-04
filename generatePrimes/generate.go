/*Generate prime numbers and store them in redis.
The generated prime numbers are stored under the key  "max"
and the prime numbers are stored under the key "primes".
*/
package generatePrimes

import (
	"github.com/vinodgoprograms/primenumbers/redisConnection"
	"fmt"
)

/*This external function is used to generate prime numbers under the 
given number. The generated values will be stored in a redis db

Input: Integer
Output:
  0 - SUCCESS
 -1 - ERROR */
func GenerateAndStore(primeUnder int) int {

	primesList := generatePrimes(primeUnder)
	if len(primesList) == 0 {
		fmt.Println("Error: Failed to generate prime numbers")
		return -1
	}
	err := storeInRedis(primeUnder, primesList)
	if err == 0 {
		fmt.Println("OK: Stored in redis")
		return 0
	}
	fmt.Println("Error: Failed to store in redis")
	return -1
}
//internal function 
//used to generate prime numbers
//Input: Integer
//Returns: array of prime numbers
func generatePrimes(primeUnder int) []int {
	result := make([]int, 0)

	for i := 2; i <= primeUnder; i++ {
		prime := 1
		for j := 2; j < i/2+1; j++ {
			if i%j == 0 {
				prime = 0
				break
			}
		}
		if prime == 1 {
			result = append(result, i)
		}
	}
	//TODO: make it fit the buffer width
	fmt.Println("Primes:", result)
	return result

}
//internal function 
//used to store prime numbers and max integer in redis
//Input: Integer, Array of integers
//Returns: 
//0 - On Success
//-1 - On Failure
func storeInRedis(primeUnder int, primesList []int) int {

	client := redisConnection.GetRedisClient()
	if client == nil {
		fmt.Println("Error: failed to create the client")
		return -1
	}

	maxKey := redisConnection.PrimesMax
	primesKey := redisConnection.PrimesKey
	client.Set(maxKey, primeUnder, 0)

	for i := 0; i < primeUnder; i++ {
		client.SetBit(primesKey, int64(i), 0)
	}
	for i := 0; i < len(primesList); i++ {
		//fmt.Println(" primesList[i] ", primesList[i])
		client.SetBit(primesKey, int64(primesList[i]), 1)
	}
	return 0
}
