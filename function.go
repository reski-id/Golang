package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("Hello World")
}

//case_Function[1] -- adding two integer
func Penjumlahan(number1 int32, number2 int32) int32 {
	return number1 + number2
}

//case_Function[2] -- cek jika number ganjil return true
func Is_Ganjil(number int) bool {
	if number%2 == 1 {
		return true
	}
	return false
}

func main() {
	for i := 0; i < 3; i++ {
		sayHello()
	}

	fmt.Println(Penjumlahan(45, 23))
	fmt.Println(Is_Ganjil(9))
}

/** case [Function]
1. Function Penjumlahan
2. Function Mencari Bilangan Ganjil
3. Fibonaccy
4. Prima
**/
