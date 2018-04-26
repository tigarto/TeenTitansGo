/*
Enunciado del problema:
Escribir un programa que lea dos números x y n y en una función, calcule la suma de la progresión geométrica:

1 + x + x^2 + x^3 + x^4 + ⋯ + x^n

Comando de ejecucion:

go run sol_problema2.go

*/

package main

import "fmt"
import "math"

/*
Determina si el numero num1 es multiplo del num2, es decir
si num1 es varias veces num2 (num1=N*num2)
*/
func calcularSerie(x float64, n int) float64 {
  var r float64 = 0
  for i := 0; i <= n; i++ {
    r = r + math.Pow(x,float64(i))
  }
  return r
}

func main() {
  var x, s float64
  var n int
  fmt.Print("Ingrese el valor de x: ")
  fmt.Scanln(&x)
  fmt.Print("Ingrese el valor de n: ")
  fmt.Scanln(&n)
  s = calcularSerie(x,n)
  fmt.Println("\n1 + x + x^2 + x^3 + x^4 + ⋯ + x^n =",s)
}
