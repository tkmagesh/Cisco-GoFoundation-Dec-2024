package main

import "fmt"

func main() {
	/*
		var nos [5]int // memory allocated & initialised
		nos[0] = 3
		nos[1] = 4
		nos[2] = 1
		nos[3] = 2
		nos[4] = 5
	*/

	// var nos [5]int = [5]int{3, 1, 4, 2, 5}

	// type inference
	// var nos = [5]int{3, 1, 4, 2, 5}

	// :=
	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	fmt.Println("Iterating using indexer")
	for i := 0; i < 5; i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	fmt.Println("Iterating using range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	/*
		nos2 := nos //creating a copy of nos
		nos2[0] = 9999
		fmt.Printf("nos[0] = %d, nos2[0] = %d\n", nos[0], nos2[0])
	*/

	/*
		nos2 := &nos
		(*nos2)[0] = 9999
		fmt.Printf("nos[0] = %d, nos2[0] = %d\n", nos[0], (*nos2)[0])
	*/

	// accessing the members of a pointer doesnot require deferencing

	nos2 := &nos
	nos2[0] = 9999
	fmt.Printf("nos[0] = %d, nos2[0] = %d\n", nos[0], nos2[0])

	sort( /* ? */ )  // use sort() to sort the nos array
	fmt.Println(nos) // print the sorted array
}

func sort( /*  */ ) { // no return values

}
