package main;

import "fmt"

func main() {
  //declaramos las variables.
  var a int
  //pedir la entrada de un numero.
  fmt.Printf("Introduzca un numero: \n")
  fmt.Scanln(&a)
  //incrementar
  fmt.Printf("a = %d \n", a)
  a = a + 1
  fmt.Printf("a + 1 = %d \n", a)
  a++
  fmt.Printf("a++ = %d \n", a)
  //disminuir
   a = a - 1
   fmt.Printf("a - 1 = %d \n", a)
   a--
   fmt.Printf("a-- = %d \n", a)
}
