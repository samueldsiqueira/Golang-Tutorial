package main

import (
	"fmt"
)

// são sucessivas constantes não tipadas como inteiras;
//podendo ser usadas como float
// - ref/spec#Iota
// - Numa declaração de constantes, o identificador iota representa números sequenciais.
// - na pratica:
// iota, iota +1, a = iota b c, reinicia em cada const, _
const (
	a = iota //Untyped int.
	_ = iota //Untyped int ignora valor.
	c = iota //Untyped int.
	x =  iota //Untyped int.
	_ = iota //Untyped int ignora valor.
	z = iota //Untyped int.
)
func main (){
	fmt.Println(a, c, x, z)//0 2 3 5
}