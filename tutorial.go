package main

import (
	"fmt"
)

func main() {
	// Functions

	// example 1
	func() {
		fmt.Println("Hi Mohamed Haseeb")
	}()

	// example 2

	func(ele string) {
		fmt.Println(ele)
	}("Hi How are you")

	// example 3

	value := func(p, q string) string {
		return p + " " + " " + q + " " + "you are welcome"
	}
	GFG(value)

	// Blank identifier
	multiply, _ := mul_div(56, 4)

	fmt.Println("result = ", multiply)

	_, division := mul_div(56, 4)

	fmt.Println("result = ", division)

}

func GFG(i func(p, q string) string) {
	fmt.Println(i("Mohamed", "Haseeb"))
}

func mul_div(num1, num2 int) (int, int) {
	return num1 * num2, num1 / num2
}
