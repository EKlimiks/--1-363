package main

import ("fmt" 
"math")

func main() {

var r float64
fmt.Print("Введите радиус")
fmt.Scan(&r)
L := 2 * math.Pi * r
S := math.Pi * r * r
fmt.Println(L)
fmt.Println(S)

}