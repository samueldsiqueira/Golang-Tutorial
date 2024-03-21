package main

import (
	"fmt"
)
type brinco int
var x brinco
var y int
func main(){
	fmt.Printf("%v\n %T\n",x,x)
	x = 42
	fmt.Println(x,)
	y = int(x)
	fmt.Printf("%v\n %T\n",y,y)
}