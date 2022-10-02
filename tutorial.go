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

}

type author struct {
	name      string
	branch    string
	particles int
	salary    int
}

func (a author) show() {
	fmt.Println("Author's Name : ", a.name)
	fmt.Println("Branch Name : ", a.branch)
	fmt.Println("Published articles : ", a.particles)
	fmt.Println("Salary : ", a.salary)
}
