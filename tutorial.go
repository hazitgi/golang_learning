package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Strings in Golang")

	str1 := "!! Mohamed Haseeb K !!"

	fmt.Println("String before trimming")
	fmt.Println("String 1 : ", str1)

	trim_str1 := strings.Trim(str1, "!")

	fmt.Println("Strings after trimming : ")
	fmt.Println("Result 1 : ", trim_str1)

}
