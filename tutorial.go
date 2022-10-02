package main

import "fmt"

func main() {
	/*
		Nested struct
	*/

	/* Example 1  Start*/

	result := HR{
		details: Author{"Mohamed Haseeb K", "Tirurangadi", 25},
	}

	fmt.Println("Details of Author")
	fmt.Println(result)
	/* Example 1  End*/

	/* Example 2 Start */
	result2 := Teacher{
		name:    "Mohamed Haseeb K",
		subject: "Go / Golang",
		exp:     2,
		details: Student{name: "Aflah", branch: "Commerce", year: 2},
	}

	fmt.Println("Details of the Teacher")
	fmt.Println("Teacher's name : ", result2.name)
	fmt.Println("Subject : ", result2.subject)
	fmt.Println("Experience : ", result2.exp)

	fmt.Println("Details of student")
	fmt.Println("Student's name : ", result.details.name)
	fmt.Println("Student's branch name: ", result2.details.branch)
	fmt.Println("Year : ", result2.details.year)

	fmt.Println(result2)

	/* Example 2 End */

}

/* 	Nested struct Syntax */
/*

	type struct_name_1 struct{
  		// Fields
	}

	type struct_name_2 struct{
  		variable_name  struct_name_1
	}

*/

/* Example 1 */
type Author struct {
	name, branch string
	year         int
}

type HR struct {
	details Author
}

/* Example 2  Start */
type Student struct {
	name, branch string
	year         int
}

type Teacher struct {
	name, subject string
	exp           int
	details       Student
}

/* Example 2  End*/
