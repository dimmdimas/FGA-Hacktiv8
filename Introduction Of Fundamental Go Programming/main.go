package main

import "fmt"

func main() {
	dapat("dimas", 23, 192123)
}

func dapat(nama string, umur, kode int) {
	fmt.Printf("Nama saya %s, umur %d. [%d]\n", nama, umur, kode)
}
