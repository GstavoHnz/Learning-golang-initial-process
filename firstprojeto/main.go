package main

import (
	"fmt"
	"net/http"
)
import "os"

func main() {

	setComand := showIntroduction()

	if setComand == 1 {
		fmt.Println("1 - Iniciar monitoramento")
		fmt.Println("2 - Exibir Logs")
		fmt.Println("0 - Sair do Programa ")

		var comando int

		fmt.Scan(&comando)
		fmt.Println("A opção escolhida foi", comando)

		switch comando {
		case 1:
			initMonitoring()
		case 2:
			showLogs()
		case 0:
			fmt.Println("Saindo do Programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	} else {
		fmt.Println("Opção", setComand, "Saindo do sistema")
	}
}

func showIntroduction() int {
	/*
		Programa de monitoramento
	*/
	fmt.Println("Programa de monitoramento")
	const sim int = 1
	const nao int = 0
	var showComand int

	fmt.Scan(&showComand)

	return showComand
}

func initMonitoring() {
	fmt.Println("Monitorando...")
	//Monitoring real site
	site := "https://www.radom-status-code.herokuapp.com"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("O site ", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("O site ", site, "está com indisponivel no momento! Status no momento é ...", resp.StatusCode)
	}
}

func showLogs() {
	fmt.Println("Exibindo Logs")
}
