package main

import "fmt"

func main() {
	nome := "Lethicia"
	versao := 1.1 
	fmt.Println("Olá sra", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	var comando int
	fmt.Scan(&comando) 
	fmt.Println("O endereço da minha varia'vel comando é", &comando)
	fmt.Println("O comando escolhido foi", comando)
	//Scanf é a função para receber dados do usuário 
	//& - serve para indicar que vai salvar na variável oque o usuário digitar
}