package main

import (
	"fmt"
	"runtime"
)

var h = 10 //mesma coisa se usar fora o valor default do tipo é declarado
var i = 10.0 //mesma coisa se usar fora o valor default do tipo é declarado
//são tipos numéricos mas não são compatíveis tem de converter para
//transformar valor inteiro em float

func main (){
	a := "e"
	b:= "é"
	c:= "的"
	fmt.Printf("%v, %v, %v\n",a,b,c)
	d := []byte(a) // valor 101 tabela ASCII
	e := []byte(b) // UTF8 vem depois da tabela ASCII precisa de tem 2 bytes [195 169]
	f := []byte(c) // se precisar de 3 ou 4 bytes  utilizamos o tipo rune em GO que representa UTF8.
	fmt.Printf("%v, %v, %v\n", d, e, f)
	
	//EXPLICANDO TIPAGEM NÚMERICA
	x:= 10 //DEIXAR O GO TIPAR USANDO ATRIBUIÇÃO RÁPIDA A "MARMOTA"
	y:= 10.0//DEIXAR O GO TIPAR USANDO ATRIBUIÇÃO RÁPIDA A "MARMOTA"
	fmt.Printf("%v, %T\n",x,x)
	fmt.Printf("%v, %T\n",y,y)
	fmt.Printf("%v, %T\n",h,h)
	fmt.Printf("%v, %T\n",i,i)


	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}

/*todos os tipoos numericos são destirons, exeto:
		- bytes = uint8
		- rune = int32 (UTF8) - strings são feitos de runes, cada rune
		 é um carácter esse carácter pode ter 1, 2 , 3 ou 4 bytes
	
*/