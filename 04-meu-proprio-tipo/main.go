package main

import (
	"fmt"
)

type hotdog int //tipo base

var b hotdog //tipado com a tipo base


func main(){
	x := 10 // não pode receber o tipo base, pois usou operador que já tipa a variável 
	//mesmo sendo igual não cria uma equivalência
	fmt.Printf("%T\n", x)
	fmt.Printf("%T", b)
}