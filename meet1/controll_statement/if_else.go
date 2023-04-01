package main

import "fmt"

func main() {
	var point uint8 = 68

	if point >= 70 {
		fmt.Println("Kamu Lulus")
	} else {
		fmt.Println("Tidak lulus")
	}

	if point >= 90 {
		fmt.Println("Grade A")
	} else if point >= 80 {
		fmt.Println("Grade B")
	} else if point >= 70 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Grade D")
	}
}
