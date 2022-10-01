package main

import "fmt"

func main() {
	// static declaration
	// var x float64
	// x = 40356.14

	// Dynamic Declaration

	// y := 56

	// Printing

	// fmt.Println("Hello, World")
	// fmt.Println(x)
	// fmt.Printf("x is a type %T\n", x)
	// fmt.Printf("y is of type %T\n", y)

	// Mixed variable Declaration in GO

	// fmt.Println("Mixed Data types declaration")
	// var a, b, c = 3, 5, "Mohamed Haseeb K"
	// fmt.Println(a)
	// fmt.Printf("a data type %T\n", a)
	// fmt.Println(b)
	// fmt.Printf("a data type %T\n", b)
	// fmt.Println(c)
	// fmt.Printf("a data type %T\n", c)

	// fmt.Println("return multiple values from Function")
	// fmt.Println(swap("Haseeb", "Mohamed"))

	// var greeting = "greeting"

	// fmt.Printf("normal string: ")
	// fmt.Printf("%s", greeting)

	// println()
	// fmt.Printf("hex bytes: ")

	// for i:=0;i<len(greeting);i++{
	// 	fmt.Printf("%x",greeting[i])
	// }

	// Arrays
	// var variable_name [SIZE] variable_type
	// var balance = [10]float32{2,3,4}
	// fmt.Println(balance)

	// loops

	// String as a range in the for loop
	// for i, j := range "XabCd" {
	// 	fmt.Printf("The index number of %U is %d\n", j, i)
	// }

	// var value string = "five"
	// switch value {
	// case "one":
	// 	fmt.Println("C#")
	// case "two", "three":
	// 	fmt.Println("Go")
	// case "four", "five", "six":
	// 	fmt.Println("Java")
	// default:
	// 	fmt.Println("He")
	// }

	// simple for loop
	// for i := 0; i < 100; i++ {
	// 	fmt.Println(i)
	// }

	// infinit loop
	// for {
	// 	fmt.Printf("Mohamed Haseeb")
	// }

	// for loop as while Loop

	// for condition{
	// statement..
	// }

	// value := 0
	// var sum int
	// for value <= 50 {
	// 	fmt.Print(value, "\n")
	// 	sum += value
	// 	value++
	// }
	// fmt.Print(sum)

	// Simple range in for loop

	// for i, j:= range rvariable{
	// statement..
	//  }

	// friends_list := []string{"Mohamed Haseeb", "Akhil", "Faris", "Niyas"}
	// for i, j := range friends_list {
	// 	fmt.Println(i, j)
	// }

	//  Using for loop for strings

	// for i, j := range "Haseeb" {
	// 	fmt.Printf("The index number of %U is %d \n", j, i)
	// }

}

func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func swap(x, y string) (string, string) {
	return y, x
}
