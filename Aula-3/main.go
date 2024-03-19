package main

import (
	"fmt"
)

func main (){
	var f float64 = 3.8
	c := int(f)
	fmt.Printf("%v:%T\n", c, c)
}