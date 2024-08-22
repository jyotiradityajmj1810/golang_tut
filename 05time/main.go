package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("time in golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2002,time.October,18,23,23,0,0,time.UTC) 

	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}

//GOOS="linux" go build 
// it is used to build executables for the specific os 