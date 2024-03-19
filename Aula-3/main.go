package main

import (
	"fmt"
)

func main (){
	var i int = 1
	c := int64(i)
	fmt.Printf("%v:%T\n", c, c)
}