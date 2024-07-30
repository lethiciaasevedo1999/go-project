package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time" //pacote utilizado para o espaçamento de tempo para executar o looping
)

const monitoramentos = 2
const delay = 5

func main() {

	exibeIntroducao()
	
	registraLog("site-falso", false)



	for {

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

}

func exibeMenu() {

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func exibeIntroducao() {
	nome := "Lethicia"
	versao := 1.1
	fmt.Println("Olá sra", nome)
	fmt.Println("Este programa está na versão", versao)
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println("")

	return comandoLido
}

func iniciarMonitoramendo() {
	fmt.Println("Monitorando...")

	//sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br", "https://caelum.com.br"}

	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos ; i++{
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

		
	fmt.Println("")
}

func testaSite(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, "foi carregado com sucesso")
		registraLog(site, true)
	} else {
		fmt.Println("Site: ", site, "está com problemas. Status Code: ", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")
	

	if err != nil{
		fmt.Println("Ocorreu um erro: ", err)
	}

	
	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		
		sites = append(sites, linha)

		if err == io.EOF{
			break
		}
		
	}
	
	arquivo.Close()

	return sites
}

func registraLog(site string, status bool){
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(arquivo)

}