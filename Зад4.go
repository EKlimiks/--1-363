package main

import "fmt"      

func main() {
var F float64
fmt.Print("Введите темп фаренгейту")
fmt.Scan(&F)
T:= 5.0/9.0*(F-32.0)
fmt.Println(T)
}