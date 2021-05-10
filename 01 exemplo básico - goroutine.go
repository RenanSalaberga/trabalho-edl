// este é um programa simples cujo objetivo é mostrar algumas das possibilidades que temos com as goroutines
// são criadas duas funções que imprimem conteúdo na tela do usuário
// uma sendo executada a cada 500 milisegundos e outra a cada 2 segundos

package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string) // canal de comunicação 1
	c2 := make(chan string) // canal de comunicação 2

	go func() {
		for {
			c1 <- "A cada 500ms"               // envia a mensagem entre aspas pelo canal c1
			time.Sleep(time.Millisecond * 500) // pausa a execução por 500 milisegundos
		}
	}() // essa sintaxe permite declarar uma função e já executá-la

	go func() {
		for {
			c2 <- "A cada 2 segundos"   // envia a mensagem entre aspas pelo canal c2
			time.Sleep(time.Second * 2) // pausa a execução por 2 segundos
		}
	}()

	for {
		select { // o comando select executa o bloco de código do Canal que está disponível no momento
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

}
