package main

import (
	"fmt"
)

/*
###Overflow

- Em uint16, por exemplo, vai de 0 a 65535
- Que acontece se tentar usar 65536?
- Ou se a gente estiver em 65535 e tentar somar mais 1?
*/

func main (){
	var i uint16
	i = 65535
	fmt.Println(i)
	i++
	fmt.Println(i) //0 como um odômetro de carro

}