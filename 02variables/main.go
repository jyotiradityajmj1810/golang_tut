package main

import "fmt"

const CollageName string = "xyz"

func main() {
	fmt.Println("variables in golang ")

	var userName string = "jmj"
	fmt.Println(userName)
	fmt.Printf("the type of variable is = %T \n", userName)

	var userAge int = 22
	fmt.Println(userAge)
	fmt.Printf("the type of variable is = %T \n", userAge)

	var userCgpa float32 = 8.9
	fmt.Println(userCgpa)
	fmt.Printf("the type of variable is = %T \n", userCgpa)

	var islogedin bool = true
	fmt.Println(islogedin)
	fmt.Printf("the type of variable is = %T \n", islogedin)

	//default values and some aliases (golang dose not put any garbage value in noting is assigned to it) .
	var marks int
	fmt.Println(marks)
	fmt.Printf("the type of variable is = %T \n", marks)

	var regNo string
	fmt.Println(regNo)
	fmt.Printf("the type of variable is = %T \n", regNo)

	// implicit type
	var class = "3b"
	fmt.Println(class)
	fmt.Printf("the type of variable is = %T \n", class)

	// no var type
	gender := "male"
	fmt.Println(gender)
	fmt.Printf("the type of variable is = %T \n", gender)

	// call of the constant
	fmt.Println(CollageName)
	fmt.Printf("the type of variable is = %T \n", CollageName)

	var a,b int = 23,24
	fmt.Println(a,b)

	const c  = 6600000000000
	fmt.Println(int64(c))
}
