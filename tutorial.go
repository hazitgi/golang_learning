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

}

func GFG(i func(p, q string) string) {
	fmt.Println(i("Mohamed", "Haseeb"))
}
