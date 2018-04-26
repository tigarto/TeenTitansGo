
package main

import (
	"fmt"
	"math/rand"
  "time"
)

func random(min, max int) int {
    return rand.Intn(max - min) + min
}

func lanzarMondena(numL int)  {
  for i := 0; i < numL; i++ {
    if random(0, 2) == 0 {
      fmt.Print("C")
    } else {
      fmt.Print("S")
    }
  }
  fmt.Print("\n")
}

func main() {
  //var N, M, cant, i, aleatorio int
  l := 0
  fmt.Print("Digite el numero de lanzamientos: ");
  fmt.Scanln(&l)
  rand.Seed(time.Now().Unix()) // semilla
  lanzarMondena(l)
}
