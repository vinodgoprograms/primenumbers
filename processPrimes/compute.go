/* 
Package used to process the generated prime numbers
*/
package processPrimes

import (
	"github.com/vinodgoprograms/primenumbers/redisConnection"
	"fmt"
	"os"
)
//Internal function
//reads two integers and process the prime numbers fall between them
func readInput(inStream *os.File) (int, int) {

	if inStream == nil {
		inStream = os.Stdin
	}
	var low int
	var high int
	fmt.Print("Enter a lower bound: ")
	fmt.Fscanf(inStream, "%d", &low)
	fmt.Print("Enter an upper bound: ")
	fmt.Fscanf(inStream, "%d", &high)
	//fmt.Println(low, high)
	return low, high

}
/*
Function to compute sum and mean of all the prime numbers between two integers.

Input:
 inStream - input stream to read the inputs from.
Output:
 result: 0 (Success) -1 (Failure)
 sum: sum of all prime numbers between the two integers read from input
 mean: mean of all prime numbers between the two integers read from input
*/
func ComputeSumAndMean(inStream *os.File) (int, int64, int64) {

	var sum int64
	var mean int64
        result := make([]int, 0)

	key := redisConnection.PrimesKey
	client := redisConnection.GetRedisClient()
	Max, err := client.Get(redisConnection.PrimesMax).Int64()
	if err != nil {
		fmt.Println("Error: Retrieving the prime max")
		return -1, 0, 0
	}
	for {
		low, high := readInput(inStream)

		//sum and mean will return previously values
		//this is to make the test cases work
		if low <= 0 || low > high {
                        fmt.Println("Error: invalid low & high")
			return 0, sum, mean
		}
		if int64(low) >= Max || int64(high) >= Max {
                        fmt.Println("Error: invalid range")
			return 0, 0, 0
		}

		sum = 0
		count := 0

		for i := low; i <= high; i++ {
			val, _ := client.GetBit(key, int64(i)).Result()
			if val == 1 {
				count++
                                result = append(result, i)
				sum = sum + int64(i)
			}
		}
		fmt.Println("Result:")
		fmt.Println("Prime numbers:", result)
		fmt.Println("Sum:", sum)
		mean = sum / int64(count)
		fmt.Println("Mean:", mean)
	}
	return 0, sum, mean

}
