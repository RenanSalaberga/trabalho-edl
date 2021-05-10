// programa que calcula os primeiros N números da sequência de fibonacci
// neste exemplo cada número será calculado por uma goroutine diferente
// se tivermos um número alto de cores no nosso processador, poderemos criar mais goroutines, melhorando a performance do programa

package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 100)    // canal responsável informar o n-ésimo número que deverá ser calculado
	results := make(chan int, 100) // canal responsável por imprimir o resultado na tela do usuário conforme eles vão ficando prontos
	// ambos os valores de 100 representam o tamanho do buffer

	// o ideal é que você chame a quantidade de workers que você possui de cores na sua CPU
	// desta maneira, cada worker criará uma thread em cada core
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	// este loop é responsável por criar os jobs que serão executados adiante
	// ou seja, os 100 primeiros números da sequência de fibonacci
	// cada n-ésimo número que deseja ser calculado
	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	// este loop imprimirá o resultado dos números conforme eles vão ficando prontos (não necessariamente ele imprimirá a sequência na ordem)
	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}
}

// a função worker recebe dois parâmetros dos canais que criamos
// job informa qual n-ésimo número estamos calculando
// results informa o número de fibonacci referente a esse n-ésimo número
func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs { // este loop percerorré os números que desejamos calcular. neste exemplo, os números de 1 a 100
		results <- fib(n) // conforme os resultados ficarão prontos eles serão transmitidos para o canal results
	}
}

func fib(n int) int { // função clássica para calcular o número de fibonacci utilizando recursão
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
