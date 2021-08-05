package main
/* Trabajando con condicionales en Go
if -- else..*/
import "fmt"

func main() {
  //declaracion de variables.
  var (
    edad int
    nombre string
  )
  //pedida de entrada de datos.
  fmt.Printf("Introduce Tu nombre: \n")
  fmt.Scanln(&nombre)
  fmt.Printf("Introduce tu edad aqui: \n")
  fmt.Scanln(&edad)
  if edad >= 18 {
    fmt.Printf("Tu nombre es: %s \nTu edad: %d \nDisfruta de nuestra plataforma!!", nombre, edad)
  } else {
    fmt.Printf("Lo sentimos!! \nEres menor de edad y no puedes ingresar.")
  }
}
