package main

import "fmt"

func main() {
	var discount = 0.1 * 100
	fmt.Println(discount)

	var right, left = 13, 13
	fmt.Println("Balanced", (right == left))

	var total = 10000
	fmt.Println("Voucher", (total >= 10000) && (total <= 10000))

	// total += 1000
	// Sama dengan: total = total + 1000

	// total *= 2
	// Sama dengan: total = total * 2

	// Auti Incriment / Decriment
	// i++ / i = i + 1
	// i-- / i = i - 1

	var opr uint8 = 255
	fmt.Println("Total:", opr+1) // Expect: 256, actual: 0
}
