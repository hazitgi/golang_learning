package main

import "fmt"

func main() {
	// Syntax

	// func(reciver_name Type) method_name(parameter_list)(return_type){
	// 	// Code
	// }

	// Method with struct type receiver

	// res := author{
	// 	name:      "Mohamed Haseeb",
	// 	branch:    "Tirurangadi",
	// 	particles: 50,
	// 	salary:    365894,
	// }

	// res.show()

	// Method with Non-Struct Type Receiver

	// value1 := data(23)
	// value2 := data(56)
	// res2 := value2.multiply(value1)
	// fmt.Println("Final result ", res2)

	// Methods with Pointer Receiver

	// Syntax
	/*
		func (p *Type) method_name(...Type) Type {
		// Code
		}
	*/

	// emplData := EmployeeInfo{
	// 	name:   "Mohamed Haseeb",
	// 	branch: "Tirurangadi",
	// }

	// fmt.Println("Employee name: ", emplData.name)
	// fmt.Println("Branch Name (Before) : ", emplData.branch)

	// p := &emplData

	// p.changeBranch("Chemmad")
	// println("____________________________")
	// fmt.Println("Employee name: ", emplData.name)
	// fmt.Println("Branch Name (Before) : ", emplData.branch)

	// Method Can Accept both Pointer and Value

	newRes := author{
		name:   "Faris Poongadan",
		branch: "Tirurangadi",
	}

	fmt.Println("Branch name (before) : ", newRes.branch)

	// Calling the show_1 method
	// (Pointer methods) with value

	newRes.show_1("Thalappra")
	fmt.Println("Branch name (After) : ", newRes.branch)

	// calling the show_2 method
	// (value method) with a pointer
	(&newRes).show_2()
	fmt.Println("Author's name (after) : ", newRes.name)

}

type author struct {
	name      string
	branch    string
	particles int
	salary    int
}

// type data int

// func (a author) show() {
// 	fmt.Println("Author's Name : ", a.name)
// 	fmt.Println("Branch Name : ", a.branch)
// 	fmt.Println("Published articles : ", a.particles)
// 	fmt.Println("Salary : ", a.salary)
// }

// func (d1 data) multiply(d2 data) data {
// 	return d1 * d2
// }

// type EmployeeInfo struct {
// 	name      string
// 	branch    string
// 	particles int
// }

// func (a *EmployeeInfo) changeBranch(abranch string) {
// 	(*a).branch = abranch
// }

func (a *author) show_1(abranch string) {
	(*a).branch = abranch
}
func (a author) show_2() {
	a.name = "Mohamed Haseeb K"
	fmt.Println("Author's name (before) : ", a.name)
}
