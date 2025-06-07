package main

import (
	"fmt"
)

func inverte(amudar bool) bool {
	return !amudar
}
func soma(a int, b int) int {
	return a + b
}
func main() {
	var numero uint = 10
	var texto string = "Olá, mundo!"
	var pi float32 = 3.14159265358979323846
	var pi64 float64 = 3.14159265358979323846
	var booleano bool = true
	var negativo int = -10

	fmt.Println("numero:", numero, "texto:", texto, "pi:", pi, "pi64:", pi64, "booleano:", booleano, "negativo:")
	fmt.Println("========================")
	fmt.Printf("numero: %d, texto: %s, pi: %f, pi64: %f, booleano: %t, negativo: %d\n", numero, texto, pi, pi64, booleano, negativo)
	fmt.Println("========================")
	fmt.Println(numero)
	fmt.Println(texto)
	fmt.Println(pi)
	fmt.Println(pi64)
	fmt.Println(booleano)
	fmt.Println(negativo)
	fmt.Println("========================")
	var estaChovendo bool = true
	var temSol bool = false
	fmt.Println(estaChovendo)
	fmt.Println(temSol)
	fmt.Println("========================")
	inverte(estaChovendo)
	fmt.Println(estaChovendo)
	inverte(temSol)
	fmt.Println(temSol)
	fmt.Println("========================")

	var numeros []int = []int{1, 2, 3, 4, 5}
	fmt.Println(numeros)
	fmt.Println("========================")
	for i := 0; i < len(numeros); i++ {
		fmt.Println(numeros[i])
	}
	fmt.Println("========================")
	var palavras []string = []string{"uma", "palavra", "qualquer"}
	fmt.Println(palavras)
	fmt.Println("========================")
	for i := 0; i < len(palavras); i++ {
		fmt.Println(palavras[i])
	}
	fmt.Println("========================")
	palavras = append(palavras, "depois")
	for i := 0; i < len(palavras); i++ {
		fmt.Println(palavras[i])
	}
	fmt.Println("========================")
	fmt.Println("========================")

	var inputed int
	fmt.Print("Type a number: ")
	fmt.Scanln(&inputed)
	fmt.Println("Your number is:", inputed)
	fmt.Println("========================")

	var inputedNome string
	fmt.Print("Digite o nome: ")
	fmt.Scanln(&inputedNome)
	fmt.Println("Olá ", inputedNome)

	fmt.Println("========================")
	var inputedHora int
	fmt.Println("Digite a hora: ")
	fmt.Scanln(&inputedHora)
	// hora := 12
	hora := inputedHora
	if hora < 12 {
		fmt.Println("Bom dia!")
	} else if hora < 18 {
		fmt.Println("Boa tarde!")
	} else {
		fmt.Println("Boa noite!")
	}

	fmt.Println("========================")
	day := 5
	switch day {
	case 1:
		fmt.Println("Inicio de Semana")
	case 2, 3, 4, 5:
		fmt.Println("Dia Util")
	case 7:
		fmt.Println("Final de Semana")
	default:
		fmt.Println("Dia inválido")
	}
	fmt.Println("========================")
	var inputedSenha string
	fmt.Println("Digite a senha: ")
	fmt.Scanln(&inputedSenha)

	for inputedSenha != "123" {
		fmt.Println("Senha incorreta")
		fmt.Println("Digite a senha: ")
		fmt.Scanln(&inputedSenha)
	}
	if inputedSenha == "123" {
		fmt.Println("Acesso permitido")
	} else {
		fmt.Println("Acesso negado")
	}
	fmt.Println("========================")
}
