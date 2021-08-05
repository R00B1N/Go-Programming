package main;

import (
	"fmt"
	"math"
)

func main() {
	var(
		n1 float64
		n2 float64
	)
	fmt.Printf("\t\t// Funciones Matematicas //  \n\n");
	fmt.Printf("\nIntroduce un numero: ");
	fmt.Scanln(&n1);
	fmt.Printf("\nIntroduce otro numero: ");
	fmt.Scanln(&n2);
	//elevando a la raiz cuadrada.
	c := math.Pow(n1, 2);
	fmt.Printf("%.2f^%d = %.2f \n", n1, 2, c);
	//imprimiendo el seno.
	c = math.Sin(n1);
	fmt.Printf("Sin(%.2f) = %.2f \n", n1, c);
	//imprimiendo el coseno.
	c = math.Cos(n2);
	fmt.Printf("Cos(%.2f) = %.2f \n", n2, c);
	//calcular la raiz cuadrada de un numero.
	c = math.Sqrt(n1 * n2);
	fmt.Printf("Sqrt(%.2f * %.2f) = %.2f \n", n1, n2, c);
}
