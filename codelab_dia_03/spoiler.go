package main

import (
	"fmt"
)

func main()  {
	fmt.Println("<<< Dia 03 Exercício #01 >>>\n")
	for inteiro01 := 1; inteiro01 <= 100; inteiro01++ {
		if inteiro01 % 3 == 0 {
			fmt.Println(inteiro01)
		}
	}

	fmt.Println("\n<<< Dia 03 Exercício #02 >>>\n")
	for inteiro02 := 1; inteiro02 <= 100; inteiro02++ {
		if inteiro02 % 9 != 0 {
			continue
		}
		fmt.Println(inteiro02)
	}

	fmt.Println("\n<<< Dia 03 Exercício #03 >>>\n")
	verificaMultiploDeCinco(88948576)

	fmt.Println("\n<<< Dia 03 Exercício #04 >>>\n")
	somaTres(5, 27, 330)

	fmt.Println("\n<<< Dia 03 Exercício #05 >>>\n")
	somaN(3, 56, 89, 2,10)
}

func somaTres(x, y, z int) {
	soma := x + y + z
	fmt.Println(soma)
}

func somaN(num ...int) {
	soma := 0
	for _, operando := range num {
		soma = soma + operando
	}
	fmt.Printf("a soma dos %d operandos deu: %d\n", len(num), soma)
}

func verificaMultiploDeCinco(numero int)  {
	if numero % 5 == 0 {
		fmt.Printf("o número %d é múltiplo de 5\n", numero)
		return
	}
	fmt.Printf("o número %d não é múltiplo de 5\n", numero)
}
