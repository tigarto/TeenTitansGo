
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

func generar_respuestas(tipo int)  {
	num resp := random(0,4)
	if tipo == 0 {
		// Respuestas correctas
		switch resp {
		case 0:
			fmt.Printf("Muy bien!")
		case 1:
			fmt.Printf("Excelente!")
		case 2:
			fmt.Printf("Buen trabajo!")
		case 3:
			fmt.Printf("Sigue haciéndolo bien!")
		}
	}
	else {
		// Respuestas incorrectas
		switch resp {
		case 0:
			fmt.Printf("No. Por favor trata de nuevo.")
		case 1:
			fmt.Printf("Incorrecto. Trata una vez más.")
		case 2:
			fmt.Printf("No te rindas!")
		case 3:
			fmt.Printf("No. Trata de nuevo")
		}
	}
}

func main() {
	const CORRECTO = 0
	const INCORRECTO = 1
  var salir string = "n"
  for salir == "n" {
		rand.Seed(time.Now().Unix())
    num1 := generar_digito()
    num2 := generar_digito()
    var resp int
    var ban bool = false
    fmt.Printf("¿Cuantas %d veces %d? -> ",num1,num2);
    fmt.Scanln(&resp)
    for ban == false {
			rand.Seed(time.Now().Unix())
      if resp == num1*num2 {
        generar_respuestas(CORRECTO)
        ban = true
      } else {
				generar_respuestas(INCORRECTO)
        fmt.Printf("¿Cuantas %d veces %d? -> ",num1,num2);
        fmt.Scanln(&resp)
      }
    }
    fmt.Printf("¿Desea salir (y/n)? ");
    fmt.Scanln(&salir)
  }
}
