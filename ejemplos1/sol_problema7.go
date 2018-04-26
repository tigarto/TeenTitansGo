
package main

import (
	"fmt"
	"math/rand"
  "time"
)

func random(min, max int) int {
    return rand.Intn(max - min) + min
}

func generar_digito() int {
  return random(0, 10)
}

func main() {
  rand.Seed(time.Now().Unix())
  var salir string = "n"
  for salir == "n" {
    num1 := generar_digito()
    num2 := generar_digito()
    var resp int
    var ban bool = false
    fmt.Printf("¿Cuantas %d veces %d? -> ",num1,num2);
    fmt.Scanln(&resp)
    for ban == false {
      if resp == num1*num2 {
        fmt.Printf("Muy bien!\n")
        ban = true
      } else {
        fmt.Printf("No. Por favor intenta nuevamente\n")
        fmt.Printf("¿Cuantas %d veces %d? -> ",num1,num2);
        fmt.Scanln(&resp)
      }
    }
    fmt.Printf("¿Desea salir (y/n)? ");
    fmt.Scanln(&salir)
  }
}
