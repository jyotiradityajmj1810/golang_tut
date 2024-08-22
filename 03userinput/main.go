package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("user input in golang")

	reader := bufio.NewReader(os.Stdin)

	// coma ok or err err syntax

	fmt.Println("enter user name ")
	uName ,_ := reader.ReadString('\n')
	fmt.Println("you entered",uName)
	fmt.Printf("the type of the user input is = %T\n", uName)

	fmt.Println("enter user age ")
	uAge ,_ := reader.ReadString('\n')
	fmt.Println("you entered ",uAge)
	fmt.Printf("the type of the user input is = %T\n", uAge)
}
