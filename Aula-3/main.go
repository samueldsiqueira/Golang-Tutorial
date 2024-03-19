package main

import (
	"fmt"
)

func main (){
	s:= "hello"
	c:= string([]byte(s))
	fmt.Printf("%v: %T", c, c)
	// c, err := strconv.ParseFloat("hello", 64)
	// fmt.Printf("%v:%T\n", c, c)
	// fmt.Println(err)
}