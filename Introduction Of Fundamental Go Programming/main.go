package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Hello World")

	go prosesPertama(8)

	prosesKedua(8)

	fmt.Println("No. of goroutines", runtime.NumGoroutine())

	time.Sleep(time.Second * 2)

	fmt.Println("Selesai")

}

func prosesPertama(index int) {
	fmt.Println("Pertama")
	for i := 0; i <= index; i++ {
		fmt.Println("i = ", i)
	}
	fmt.Println("Proses pertama selesai")
}

func prosesKedua(index int) {
	fmt.Println("Kedua")
	for i := 0; i <= index; i++ {
		fmt.Println("j = ", i)
	}
	fmt.Println("Proses kedua selesai")
}
