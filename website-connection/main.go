package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const urlVerificationDelay = 5 * time.Second

func welcome() {
	nome := "Terráqueo"
	version := 1.1

	fmt.Printf("Olá, %s!\n", nome)
	fmt.Printf("Este programa está na versão %v\n", version)
	fmt.Println("")
}

func showOptions() {
	fmt.Println("O que você deseja fazer?")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
	fmt.Println("")
}

func getCommand() int {
	var command int
	fmt.Scan(&command)

	return command
}

func verifyWebsite(url string) {
	resp, err := http.Get(url)
	
	if err != nil {
		fmt.Printf("Ocorreu um erro, tente novamente.")
		os.Exit(1)
		return
	}

	if resp.StatusCode != 200  {
		fmt.Printf("O site está com problemas. Código de resposta: %v.\n\n", resp.StatusCode)
		return
	}
	
	fmt.Printf("O site foi carregado com sucesso. Código de resposta: %v.\n\n", resp.StatusCode)
}

func getWebsitesUrl() []string {
	file, err := os.Open("./sites.txt")

	if err != nil {
		fmt.Println("Erro ao ler arquivo, tente novamente.")
		os.Exit(-1)
	}
	
	reader :=	bufio.NewReader(file)

	var sites []string

	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}
		
		if err != nil {
			fmt.Println("Erro ao ler sites do arquivo, tente novamente.")
			os.Exit(-1)
		}

		line = strings.TrimSpace(line)
		sites = append(sites, line)
	}

	return sites
}

func startMonitoring() {
	fmt.Printf("Iniciando monitoramento...\n\n")

	sites := getWebsitesUrl()

	for i := 0; i < 3; i++ {
		for _, url := range(sites) {
			fmt.Printf("Verificando a url %s...\n", url)
			
			verifyWebsite(url)
		}

		time.Sleep(urlVerificationDelay)
	}
}

func main() {
	welcome()

	for {
		showOptions()
	
		command := getCommand()
	
		switch command {
			case 0:
				fmt.Println("Programa encerrado.")
				os.Exit(0)
			case 1:
				startMonitoring()
			case 2:
				fmt.Println("Exibindo logs...")
			default:
				fmt.Println("Comando não reconhecido, tente novamente.")
				os.Exit(-1)
		}
	}
}