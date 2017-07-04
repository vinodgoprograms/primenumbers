/* Main package
   
   Generate prime numbers and computes sum and mean, based on supplied inputs. Please refer to README.md for details.
*/
package main

import (
	"github.com/vinodgoprograms/primenumbers/generatePrimes"
	"github.com/vinodgoprograms/primenumbers/processPrimes"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("usage: %s <primeUnder>\n", os.Args[0])
		os.Exit(2)
	}

	primeUnder, err := strconv.Atoi(os.Args[1:][0])

	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	if primeUnder <= 1 || primeUnder > math.MaxInt32 {
		fmt.Println("Out of range input: (valid values include: 2-2147483647")
		os.Exit(1)
	}

	ret := generatePrimes.GenerateAndStore(primeUnder)
	if ret != 0 {
		fmt.Printf("Error: Failed to generate and store %d\n", ret)
		os.Exit(1)
	}
	ret, _, _ = processPrimes.ComputeSumAndMean(nil)
	if ret != 0 {
		fmt.Printf("Error: Failed to process prime numbers %d\n", ret)
		os.Exit(1)
	}

	os.Exit(0)
}
