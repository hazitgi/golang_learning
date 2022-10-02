package main

import "fmt"

func main() {
	/*
		Nested struct
	*/


	/* Example 1 */

	result := HR{
		details: Author{"Mohamed Haseeb K", "Tirurangadi", 25},
	}

	fmt.Println("Details of Author")
	fmt.Println(result)

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
