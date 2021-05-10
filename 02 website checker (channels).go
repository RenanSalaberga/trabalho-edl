// este exemplo é um pequeno programa que ao ser executado verifica se uma lista de websites estão offline ou online
// nele conseguimos ver a execução das goroutines, bem como o uso dos canais, para receber as mensagens uma vez que elas estão prontas

package main

import (
	"fmt"
	"net/http"
)

func main() {

	c := make(chan string) // channel utilizado para a comunicação

	websites := []string{ // lista de websites que serão checados
		"https://www.uerj.br/",
		"https://stackoverflow.com/",
		"https://github.com/",
		"https://gitlab.com/",
		"http://medium.com/",
		"https://golang.org/",
		"https://www.udemy.com/",
		"https://www.coursera.org/",
	}
	for _, website := range websites { // percorre o vetor websites
		go getWebsite(website, c) // executa uma goroutine para cada elemento do vetor
	}

	for msg := range c { // conforme as mensagens são recebidas no canal c este loop é responsável exibir a mensagem na tela para o usuário
		fmt.Println(msg) // a mensagem transmitida pelo canal c é guardada na variável msg
	}
}
func getWebsite(website string, c chan string) {
	if _, err := http.Get(website); err != nil { // se não houver resposta, o website está offline
		c <- website + " is down" // envia a mensagem pelo canal c

	} else { // se ele entrou neste caso, é porque houve resposta do website, logo, ele está online
		c <- website + " is up" // envia a mensagem pelo canal c
	}

}
