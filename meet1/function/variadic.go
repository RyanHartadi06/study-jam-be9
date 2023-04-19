package main

import "fmt"

func main() {
	fmt.Println(" SUM 1, 3, 5 = ", sum(1, 3, 5))

}

func sum(numbers ...uint64) (result uint64) {
	for _, num := range numbers {
		result += num
	}
	return result
}
