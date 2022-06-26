package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	firstName := "Reski"
	sayHelloTo(firstName, "Ahmad")
	sayHelloTo("Budi", "Nugraha")
}
