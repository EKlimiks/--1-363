package main

import (
	"fmt"
	"math"
)

func main() {
	var i float64
fmt.Print("начальная сумма")
fmt.Scan(&i)
var a float64
fmt.Print("годовая ставка")
fmt.Scan(&a)
var y float64
fmt.Print("количество лет")
fmt.Scan(&y)
sum :=i * math.Pow((1 + a/100),y)
fmt.Print(sum)
}