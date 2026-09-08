package main

import "fmt"

func main() {
	var a float64
	fmt.Print("Введите сумму покупки")
	fmt.Scan(&a)
	b := a * 0.8
	fmt.Println(b)
}