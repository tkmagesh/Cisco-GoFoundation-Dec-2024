package main

import "fmt"

func main() {
	// var nos []int
	// var nos []int = []int{3, 1, 4, 2, 5}
	// var nos = []int{3, 1, 4, 2, 5}
	nos := []int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	fmt.Println("Iterating using indexer")
	for i := 0; i < len(nos); i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	fmt.Println("Iterating using range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	// adding new items
	nos = append(nos, 30)
	nos = append(nos, 10, 20)
	hundreds := []int{400, 700, 300}
	nos = append(nos, hundreds...)
	fmt.Println(nos)

	nos2 := nos
	nos2[0] = 9999
	fmt.Printf("nos[0] = %d, nos2[0] = %d\n", nos[0], nos2[0])

	fmt.Println("After sorting")
	sort(nos)
	fmt.Println(nos)

	// slicing
	fmt.Println("nos[2:5] =", nos[2:5]) // 2..4
	fmt.Println("nos[2:] =", nos[2:])   // 2..len(nos)
	fmt.Println("nos[:5] =", nos[:5])   // 0..4

	// using make()
	list := make([]int, 4)
	fmt.Println(list)
}

func sort(list []int) { // no return values
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
