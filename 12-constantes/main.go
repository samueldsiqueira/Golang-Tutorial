package main

import (
	"fmt"
)

//### CONTANTES
// - São valores imutáveis;
// - Podem ser tipadas ou não:
// - const oi = "bom dia"
// - const oi string = "bom dia"
// - As não tipadas só terão um tipo atribuido a elas quando forem usadas.
// - ex. qual o tipo de 42? int? uint? float64?
// - ou seja, é uma flexibilidade conveniente.
// na prática: int, float, string.
// const x = y
//const (x=y)

const (
	x = 10
	y = 20
	z = 30
)
//podemos declarar mais de uma de uma vez


func main (){
	fmt.Println(x, y, z)
}