package main

import "fmt"

func main() {
	// Functions

	// Declaration

	/*
		 	func function_name(parameter-list)(return_type){
				body
			}
	*/

	fmt.Printf("Area of rectagle is : %d\n", area(12, 50))
}

func area(length, width int) int {
	Ar := length * width
	return Ar
}
