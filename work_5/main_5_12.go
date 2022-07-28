package main

import (
	"fmt"
//	"math"
)

func main(){
	maxMph := 500
	passengerCapacity := 4
	airbags := true
	familyCar := passengerCapacity > 2 && airbags
	sportsCar := maxMph > 100 || passengerCapacity == 2
	canCategorize := !familyCar && !sportsCar
	
	fmt.Println(familyCar)
	fmt.Println(sportsCar)
	fmt.Println(canCategorize)
}