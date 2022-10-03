package main

import "fmt"

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

}
