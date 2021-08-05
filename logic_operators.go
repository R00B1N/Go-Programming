package main
/*Operadores logicos en Go
"&& -> and"
"|| -> or"
"! -> not"
*/

import "fmt"

func main() {
  //declaracion de variables.
  var (
    num1 int
    num2 int
  )
  //pedido de entrada de datos.
  fmt.Printf("Introduce un numero: ")
  fmt.Scanln(&num1)
  fmt.Printf("Introduce otro numero: ")
  fmt.Scanln(&num2)
  //arrojar los resultados.
  fmt.Println(num1 > num2 && num1 != num2)
  fmt.Println(!(num1 >= num2))
  fmt.Println(num1 == num2 || num1 > num2)
}
