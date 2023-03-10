package main

import "fmt"

type MenghitungHuruf struct {
	huruf  string
	jumlah int
}

func main() {
	input := "selamat malam"

	for i := 0; i < len(input); i++ {
		fmt.Printf("%c \n", input[i])
	}

	var hitung = [13]MenghitungHuruf{
		{"s", 0},
		{"e", 0},
		{"l", 0},
		{"a", 0},
		{"m", 0},
		{"t", 0},
		{" ", 0},
	}

	for i := 0; i < len(input); i++ {

		for u := 0; u < len(input); u++ {
			if hitung[u].huruf == string(input[i]) {
				hitung[u].jumlah = hitung[u].jumlah + 1
				continue
			}
		}

	}

	for _, a := range hitung {
		fmt.Printf("%+v\n", a)
	}
}
