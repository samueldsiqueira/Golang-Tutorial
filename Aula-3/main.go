package main

import (
	"fmt"
)

func main (){
	var i int = 129
	c := int8(i)
	fmt.Printf("%v:%T\n", c, c)
}