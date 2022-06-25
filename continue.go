package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Perulangan ke", i)
	}
}

/** case loop-continue[1]
	1. Temukan Angka Ganjil
	2. Buatlah array 1-9, Jika ada 4689 skip, print angka selain 4689
**/
