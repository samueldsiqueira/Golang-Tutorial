package main

import (
	"fmt"
)

func main(){
	//ARITMÉTICOS
	sum := 999 + 1
	sub := 999 - 2
	division := 999 / 2
	restOfDivision := 999 % 5

	fmt.Println("Aqui está sua lista:",sum, sub, division,restOfDivision)
	fmt.Printf("%v, %T", restOfDivision, restOfDivision)

	//Atribuição
	var variable string = "hello"
	variable2 := "hellow there"
	fmt.Println(variable, variable2)

	//relacional

	fmt.Println(999 > 2) //true
	fmt.Println(999 >= 2) //true
	fmt.Println(999 == 2) //false
	fmt.Println(2 <= 999) //true
	fmt.Println(2 > 999) //false
	fmt.Println(999 != 2) //true

	//logico

	truthy, falsy := true, false
	fmt.Println(truthy && falsy) //false
	fmt.Println(truthy || falsy) //true
	fmt.Println(truthy || !falsy) //true

	// UNÁRIO
number := 10
number++ //11
number +=15 //26
fmt.Println(number)
}