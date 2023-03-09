package main

import "fmt"

func main() {

	i := 0
	j := 0

	for i < 5 {
		fmt.Println("Nilai i = ", i)
		i++
	}

	for j < 5 {
		fmt.Println("Nilai j = ", j)
		j++
	}

	for a, b := range "CАШАВО" {
		fmt.Printf("character %#U starts at byte position %d\n", b, a)
	}

	for {
		if j++; j <= (10) {
			fmt.Println("Nilai j = ", j)
		} else {
			break
		}
	}
}
