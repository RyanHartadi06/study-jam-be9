package main

import "fmt"

/*
	closure adalah fungsi yag disimpan dalam variable (anonymus func)
	declaration :
	func main(){
		// fungsi yang disimpan dalam bentuk variable
		var getMaxNumbers = func(numbers []int){
		}

		[] IIFE (Immediately-Invoked Function Expression)
		var maxNumbers = func(numbers []int) int {
		}(numbers)
	}

	// FUngsi juga bisa mengembalikan suatu fungsi

	func findMax(numbers []int, max int) (int, func() []int){

	}
*/

func main() {
	var numbers = []int{2, 3, 0, 4, 3, 2, 0, 3}
	var text = []string{"asd", "qwe"}
	// NOTE : Jadi fungsi dibawah itu variable terdapat func tapi harus assign params
	var getMaxNumbers = func(numbers []int) int {
		if len(numbers) == 0 {
			return 0
		}

		var max int

		for _, val := range numbers {
			if val > max {
				max = val
			}
		}

		return max

	}

	var cekString = func(text []string) []string {

		var result = make([]string, 0, len(text))

		for _, data := range text {
			if data == "asd" {
				result = append(result, "Hello World")
			}
			result = append(result, data)
		}

		return result
	}

	fmt.Println("Fungsi Sendiri Closure", cekString(text))

	fmt.Println("Closure:", getMaxNumbers(numbers))

	// NOTE : Jadi fungsi dibawah itu variable terdapat func tapi tidak assign params ketika print
	// Params langsung di assign ketika dia load di awal
	var maxNumber = func(numbers []int) int {
		var max int

		for _, val := range numbers {
			if val > max {
				max = val
			}
		}

		return max
	}(numbers)

	fmt.Println("IIFE: ", maxNumber)

	var length, details = findMax(numbers, 2)
	fmt.Println("Length", length)
	fmt.Println("Details", details())
}

func findMax(numbers []int, max int) (int, func() []int) {
	//var result []int //Kurang efektif untuk pemakaian memori
	var result = make([]int, 0, len(numbers))

	for _, num := range numbers {
		if num <= max {
			result = append(result, num)
		}
	}

	return len(result), func() []int {
		return result
	}
}
