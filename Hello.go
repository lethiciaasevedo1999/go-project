package main

import "fmt"
import "reflect" // pacote utilizado para saber qual é o tipo da variável utilizada

func main() {
	var nome = "Lethicia" // ou nome := “Lethicia” 
	var idade  = 25 // ou idade := 25
	var versao  = 1.1 // ou versao := 1.2
	fmt.Println("Olá",nome, "sua idade é:", idade, "anos")
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("O tipo da variável nome é:", reflect.TypeOf(nome))
}