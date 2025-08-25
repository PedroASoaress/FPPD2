// por Fernando Dotti - PUCRS
// EXERCÍCIO:  dado o programa abaixo
//    1) quantos processos concorrentes são gerados ?
//    sao gerados 40 goroutines ou seja 40 processos são gerados
//    2) execute e observe: que se pode supor sobre a velocidade relativa dos mesmos ?
//    as goroutines nao tem prioridade garantida, entao a velocidade relativa pode variar a cada execucao(não é deterministica), ou seja, alguns podem
//    imprimir mais vezes que outros
// OBSERVACAO:o sleep no método main serve para este nao acabar, o que acabaria todos processos em execucao
//     mais adiante veremos outras formas de sincronizar isto

package main

import (
	"fmt"
	"time"
)

var N int = 40

func geraNespacos(n int) string {
	s := "  "
	for j := 0; j < n; j++ {
		s = s + "   "
	}
	return s
}

func funcaoA(id int, s string) {
	for {
		fmt.Println(s, id)
	}
}
func main() {
	for i := 0; i < N; i++ {
		go funcaoA(i, geraNespacos(i))
	}
	for true {
		time.Sleep(100 * time.Millisecond)
	}
}
