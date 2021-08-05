package main

import "fmt"

func main() {
  //declaracion de variables.
  var (
    n1 int
    n2 int
  )
  //pedida de entrada de datos.
  fmt.Printf("Introduce un numero: \n")
  fmt.Scanln(&n1)
  fmt.Printf("Introduce otro numero: \n")
  fmt.Scanln(&n2)
  //llamar a nuestra funcion llamada verifica.
  fmt.Println(n1 == n2)
  fmt.Println(n1 != n2)
  fmt.Println(n1 < n2)
  fmt.Println(n1 > n2)
  fmt.Println(n1 >= n2)
  fmt.Println(n1 <= n2)
}
