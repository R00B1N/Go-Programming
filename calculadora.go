package main

import "fmt"

func main(){

	var (
		opcion int
		n1 int
		n2 int
		operacion int
	)

	//Creando mi primera calculadora en Go

	fmt.Println("Hola, Bienvenido a mi programa basico de calculadora.")

	fmt.Print("\n// Menu //")
	fmt.Print("\n1-Sumar. \n2-Restar. \n3-Multiplicar. \n4-Dividir.")
	fmt.Print("\nEscoge una opcion: ")
	fmt.Scanln(&opcion)

	if opcion == 1{
		fmt.Print("\nVamos a sumar!!")
		fmt.Print("\nIntroduce el primer numero para nuestra suma: ")
		fmt.Scanln(&n1)
		fmt.Print("\nIntroduce el segundo numero para nuestra suma: ")
		fmt.Scanln(&n2)
		operacion = n1 + n2
		fmt.Println("La suma es igual a: ", operacion)
	}else if opcion == 2{
		fmt.Print("\nVamos a restar!!")
		fmt.Print("\nIntroduce el primer numero para nuestra resta: ")
		fmt.Scanln(&n1)
		fmt.Print("\nIntroduce el segundo numero para nuestra resta: ")
		fmt.Scanln(&n2)
		operacion = n1 - n2
		fmt.Println("La resta es igual a: ", operacion)
	}else if opcion == 3{
		fmt.Print("\nVamos a restar!!")
		fmt.Print("\nIntroduce el primer numero para nuestra multiplicacion: ")
		fmt.Scanln(&n1)
		fmt.Print("\nIntroduce el segundo numero para nuestra multiplicacion: ")
		fmt.Scanln(&n2)
		operacion = n1 * n2
		fmt.Println("La multiplicacion es igual a: ", operacion)
	}else {
		if opcion == 4{
			fmt.Print("\nVamos a restar!!")
			fmt.Print("\nIntroduce el primer numero para nuestra division: ")
			fmt.Scanln(&n1)
			fmt.Print("\nIntroduce el segundo numero para nuestra division: ")
			fmt.Scanln(&n2)
			operacion = n1 / n2
			fmt.Println("La division es igual a: ", operacion)
		}
	}
}
