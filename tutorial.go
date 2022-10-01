package main

import "fmt"

func main() {

	// Switch Statement in Go

	// Expression Switch
	// Type Switch

	// Expression Switch

	// switch optstatement; optexpression{
	// 	case expression1: Statement..
	// 	case expression2: Statement..
	// 	...
	// 	default: Statement..
	// 	}

	// // Expression Switch

	// switch day := 4; day {
	// case 1:
	// 	fmt.Println("Monday")
	// case 2:
	// 	fmt.Println("Tuesday")
	// case 3:
	// 	fmt.Println("Wednesday")
	// case 4:
	// 	fmt.Println("Thursday")
	// case 5:
	// 	fmt.Println("Friday")
	// case 6:
	// 	fmt.Println("Saturday")
	// case 7:
	// 	fmt.Println("Sunday")
	// default:
	// 	fmt.Println("Invalid")
	// }

	// Exampole 2

	// number1 := 1
	// number2 := 4

	// switch {
	// case number1 == 2:
	// 	fmt.Print("number1 is 1")
	// case number2 == 4:
	// 	fmt.Print("number2 is 4")
	// }

	// Switch statement without default statement
	// Multiple values in case statement

	// value := "five"

	// switch value {
	// case "one":
	// 	fmt.Println("Hai case one")
	// case "two", "three":
	// 	fmt.Println("case is two , three")
	// case "four", "five", "six":
	// 	fmt.Println("case four, five, six:")
	// }

	// 2. Type Switch

	/*
		switch optstatement; typeswitchexpression{
			case typelist 1: Statement..
			case typelist 2: Statement..
			...
			default: Statement..
		}
	*/

	var value interface{}
	switch q := value.(type) {
	case bool:
		fmt.Println("value is of boolean type")
	case float32:
		fmt.Println("value is of float32 type")
	case int:
		fmt.Println("value is of int type")
	default:
		fmt.Printf("value is of type : %T", q)

	}

}
