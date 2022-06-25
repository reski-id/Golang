package main

import "fmt"

func main() {
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	slice := []string{"Reski", "Denny", "Sherly", "Budi", "Joko"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice { //for..range
		fmt.Println("Index", i, "=", value)
	}

	person := make(map[string]string) //bikin map
	person["name"] = "Reski"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}

	var number = []int{51, 4, 5, 6, 24, 11, 22, 33}

	//find-> Bilangan Ganjil Genap case_loop[1]
	for _, num := range number {
		if num%2 == 0 {
			fmt.Printf("%v Genap\n", num)
		} else {
			fmt.Printf("%v Ganjil\n", num)
		}
	}

}

/** case [loop]
1. Bilangan Ganjil Genap
2. Bilangan prima
3. Fibonaccy
**/
