package main

import (
	"fmt"
)

type hotdog int //tipo base

var b hotdog = 101  //tipado com a tipo base

func main() {

	//em go só temos conversion, não tem coercion ou casting
	x := 10
	fmt.Printf("%v\n",x) //10

	x = int(b)
	fmt.Printf("%v\n, %T",x, x) //101
//dessa forma convertemos o tipo
}