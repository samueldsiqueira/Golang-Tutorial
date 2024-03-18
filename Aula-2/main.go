package main

import "fmt"


func main(){
	var r rune = '😀'
	var s = "Hello World!!!"
	fmt.Println(s)
	fmt.Println(r)
}

// func main(){
// 	var u uint8 = 255
// 	fmt.Println(u+1)
// }

// func main() {
// 	var b bool = true
// 	fmt.Println(b)
// }

/*
Tipos Básicos:

bool 

string 

int int8 int16 int32 int64 
uint uint8 uint16 uint32 uint64 uintptr 

byte // alias para uint8 

rune // alias para int32 
     // representa um ponto de código Unicode 

float32 float64 

complex64 complex128


O exemplo mostra variáveis ​​de vários tipos, e também que declarações de variáveis ​​podem ser "fatoradas" em blocos, como acontece com instruções de importação.

Os tipos int, uint e uintptr geralmente têm 32 bits de largura em sistemas de 32 bits e 64 bits em sistemas de 64 bits. Quando você precisar de um valor inteiro, deverá usá-lo, int a menos que tenha um motivo específico para usar um tipo inteiro dimensionado ou sem sinal.



OBS RUNES, PARECE COMPLICADO...
*/