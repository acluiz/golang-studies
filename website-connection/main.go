package main

import (
	"fmt"
	"os"
)

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


func main() {
	welcome()
	showOptions()

	command := getCommand()

	switch command {
		case 0:
			fmt.Println("Programa encerrado.")
			os.Exit(0)
		case 1:
			fmt.Println("Iniciando execução...")
		case 2:
			fmt.Println("Iniciando execução...")
		default:
			fmt.Println("Comando não reconhecido, tente novamente.")
			os.Exit(-1)
	}
}