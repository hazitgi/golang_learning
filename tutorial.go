package main

import "fmt"

func main() {
	/* Arrays */

	// numbers := [10]int{4, 5, 6, 8, 9, 8, 9, 45, 6, 4}

	// fmt.Println(numbers)

	/* Example 2 */
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(numbers[i])
	// }

	// var numbers2 [10]int

	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("Enter %d element : ", i)
	// 	fmt.Scanf("%d", &numbers2[i])
	// }

	// fmt.Println(sum(numbers2[:]))

	myArray := [...]int{256, 4, 6, 8, 9, 55, 66, 77, 86, 45}
	fmt.Println(myArray)

	/*  Example 3
	an array is of value type not of reference type.
	*/
	new_array := myArray
	fmt.Println("new array before", new_array)

	new_array[0] = 65324
	fmt.Println("Original Array after ", myArray)
	fmt.Println("new_array Array after ", new_array)

	/*
		Compare two arrays
	*/
	arr1 := [3]int{9, 7, 6}
	arr2 := [...]int{9, 7, 6}
	arr3 := [3]int{9, 5, 6}

	// Comparing arrays using == operator
	fmt.Println(arr1 == arr2)
	fmt.Println(arr2 == arr3)
	fmt.Println(arr1 == arr3)

	// This will give and error because the
	// type of arr1 and arr4 is a mismatch
	/*
		arr4 := [4]int{9, 7, 6}
		fmt.Println(arr1 == arr4)
	*/

}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
