package main

import "fmt"

func main() {
	defer fmt.Println("DEFER")

	fmt.Println("Ga Defer")
}
