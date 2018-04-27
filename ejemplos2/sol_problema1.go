package main

import (
	"fmt"
  "strings"
)

func contarCaracter(v string, c string) int {
  cnt := 0
  for _,e := range(v) {
    if string(e) == c {
      cnt = cnt + 1
    }
  }
  return cnt
}

/*
func longitud(v string) int {
  i := 0
  cnt := 0
  for v[i] != 0 {
    cnt = cnt + 1
    i = i + 1
    fmt.Println(i)
    fmt.Println(string(v[i]))
  }
  return cnt
}
*/

func main() {
  var s string = "hohhhhhhhhla"
  fmt.Println(contarCaracter(s,"h"))
  fmt.Println(strings.Count(s,"h"))
  fmt.Println(len(s))
}
