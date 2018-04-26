package main

import (
	"fmt"
)



func discriminante(a, b, c int) int  {
  return b*b - 4*a*c
}


func test_discriminante()  {
  x := 1
  y := 2
  z := 6
  w := 0
  w = discriminante(1,0,1)
  fmt.Printf("%d\n",w)
  w = discriminante(y,z,x)
  fmt.Printf("%d\n",w)
  z = z - 1
  w = discriminante(y,z,5)
  fmt.Printf("%d\n",w)
}

func obtener_raices(a, b, c int) string {
  D := discriminante(a, b, c)
  if D > 0 {
    // Raices reales y distintas
  } else if D == 0 {
    // Raices reales e iguales
  } else {
    // Raices complejas conjugadas
  }



}


func main() {
  test_discriminante()
}
