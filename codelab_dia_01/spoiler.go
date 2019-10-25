package main

import "fmt"

func main() {

	fmt.Println("\n<<< Dia 01 Exercício #01 >>>")
	var numero01 int
	numero01 = 47
	fmt.Println("\nvalor 1:", numero01)

	var numero02 int
	numero02 = 60031
	fmt.Println("valor 2:", numero02)

	var numero03 int
	numero03 = 13
	fmt.Println("valor 3:", numero03)

	var numero04 int
	numero04 = 7
	fmt.Println("valor 4:", numero04)

	fmt.Println("\n<<< Dia 01 Exercício #02 >>>")
	A := 230
	B := 27
	somaAB := A + B
	subtraçãoAB := A - B
	fmt.Println("\nA:", A)
	fmt.Println("B:", B)
	fmt.Println("soma:", somaAB)
	fmt.Println("subtração:", subtraçãoAB)

	fmt.Println("\n<<< Dia 01 Exercício #03 >>>")
	multiplicaçãoAB := A * B
	divisãoAB := float64(A) / float64(B)
	fmt.Println("\nmultiplicação:", multiplicaçãoAB)
	fmt.Println("divisão:", divisãoAB)

	fmt.Println("\n<<< Dia 01 Exercício #04 >>>")
	fmt.Printf("\nOs tipos de A e B são, respectivamente: %T e %T", A, B)
	fmt.Printf("\nOs tipos de somaAB e subtraçãoAB são, respectivamente: %T e %T", somaAB, subtraçãoAB)
	fmt.Printf("\nOs tipos de multiplicaçãoAB e divisãoAB são, respectivamente: %T e %T\n", multiplicaçãoAB, divisãoAB)

	fmt.Println("\n<<< Dia 01 Exercício #05 >>>")
	var preçoDoPão float64 = 4.60
	var preçoDaAveia float64 = 5
	var preçoDoAzeite float64 = 19.95
	var preçoDoSuco float64 = 7
	var preçoDaÁgua float64 = 2.15
	var preçoDoKGMaçã float64 = 8.90
	var preçoDoKGBanana float64 = 4.5

	valorTotal := (preçoDoPão * float64(3)) + (preçoDaAveia * float64(0)) + (preçoDoAzeite * float64(1)) + (preçoDoSuco * float64(2)) + (preçoDaÁgua * float64(5)) + (preçoDoKGMaçã * float64(1.5)) + (preçoDoKGBanana * float64(0))

	fmt.Println("\n\nValor Total da Compra: R$", valorTotal)

	fmt.Println("\n<<< Dia 01 Exercício #06 >>>")
	fmt.Println(`> A palavra "var" é uma keyword e não pode ser utilizada como identificador.
As keywords da linguagem podem ser utilizadas apenas para os seus fins.
Por exemplo, a finalidade de "var" é declarar uma variável, como no código acima.
Todas as keywords de Go estão expostas neste link:
- https://golang.org/ref/spec#Keywords`)
	var umNumero int
	umNumero = 27
	fmt.Println("\nvalor da variável umNumero:", umNumero)

	fmt.Println("\n<<< Dia 01 Exercício #07 >>>")
	X := 15
	var Y int = 31
	var Z int
	Z = 47

	fmt.Println("\n\t X \t Y \t Z")
	fmt.Printf("base 2 : %b \t%b \t%b", X, Y, Z)
	fmt.Printf("\nbase 10: %d \t%d \t%d", X, Y, Z)
	fmt.Printf("\nbase 16: %#X \t%#X \t%#X\n", X, Y, Z)

	fmt.Println("\n<<< Dia 01 Exercício #08 >>>")
	var x int
	fmt.Printf("\n\nvalor de x: %d, tipo de x: %T\n", x, x)

	fmt.Println(`> Quando uma variável é criada e nenhum valor é explicitamente
atribuído a ela, ela recebe um valor padrão.
Esse valor depende do tipo da variável e se chama "valor zero" em Go.
Para o tipo int o valor zero é 0, como podemos ver ao dar Run nesse código.`)
}
