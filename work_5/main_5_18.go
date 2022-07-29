package main

import (
	"fmt"
	"strconv"
)

func main (){
	val1 := "0"
	bool1, b1err := strconv.ParseBool(val1)
	
	if b1err == nil{
		fmt.Println("Parsed Value is ",bool1)
	} else {
		fmt.Println("Canot Parse "val1)
	}
}