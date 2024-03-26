package main

import (
	"fmt"
)

const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB // 1 << (2 * 10)
	GB // 1 << (3 * 10)
	TB // 1 << (4 * 10)
)
// - Deslocamento de bits é quando deslocamos digitos binários para a esquerda ou direita:
// - Na prática:
// - %d %b
// - x << y
// - iota *10 << 10 = kb, mb, gb
func main () {
	x := 1
	y := x << 1
	z := x >> 1
	fmt.Printf("%b\n", x)
	fmt.Printf("%b\n", y)
	fmt.Printf("%b\n", z)

	fmt.Println("binary\t\t\t\tdecimal")
	fmt.Printf("%b\t\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)

	//o operador de deslocamento de bits é ess "<<" para um lado e ">>"
}