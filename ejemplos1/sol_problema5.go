package main

import (
	"fmt"
  "math"
)

func discriminante(a, b, c float64) float64  {
  return b*b - 4*a*c
}

func obtener_raices(a, b, c float64) {
  var D float64 = discriminante(a, b, c)
  if D > 0 {
    // Raices reales y distintas
    var x1 float64 = ((-b) - math.Sqrt(D))/(2*a)
    var x2 float64 = ((-b) + math.Sqrt(D))/(2*a)
    fmt.Printf("%.1f, %.1f\n",x1, x2)
  } else if D == 0 {
    // Raices reales e iguales
    var x float64 = -b/(2*a)
    fmt.Printf("%.1f con multiplicidad 2.\n",x)
  } else {
    // Raices complejas conjugadas
    var pR float64 = -b/(2*a)
    var pI float64 = math.Sqrt(-D)/(2*a)
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

func test_discriminante()  {
  var x float64 = 1
  var y float64 = 2
  var z float64 = 6
  var w float64 = 0
  w = discriminante(1,0,1)
  fmt.Printf("%d\n",w)
  w = discriminante(y,z,x)
  fmt.Printf("%d\n",w)
  z = z - 1
  w = discriminante(y,z,5)
  fmt.Printf("%d\n",w)
}

func test_bachiller()  {
  var a float64 = 1
  var b float64 = 2
  var c float64 = 2
  fmt.Printf("(%.1f)x^2 + (%.1f)*x + (%.1f) = 0\n",a , b, c)
  obtener_raices(a,b,c)
  a = 1
  b = 3
  c = 2
  fmt.Printf("\n(%.1f)x^2 + (%.1f)*x + (%.1f) = 0\n",a , b, c)
  obtener_raices(a,b,c)
  a = 1
  b = 2
  c = 1
  fmt.Printf("\n(%.1f)x^2 + (%.1f)*x + (%.1f) = 0\n",a , b, c)
  obtener_raices(a,b,c)
  a = 2
  b = 3
  c = 1
  fmt.Printf("\n(%.1f)x^2 + (%.1f)*x + (%.1f) = 0\n",a , b, c)
  obtener_raices(a,b,c)
  a = 2.1
  b = 3.2
  c = 1.7
  fmt.Printf("\n(%.1f)x^2 + (%.1f)*x + (%.1f) = 0\n",a , b, c)
  obtener_raices(a,b,c)
  a = 2
  b = 2
  c = 0.5
  fmt.Printf("\n(%.1f)x^2 + (%.1f)*x + (%.1f) = 0\n",a , b, c)
  obtener_raices(a,b,c)
}

func main() {
  var a float64
  var b float64
  var c float64
  fmt.Printf("Ingrese los coeficientes:\n")
  fmt.Printf("a -> ")
  fmt.Scanln(&a)
  fmt.Printf("b -> ")
  fmt.Scanln(&b)
  fmt.Printf("c -> ")
  fmt.Scanln(&c)
  fmt.Printf("\nEcuacion: (%.1f)x^2 + (%.1f)*x + (%.1f) = 0\n",a , b, c)
  fmt.Printf("\nRaices de la ecuacion ->\n")
  obtener_raices(a, b, c)

}
