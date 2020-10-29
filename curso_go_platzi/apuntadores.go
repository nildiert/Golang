package main

import "fmt"

func main() {
	x := 25
	fmt.Println(x)
	cambiarValor(x)
	y := &x
	fmt.Println(x)
	fmt.Println(*y)
}

func cambiarValor(a int) {
	a = 36
}
