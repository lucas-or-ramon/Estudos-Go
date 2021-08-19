package main

import (
	"fmt"

	"banco/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) (string, float64)
}

func main() {
	contaExemplo := contas.ContaCorrente{}
	contaExemplo.Depositar(100)
	PagarBoleto(&contaExemplo, 40)

	fmt.Println(contaExemplo.ObterSaldo())
}
