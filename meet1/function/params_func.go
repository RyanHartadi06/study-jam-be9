package main

import (
	"fmt"
	"strings"
)

func filter(data []string, condition func(string) bool) []string {
	var result = make([]string, 0, len(data))

	for _, val := range data {
		if ok := condition(val); ok {
			result = append(result, val)
		}
	}
	return result
}

func main() {
	var data = []string{"Hello", "World", "John", "Heru", "Wick"}

	var filterContains0 = filter(data, func(text string) bool {
		return strings.Contains(text, "o")
	})

	fmt.Println("Res filterContains0 : ", filterContains0)

	var filterLength5 = filter(data, func(text string) bool {
		return len(text) == 5
	})

	fmt.Println("filterLength 5: ", filterLength5)
}
