package main

import (
	"fmt"
)

func main() {
	fmt.Println("<<< Dia 04 Exercício #01 >>>\n")
	var decadas [10]int
	decadas[0] = 00
	decadas[1] = 10
	decadas[2] = 20
	decadas[3] = 30
	decadas[4] = 40
	decadas[5] = 50
	decadas[6] = 60
	decadas[7] = 70
	decadas[8] = 80
	decadas[9] = 90

	fmt.Println(decadas)
	for indice, valor := range decadas {
		fmt.Printf("índice: %d; valor: %d\n", indice, valor)
	}

	fmt.Println("\n<<< Dia 04 Exercício #02 >>>\n")
	conjuntoOperandos:= []int{3, 7, 27, 90, 313, 700}
	somaElementos(conjuntoOperandos)

	fmt.Println("\n<<< Dia 04 Exercício #03 >>>\n")
	aFazer := make([]string, 5)
	aFazer[0] = "Estudar Go"
	aFazer[1] = "Comprar bananas"
	aFazer[2] = "Passar na lavanderia"
	aFazer[3] = "Ir à academia"
	aFazer[4] = "Marcar date"

	for indice, tarefa := range aFazer {
		fmt.Printf("índice %d) %s\n", indice, tarefa)
	}

	fmt.Println("\n<<< Dia 04 Exercício #04 >>>\n")
	numeros := []int{0, 1, 1, 2, 3, 5, 8}
	fmt.Print("elementos da slice <numeros> = ")
	for _, valor := range numeros {
		fmt.Printf("%d ", valor)
	}
	numeros = append(numeros, 13, 21, 34)
	fmt.Print("\nelementos da slice <numeros> após o append = ")
	for _, valor := range numeros {
		fmt.Printf("%d ", valor)
	}

	fmt.Println("\n\n<<< Dia 04 Exercício #05 >>>\n")
	parksAndRecreation := []string{
		"Leslie Knope",
		"Andy Dwyer",
		"April Ludgate",
		"Ron Swanson",
		"Ann Perkins",
		"Ben Wyatt",
		"Chris Traeger",
		"Donna Meagle",
		"Tom Haverford",
	}

	for i := 0; i < len(parksAndRecreation); i++ {
		fmt.Printf("%d.\t%s\n", i+1, parksAndRecreation[i])
	}
}

func somaElementos(operandos []int) {
	soma := 0
	for _, valor := range operandos {
		soma = soma + valor
	}
	fmt.Println("soma dos elementos da slice:", soma)
}
