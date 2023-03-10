package main

import (
	"fmt"
	"os"
)

func main() {
	data := map[string]string{
		"Nama":      os.Args[1],
		"Alamat":    os.Args[2],
		"Pekerjaan": os.Args[3],
		"Alasan":    os.Args[4],
	}
	displayData(data)
}

func displayData(data map[string]string) {
	fmt.Println("Nama		: ", data["Nama"])
	fmt.Println("Alamat		: ", data["Alamat"])
	fmt.Println("Pekerjaan	: ", data["Pekerjaan"])
	fmt.Println("Alasan		: ", data["Alasan"])
}
