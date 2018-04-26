/*
Enunciado del problema:
Escriba una función llamada múltiplo que tome dos enteros y
determine si el segundo es múltiplo del primero. La función
deberá tomar dos argumentos enteros y devolver 1 si el segundo
es un múltiplo del primero y 0 si no. Luego, utilice esta
función en un programa que acepte como entrada una serie de
pares de enteros.

Comando de ejecucion:

go run sol_problema1.go

*/

package main

import "fmt"

/*
Determina si el numero num1 es multiplo del num2, es decir
si num1 es varias veces num2 (num1=N*num2)
*/
func esMultiplo(num1, num2 int) int {
  if num1%num2 == 0 {
    return 1
  } else {
    return 0
  }
}

func main() {
    var N1, N2 int
    fmt.Print("Ingrese el primer numero: ")
    fmt.Scanln(&N1)
    fmt.Print("Ingrese el segundo numero: ")
    fmt.Scanln(&N2)
    fmt.Println()
    if esMultiplo(N2,N1) == 1 {
      //fmt.Println("%d es multiplo de %d",N2,N1)
      fmt.Println(N2,"es multiplo de",N1)
    } else {
      fmt.Println(N2,"no es multiplo de",N1)
    }
}
