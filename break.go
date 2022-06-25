package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan ke ", i)
	}

	//1. Case loop-break[1] : Loop array bilangan break pada angka 3
	var number = []int32{9, 8, 7, 6, 5, 4, 3, 2, 1}

	for _, value := range number {
		if value == 3 {
			fmt.Println("break pada angka 3")
			break
		}
		fmt.Println(value)
	}
}

/** case loop-break[1]
	1. Loop array bilangan break pada angka 3
**/
