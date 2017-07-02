package test

import "testing"
import "fmt"
import "strconv"
import "os"
import "io"
import "io/ioutil"
import "../redisConnection"
import "../generatePrimes"
import "../processPrimes"

func TestGenerateAndStore(t *testing.T) {
	fmt.Println("Testing GenerateAndStore..")
	ret := generatePrimes.GenerateAndStore(200)
	if ret != 0 {
		t.Error("Failed")
	}
	client := redisConnection.GetRedisClient()
	key := redisConnection.PrimesKey

	count := 0
	input := []int{2, 113, 100, 200, 199, 44444}
	expected := []int64{1, 1, 0, 0, 1, 0}
	result := []int64{-1, -1, -1, -1, -1, -1}
	for _, i := range input {
		val, _ := client.GetBit(key, int64(i)).Result()
		result[count] = val
		count++
	}
	count = 0
	for _, _ = range result {
		fmt.Println("result[count] ", count, result[count])
		if result[count] != expected[count] {
			t.Error("Unexpected value")
		}
		count++
	}
	fmt.Println(" Expected", expected, "Received", result)
}

func TestSumAndMean(t *testing.T) {
	fmt.Println("Testing Sum and Mean..")
	ret := generatePrimes.GenerateAndStore(200)
	if ret != 0 {
		t.Error("Failed")
	}

	input := []int{3, 7, 9, 40, 5, 300}
	expected := []int64{15, 5, 180, 22, 0, 0}
	result := []int64{-1, -1, -1, -1, -1, -1}
	i := 0

	for {
		if i >= len(input) {
			break
		}
		in, err := ioutil.TempFile("", "")
		if err != nil {
			t.Fatal(err)
		}

		_, err = io.WriteString(in, strconv.Itoa(input[i])+"\n"+strconv.Itoa(input[i+1])+"\n")
		if err != nil {
			t.Fatal(err)
		}

		_, err = in.Seek(0, os.SEEK_SET)
		if err != nil {
			t.Fatal(err)
		}

		ret, result[i], result[i+1] = processPrimes.ComputeSumAndMean(in)
		if ret != 0 {
			t.Error("Failed")
		}
		i = i + 2
		in.Close()
	}

	count := 0
	for _, _ = range result {
		fmt.Println("result[count] ", count, result[count])
		if result[count] != expected[count] {
			t.Error("Unexpected result")
		}
		count++
	}
	fmt.Println(" Expected", expected, "Received", result)
}
