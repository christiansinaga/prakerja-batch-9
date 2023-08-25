package main

import "fmt"

func main() {
	max := 50

	fmt.Println("Bilangan kelipatan 7 yang kurang dari atau sama dengan 50:")

	for i := 0; i <= max; i += 7 {
		fmt.Println(i)
	}
}
