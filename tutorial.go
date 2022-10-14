package main

import "fmt"

func main() {
	fmt.Println("Pointers in Golang")

	x:=0xFF
	y:=0x9C

	fmt.Printf("Type of variable x is %T\n",x)
	fmt.Printf("Value of x in hexadecimal is  %X\n",x)
	fmt.Printf("Value of x in decimal is  %v\n",x)

	fmt.Printf("Type of variable x is %T\n",y)
	fmt.Printf("Value of x in hexadecimal is  %X\n",y)
	fmt.Printf("Value of x in decimal is  %v\n",y)

	/* Normal variable declaration */
	var a = 23

	/* Initialization of pointer s with */
	/* memory address of variable a */
	var s *int = &a

	fmt.Println("value stored in x = ", a)
	fmt.Println("Address of  x = ", &a)
	fmt.Println("Value stored in variable p = ", s)

	/* 1. The default value or zero-value of a pointer is always nil.  */

	var b *int
	fmt.Println(" b = ",b)


	ab:=458
	ac:=&ab

	fmt.Println("Values stored in ab = ",ab)
	fmt.Println("Address of y = ", &ab)
	fmt.Println("Value stored in pointer varibale p = ", ac)

	/* 
		This is a dereferencing a pointer
		using * operator before a pointer
		variable to access the value stored
		at the variable at which it is pointing
	*/
	fmt.Println("Value stored in ab(*ac) = ", *ac)


	fmt.Println("=======================================")
	var aq = 458
	aw:=&aq

	fmt.Println("values stored in aq before changeing = ", aq)
	fmt.Println("Address of aq = ", &aq)
	fmt.Println("Value stored in pointer variable p = ", aw)




	fmt.Println("value stored in aq(*aw) before changing = ",*aw)

	*aw = 500
	fmt.Println("values stored in aq after changeing = ", aq)
	fmt.Println("values stored in aq(*aw) after changeing = ", *aw)



}
