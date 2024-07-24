package main

import (
	"fmt"
	"go/printer"
	"net/http"
	"os"
)

func main() {
	
	exibeIntroducao()
	exibeMenu()

	comando := leComando()
	
	//Scanf é a função para receber dados do usuário 
	//& - serve para indicar que vai salvar na variável oque o usuário digitar


	switch comando {
	case 1:
		iniciarMonitoramendo()
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Não conheço esse comando...")
		os.Exit(-1)
	}
	//O switch não precisa do "break" em Go	
	
}


func exibeMenu(){

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func exibeIntroducao(){
	nome := "Lethicia"
	versao := 1.1 
	fmt.Println("Olá sra", nome)
	fmt.Println("Este programa está na versão", versao)
}

func leComando() int{
	var comandoLido int
	fmt.Scan(&comandoLido) 
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramendo(){
	fmt.Println("Monitorando...")
	site := "https://www.alura.com.br"
	resp, _ := http.Get(site)
	//fmt.Println(resp)

	if resp.StatusCode == 200{
		fmt.Println("Site: ", site, "foi carregado com sucesso")
	}else{
		fmt.Println("Site: ", site, "está com problemas. Status Code: ", resp.StatusCode)
	}
}