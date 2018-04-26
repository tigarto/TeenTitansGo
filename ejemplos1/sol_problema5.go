package main

import (
	"fmt"
  "math"
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

func obtener_raices(a, b, c int) {
  var D float64 = float64(discriminante(a, b, c))
  if D > 0 {
    // Raices reales y distintas
    // x1 = (-b - D^(1/2))/(2*a)
    var x1 float64 = (float64(-b) - math.Sqrt(D))/float64(2*a)
    //x2 = (-b + D^(1/2))/(2*a)
    //var x2 float64 = (float64(-b) + math.Sqrt(D))/float64(2*a)
    fmt.Printf("%.1f con multiplicidad 2.\n",x1)
  } else if D == 0 {
    // Raices reales e iguales
    // x1 = x2 = -b/(2*a)
    var x1 float64 = float64(-b/(2*a))
    var x2 float64 = float64(-b/(2*a))
    fmt.Printf("%.1f, %.1f\n",x1, x2)
  } else {
    // Raices complejas conjugadas
    // x1 = (-b/(2*a))i - ((-D)^(1/2))/(2*a))j
    var pR float64 = float64(-b/(2*a))
    var pI float64 = math.Sqrt(-D)/float64(2*a)
    if pR != 0 {
      if pI > 0 {
        fmt.Printf("%.1f + %.1fj, %.1f - %.1fj\n",pR, pI,pR,pI)
      } else if pI < 0 {
        fmt.Printf("%.1f + %.1fj, %.1f - %.1fj\n",pR, -pI, pR,-pI)
      }
    } else {
      if pI > 0 {
        fmt.Printf("%.1fj, -%.1fj\n",pI, pI)
      } else {
        fmt.Printf("%.1fj, -%.1fj\n",-pI, -pI)
      }
    }
  }
}


func main() {
  test_discriminante()
}
