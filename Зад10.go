package main

import (
	"fmt"
	"math"
)


func main() {
	var a, b int
	fmt.Print("Введите два целых числа")
	fmt.Scan(&a, &b)
	res := float64(a) / float64(b)
	c := math.Round(res)
	d :=  math.Floor(res)
	fmt.Println("Результат",res)
	fmt.Println("Округление до ближайщего целого",c)
	fmt.Println("Округление вниз",d)
	

}