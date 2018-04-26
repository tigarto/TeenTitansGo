
package main

import (
	"fmt"
	"math/rand"
  "time"
)

func main() {
  //var N, M, cant, i, aleatorio int
  var N, M, cant, aleatorio int
  fmt.Print("Digite la cantidad de números que desea generar: ");
  fmt.Scanln(&cant);
  fmt.Print("Digite los limites (primero el superior, luego el inferior): ");
  fmt.Scanf("%d",&N);
  fmt.Scanf("%d",&M);
  rand.Seed(time.Now().Unix()) // semilla
  for i := 0; i < cant; i++ {
    aleatorio = rand.Intn(N - M) + M
    fmt.Println(aleatorio)
  }
}
