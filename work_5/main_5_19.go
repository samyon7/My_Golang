package main

import (
	"fmt"
	"strconv"
)

func main(){
	val1 := "0"
	
	if bool1, b1err := strconv.ParseBool(val1); b1err == nil{
		fmt.Println("Parsed Value :",bool1)
	} else {
		fmt.Println("Cannot parse")
	}
}