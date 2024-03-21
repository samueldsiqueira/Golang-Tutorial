package main

import (
	"fmt"
)

func main(){

	//
	//é uma string interpretada
	x := "oi bom dia\ncomo vai?\tespero que tudo bem.\n"
	//agora vamos ver as raw string literal
	y := `
	 oi boa noite
	 como vai?
	 espero que bem
	 `
	 //a diferença entre elas é que tudo que é escrito entre o acento agudo é interpretado
	 //como está escrito, já no outro precisamos dar os comandos de tab ou quebra de linha (/t e /n )
	fmt.Println(x,y)

	a:="oi"
	b:="bom dia"
	c:= fmt.Sprint(a, b)
	//retorno vai ser oibom dia, ela junta todos argumentos em uma string.
	//SprintF não joga na tela o valor ele só junta os valores.
	fmt.Println(c)

	/*Fprint é um file print
	é usado para muitas coisas
	*/
}