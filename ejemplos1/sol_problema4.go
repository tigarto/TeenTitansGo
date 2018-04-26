/*
Escriba una función que tome un valor entero y devuelva el
número con sus dígitos en reversa. Por ejemplo, dado el
numero 7631, la función deberá devolver 1367.
*/

package main

import (
	"fmt"
)


func voltearNumero(num int) int {
  inv_num := 0
  for num > 0 {
    inv_num = inv_num*10
		inv_num = inv_num + num%10
    num = num/10
	}
  return inv_num
}

func main() {
  num := 0
  fmt.Print("Digite el numero: ");
  fmt.Scanln(&num);
  fmt.Println("El numero invertido es",voltearNumero(num));
}
