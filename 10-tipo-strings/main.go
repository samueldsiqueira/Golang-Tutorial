package main

import (
	"fmt"
)

//strings são imutáveis
//para mudar tenho que criar uma nova variavel e alterar anterior apartir dela.
//string é um slice de bytes
//%#U = Unical unicode point
//%#x = Decimal
//ASCII é um subset de utf8 que é o que realmente GO usa

func main(){
	s := "ascii êøâ 香"
	// sb := []byte(s)

	for _, v := range s {
		fmt.Printf("%b - %T - %#U - %#x\n", v, v, v, v,)
	}

	fmt.Println(`░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░►
	
	`)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%b - %T - %#U - %#x\n", s[i], s[i], s[i], s[i])
			}
}