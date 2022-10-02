package main

import "fmt"

func main() {
	// Syntax

	// func(reciver_name Type) method_name(parameter_list)(return_type){
	// 	// Code
	// }

	// Method with struct type receiver

	res := author{
		name:      "Mohamed Haseeb",
		branch:    "Tirurangadi",
		particles: 50,
		salary:    365894,
	}

	res.show()


	// Method with Non-Struct Type Receiver

	value1 := data(23)
	value2 := data(56)
	res2 := value2.multiply(value1)
	fmt.Println("Final result ", res2)

}

type author struct {
	name      string
	branch    string
	particles int
	salary    int
}

type data int

func (a author) show() {
	fmt.Println("Author's Name : ", a.name)
	fmt.Println("Branch Name : ", a.branch)
	fmt.Println("Published articles : ", a.particles)
	fmt.Println("Salary : ", a.salary)
}

func (d1 data) multiply(d2 data) data {
	return d1 * d2
}

