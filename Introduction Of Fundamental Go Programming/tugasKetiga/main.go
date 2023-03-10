package main

import (
	"fmt"
	"strings"
)

func main() {
	kalimat := "selamat malam"
	huruf := strings.Split(kalimat, "")
	count := make(map[string]int)

	for _, h := range huruf {
		if h != "" {
			fmt.Println(h)
			count[h]++
		}
	}

	fmt.Println("Jumlah kemunculan huruf:")
	for k, v := range count {
		fmt.Printf("%s: %d\n", k, v)
	}
}
