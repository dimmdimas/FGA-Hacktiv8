package main

import (
	"fmt"
)

func main() {
	kalimat := "selamat malam"
	hitung := make(map[string]int)

	for _, char := range kalimat {

		// fmt.Printf("%T ", char)

		h := string(char)

		// fmt.Printf("%T", h)

		fmt.Println(h)

		if h != "" {
			hitung[h]++
		}
	}

	fmt.Println(hitung)
}
