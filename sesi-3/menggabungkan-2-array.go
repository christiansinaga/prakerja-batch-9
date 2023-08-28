package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	merged := make([]string, 0)

	// Create a map to track the names that have been added
	addedNames := make(map[string]bool)

	// Add names from arrayA to the merged slice and the map
	for _, name := range arrayA {
		if !addedNames[name] {
			merged = append(merged, name)
			addedNames[name] = true
		}
	}

	// Add names from arrayB to the merged slice if they are not duplicates
	for _, name := range arrayB {
		if !addedNames[name] {
			merged = append(merged, name)
			addedNames[name] = true
		}
	}

	return merged
}

func main() {
	// Test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))          // ["king", "devil jin", "akuma", "eddie", "steve", "geese"]
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))                         // ["sergei", "jin", "steve", "bryan"]
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"})) // ["alisa", "yoshimitsu", "devil jin", "law"]
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))                                          // ["devil jin", "sergei"]
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))                                                     // ["hwoarang"]
	fmt.Println(ArrayMerge([]string{}, []string{}))                                                               // []
}
