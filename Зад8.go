package main

import (
"fmt"
"math"
)
type Point struct
{
  x float64
  y float64
}
func dist(p1, p2 Point) float64 {
  return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}
func main() {
  T1 := Point{1.0, 2.0}
  T2 := Point{4.0, 6.0}
  fmt.Println("Расстояние между точками: ", dist(T1, T2))
}