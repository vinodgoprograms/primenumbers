# primenumbers
generate all the prime numbers between 1 to X where X is a command line argument to the application.

The prime numbers should be stored in a local Redis instance
Once the prime numbers are generated the application should repeatedly ask the user for a lower and upper bounds (inclusive) on the prime numbers to return along with their sum and mean

# Example

 - $ Enter a lower bound: 3
 - $ Enter an upper bound: 9
 - $ Result:
 - $ Prime numbers: [3, 5, 7]
 - $ Sum: 15
 - $ Mean: 5

Include basic unit testing around the core functionality
All code (excluding any external dependencies) should be committed to a GitHub repository.

# Dependencies 

- import "github.com/go-redis/redis"
- go version go1.8.3 darwin/amd64

# Installation

go get -u github.com/vinodgoprograms/primenumbers



