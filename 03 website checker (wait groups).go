// este exemplo é um pequeno programa que ao ser executado verifica se uma lista de websites estão offline ou online
// este programa tem o mesmo próposito do anterior, porém, ao invés de utilizarmos os canais, a concorrência é implementada através dos wait groups

package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // declaração da variável wg que é um waitGroup, o mecanismo responsável pelo sincronismo deste programa

func main() {

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
		go getWebsite(website) // executa uma goroutine para cada elemento do vetor
		wg.Add(1)              // para cada criação de uma goroutine, adicionaremos 1 ao waitGroup
	}

	wg.Wait() // como a main em Go também é considerada uma rotina, usamos esse comando para que o programa não se encerre de forma abrupta
	// caso não usássemos esse comando, o programa se encerraria quando terminasse de executar todos os comandos contidos na main

}
func getWebsite(website string) {
	defer wg.Done()                                // a cada execução de uma goroutine, decrementamos o valor do waitGroup em 1
	if res, err := http.Get(website); err != nil { // se não houver resposta, o website está offline
		fmt.Println(website, "is down") // imprime a mensagem na tela do usuário

	} else { // se ele entrou neste caso, é porque houve resposta do website, logo, ele está online
		fmt.Printf("[%d] %s is up\n", res.StatusCode, website) // imprime a mensagem na tela do usuário
	}

}
