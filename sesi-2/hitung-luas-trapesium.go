package main

import (
	"fmt"
)

func luasTrapesium(base1, base2, height float64) float64 {
	return 0.5 * (base1 + base2) * height
}

func main() {
	var base1, base2, height float64

	fmt.Print("Masukkan panjang alas atas: ")
	fmt.Scanln(&base1)

	fmt.Print("Masukkan panjang alas bawah: ")
	fmt.Scanln(&base2)

	fmt.Print("Masukkan tinggi trapesium: ")
	fmt.Scanln(&height)

	area := luasTrapesium(base1, base2, height)
	fmt.Printf("Luas trapesium: %.2f\n", area)
}
