
package main

import (
	"fmt"
	"math/rand"
  "time"
)

// Gracias a: http://golangcookbook.blogspot.com.co/2012/11/generate-random-number-in-given-range.html
func random(min, max int) int {
    return rand.Intn(max - min) + min
}


func main() {
  //var N, M, cant, i, aleatorio int
  var N, M, cant, aleatorio int
  fmt.Print("Digite la cantidad de números que desea generar: ");
  fmt.Scanln(&cant);
  fmt.Print("Digite los limites (primero el superior, luego el inferior): ");
  fmt.Scanf("%d",&N);
  fmt.Scanf("%d",&M);
  rand.Seed(time.Now().Unix())
  for i := 0; i < cant; i++ {
    aleatorio = random(M,N)
    fmt.Println(aleatorio)
  }
}
