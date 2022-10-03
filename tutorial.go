package main

import (
	"fmt"
	"sort"
)

func main() {
	/* Slice */

	/* Example 1 */
	var my_slice_1 = []string{"Geeks", "for", "Geeks", "GFG"}

	fmt.Println("My Slice 1 : ", my_slice_1)

	/* using shorthand declaration */
	my_slice_2 := []int{5, 6, 8, 9, 4, 65, 9, 4, 6, 5}
	fmt.Println("my_slice_2 : ", my_slice_2)

	/* Using an Array */

	var array_slice_1 = my_slice_1[1:2]
	array_slice_2 := my_slice_1[0:]
	array_slice_3 := my_slice_1[:2]
	array_slice_4 := my_slice_1[:]

	/* Display the result */
	fmt.Println("My Array: ", my_slice_1)
	fmt.Println("Array Slice 1: ", array_slice_1)
	fmt.Println("Array Slice 2: ", array_slice_2)
	fmt.Println("Array Slice 3: ", array_slice_3)
	fmt.Println("Array Slice 4: ", array_slice_4)

	/* using make functions */

	var new_slice = make([]int, 4, 7)
	fmt.Printf("Slice 1 = %v \n length = %d, \n capacity = %d\n", new_slice, len(new_slice), cap(new_slice))

	var new_slice_2 = make([]int, 7)
	fmt.Printf("Slice 2 = %v \n length = %d, \n capacity = %d\n", new_slice_2, len(new_slice_2), cap(new_slice_2))

	/* Multi-Dimensional Slice Start */
	multi_slice := [][]int{
		{5, 6}, {1, 6}, {5, 69},
	}

	fmt.Println(multi_slice)

	/* Multi-Dimensional Slice End */
	scl1 := []int{50, 6, 8, 9, 4, 88}
	scl2 := []int{-5, 0, -1, 6, 8, 9, 7}

	fmt.Println("Slices(Before):")
	fmt.Println("Slice 1: ", scl1)
	fmt.Println("Slice 2: ", scl2)

	// / Sorting the slice of ints
	// Using Ints function
	sort.Ints(scl1)
	// sort.Ints(scl2)

	// Displaying the result
	fmt.Println("\nSlices(After):")
	fmt.Println("Slice 1 : ", scl1)
	fmt.Println("Slice 2 : ", scl2)

	/* Sorting of sline */

	/* IntsAreSorted verify is slice is sorted or not Start */
	// func IntsAreSorted(scl []int) bool

	fmt.Println("is scl1 is sorted ? : ", sort.IntsAreSorted(scl1))
	fmt.Println("is scl2 is sorted ? : ", sort.IntsAreSorted(scl2))

	/* IntsAreSorted verify is slice is sorted or not End */

}
