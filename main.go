package main

import "fmt"

func main() {
	menu := true
	for menu {
		var opcion int
		fmt.Println("\n seleccione un numero del menu")
		fmt.Println("1. --If, evaluar numero entero --")
		fmt.Println("2. --Ejemplod de ciclo while --")
		fmt.Println("3. --Ejemplo de ciclo Dowhile --")
		fmt.Println("4. --Ejemplo de ciclo For clasico --")
		fmt.Println("5. --Switch Estaciones del año --")
		fmt.Println("0. --presione enter para salir")
		fmt.Scanf("%d", &opcion)
		switch opcion {
		case 1:
			sentenciaIf()
			break
		case 2:
			simulacionWhile()
			break
		case 3:
			simulacionDoWhile()
			break
		case 4:
			cicloForClassic()

		case 5:
			condicionSwitch()

		default:
			menu = false
		}
	}

}
func sentenciaIf() {
	fmt.Println("##### Condicionales if ####")
	var evaluar int
	fmt.Println("ingrese el valor a evaluar")
	fmt.Scanf("%d", &evaluar)
	if evaluar > 0 {
		fmt.Printf("valor %d positivo, mayor que 0 \n \n", evaluar)
	} else if evaluar < 0 {
		fmt.Printf("valor %d negativo, menor que 0 \n \n", evaluar)
	} else if evaluar == 0 {
		fmt.Printf("valor %d identico a 0 \n \n", evaluar)
	}
}
func simulacionWhile() {
	fmt.Println("##### Ciclo While ####")
	numeroWhile := 0
	for numeroWhile < 3 {
		fmt.Println(numeroWhile)
		numeroWhile++
	}

}
func simulacionDoWhile() {
	fmt.Println("##### Ejemplo de ciclo DoWhile ####")
	numeros := 1
	for isPass := true; isPass; isPass = numeros < 1 {
		fmt.Println(numeros)
		numeros++
	}
}
func cicloForClassic() {
	fmt.Println("##### Ejemplo de ciclo For ####")
	for numeroFor := 0; numeroFor <= 3; numeroFor++ {
		fmt.Println(numeroFor)
	}
}
func condicionSwitch() {
	var estacion string
	fmt.Println("##### Ingresar la estacion del año #####")
	fmt.Scanf("%s", &estacion)
	switch estacion {
	case "primavera":
		fmt.Println("Primavera.")
		fmt.Println("La primavera​ es una de las cuatro estaciones del año, sigue al invierno y precede al verano.")
		break
	case "verano":
		fmt.Println("Verano")
		fmt.Println("El verano o estío es una de las cuatro estaciones de las zonas templadas. Es la más cálida de ellas. Sigue a la primavera y precede al otoño.")
		break
	case "otono":
		fmt.Println("Otoño")
		fmt.Println("El otoño es una de las cuatro estaciones del año y una de las cuatro de las zonas templadas. Sigue al verano y precede al invierno. Astronómicamente, comienza con el equinoccio de otoño y termina con el solsticio de invierno.")
		break
	case "invierno":
		fmt.Println("Invierno")
		fmt.Println("El invierno es una de las cuatro estaciones de las zonas templadas. Sigue al otoño y precede a la primavera. Esta estación se caracteriza por días más cortos, noches más largas y temperaturas más bajas a medida que nos alejamos de la línea ecuatorial.")
		break
	default:
		fmt.Println("no se ingreso ninguna estacion")
	}
}
