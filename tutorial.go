package main

import (
	"fmt"
)

func main() {
	// Functions

	// Declaration

	/*
		 	func function_name(parameter-list)(return_type){
				body
			}
	*/

	// fmt.Printf("Area of rectagle is : %d\n", area(12, 50))

	// var p int = 10
	// var q int = 20

	// fmt.Printf("p = %d and q = %d\n", p, q)
	// swap(&p, &q)
	// fmt.Printf("p = %d and q = %d\n", p, q)

	// zero argument
	// fmt.Println(joinstr())

	// multiple arguments
	// fmt.Println(joinstr("GEEK", "GFG"))
	// fmt.Println(joinstr("Geeks", "for", "Geeks"))
	// fmt.Println(joinstr("G", "E", "E", "k", "S"))

	message := ABC()
	fmt.Println(message("Habeebee", "welcome"))

}

// func area(length, width int) int {
// 	Ar := length * width
// 	return Ar
// }

// func swap(a, b *int) int {
// 	var o int
// 	o = *a
// 	*a = *b
// 	*b = o
// 	return o
// }

// Variadic Functions in Go

// func joinstr(elements ...string) string {
// 	return strings.Join(elements, "-")
// }

func ABC() func(i, j string) string {
	myf := func(i, j string) string {
		return i + " " + j + " " + "Tirurangadi"
	}
	return myf
}
