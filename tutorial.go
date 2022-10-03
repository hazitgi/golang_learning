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

	/*
		How to copy an Array into another array in golang
	*/

	// creating a copy of an array by value
	arr5 := arr1

	arr5[2] = 6598

	fmt.Println("Create a copy of an array by value (original Array)", arr1)
	fmt.Println("Create a copy of an array by value (new Array)", arr5)

	/*  if copy an array by reference that will effect original array.
	if copy an array by value that won't effect original array
	*/

	// Creating a copy of an array by reference
	arr6 := &arr1

	arr6[2] = 3269

	fmt.Println("Create a copy of an array by value (original Array)", arr1)
	fmt.Println("Create a copy of an array by value (new Array)", arr6)

	/* Pass array in functions Start */
	fmt.Println(myFun(arr2, 3))
	/* Pass array in functions End */

}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

/* Pass array into functions Start */
func myFun(a [3]int, size int) int {
	var k, val, r int

	for k = 0; k < size; k++ {
		val += a[k]
	}
	r = val / size
	return r
}

/* Pass array into functions End */
