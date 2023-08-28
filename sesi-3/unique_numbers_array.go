package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	countMap := make(map[int]int)
	var result []int

	for _, char := range angka {
		num, err := strconv.Atoi(string(char))
		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			return nil
		}
		countMap[num]++
	}

	for num, count := range countMap {
		if count == 1 {
			result = append(result, num)
		}
	}

	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6, 3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8 7 2 5 4]
}
