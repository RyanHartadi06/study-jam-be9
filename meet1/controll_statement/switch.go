package main

import "fmt"

func main() {
	var grade = "Grade B"

	switch grade {
	default:
		fmt.Println("Point range 0 - 69")
	case "Grade A":
		fmt.Println("Point range 90 - 100")
	case "Grade B":
		fmt.Println("Point range 80 - 89")
	case "Grade C":
		fmt.Println("Point range 70 - 79")
	}

	switch {
	case grade == "Grade A":
		fmt.Println("#2 Point range 90 - 100")
	case grade == "Grade B":
		fmt.Println("#2 Point range 80 - 89")
	}

	var point uint8 = 80
	switch {
	default:
		fmt.Println("Anda Tidak Lulus")

	case point <= 100 && point >= 70:
		fmt.Println("Selamat Anda Lulus")
	}
}
