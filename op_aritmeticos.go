package main;

import "fmt";

func main() {
	//declarar las variables.
	var(
		a int
		b int
	)
	//pedir la entrada de datos al usuario.
	fmt.Println("Introduce un numero: ");
	fmt.Scanln(&a);
	fmt.Println("Introduce otro numero: ");
	fmt.Scanln(&b);
	c := a + b
	//imprimimos los valores de la suma.
	fmt.Printf("\nSuma:");
	fmt.Printf("\nEl resultado de %d + %d es igual a = %d",a, b, c);
	//imprimimos los valores de la resta.
	fmt.Printf("\nResta:");
	d := a - b
	fmt.Printf("\nEl resultado de %d - %d es igual a = %d", a, b, d );
	//imprimos los valores de la multiplicacion.
	fmt.Printf("\nMultiplicacion:");
	e := a * b
	fmt.Printf("\nEl resultado de %d * %d es igual a = %d", a, b, e);
	//imprimimos los valores de la division.
	fmt.Printf("\nDivision:");
	f := a / b
	fmt.Printf("\nEl resultado de %d / %d es igual a = %d", a, b, f);
}
