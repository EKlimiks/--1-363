
package main

import "fmt"       

func main() {
  a := 5000
  c := 256
  o := a / c
  s := a - c*o
  fmt.Println(o)
  fmt.Println(s)
}