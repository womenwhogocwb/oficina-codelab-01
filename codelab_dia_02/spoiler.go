package main

import "fmt"

func main() {
	fmt.Println("<<< Dia 02 Exercício #01 >>>")
	x := 90
	y := 75
	printaMaior(x, y)

	fmt.Println("\n<<< Dia 02 Exercício #02 >>>")
	anoNascimento := 1980
	anoAtual := 2019
	idade := anoAtual - anoNascimento

	if anoAtual >= anoNascimento {
		if idade >= 16 {
			fmt.Printf("essa pessoa tem %d anos e já pode votar.\n", idade)
		} else if idade < 16 && idade >= 0 {
			fmt.Printf("essa pessoa tem %d anos e ainda não pode votar.\n", idade)
		}
	} else {
		fmt.Println("por favor, verifique os valores fornecidos e tente novamente.")
	}

	fmt.Println("\n<<< Dia 02 Exercício #03 >>>")
	var senha string
	senha = "womenwhogo"

	var entrada string
	entrada = "golang"

	if entrada != "" && senha != "" {
		if entrada == senha {
			fmt.Println("ACESSO PERMITIDO")
		} else {
			fmt.Println("ACESSO NEGADO")
		}
	} else if entrada == "" && senha != "" {
		fmt.Println("por favor, entre sua senha.")
	} else if senha == "" {
		fmt.Println("por favor, defina uma senha.")
	} else {
		fmt.Println("deu xabu, melhor investigar")
	}

	fmt.Println("\n<<< Dia 02 Exercício #04 >>>")
	umaVariável := "olar"
	tipoDeUmaVariável := fmt.Sprintf("%T", umaVariável)

	fmt.Printf("valor: %v -> ", umaVariável)

	if tipoDeUmaVariável == "string" {
		fmt.Println("a variável é uma string!")
	} else if tipoDeUmaVariável == "int" {
		fmt.Println("a variável é um int!")
	} else if tipoDeUmaVariável == "float64" || tipoDeUmaVariável == "float32" {
		fmt.Println("a variável é um float!")
	} else if tipoDeUmaVariável == "bool" {
		fmt.Println("a variável é um bool!")
	} else {
		fmt.Println("não sabemos o tipo dessa variável :D")
	}
}

func printaMaior(x int, y int) {
	fmt.Print("O maior número é: ")
	if x > y {
		fmt.Println(x)
	} else if y > x {
		fmt.Println(y)
	} else {
		fmt.Println("x e y são iguais")
	}
}