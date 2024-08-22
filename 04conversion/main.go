package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("conversion in golang")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("enter the marks in maths")
	marksMath ,_ := reader.ReadString('\n')
	fmt.Println("math marks is = ",marksMath)

	fmt.Println("enter the marks in science")
	marksScience , _ := reader.ReadString('\n')
	fmt.Println("marks in science is = ",marksScience)

	mMarks,_:= strconv.ParseInt(strings.TrimSpace(marksMath) , 10,64)
	sMarks,_:= strconv.ParseInt(strings.TrimSpace(marksScience) ,10,64)



	fmt.Println("total marks is = ",mMarks+sMarks)

}
