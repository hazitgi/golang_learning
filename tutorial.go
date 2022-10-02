package main

import "fmt"

func main() {
	fmt.Println("Hello, world")

	// Declaring a variable of struct type
	// All the struct fields are initialized
	// with their zero value
	var a Address
	fmt.Println(a)

	a1 := Address{"Mohamed Haseeb", "Tirurangadi", "kerala", 676306}
	fmt.Println("Address1 : ", a1)

	// Naming fields while
	// initializing a struct
	a2 := Address{name: "Arshad", state: "Kerala", Pincode: 676306}
	fmt.Println("Address 2", a2)

	// Uninitialized fields are set to
	// their corresponding zero-value
	a3 := Address{name: "Hamda"}
	fmt.Println("Address 3 : ", a3)

	/*
		Pointers to struct
	*/
	emp1 := &Employee{"Anas", "Karumbil", 25, 676306}

	fmt.Println("Frist Name : ", (*&emp1.firstName))
	fmt.Println("Full Name : ", (*&emp1.firstName)+" ", (*&emp1.lastName))
	fmt.Println("Age : ", (*emp1).age)
}

type Address struct {
	name    string
	city    string
	state   string
	Pincode int
}

type Employee struct {
	firstName, lastName string
	age, salary         int
}
