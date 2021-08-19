package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	if valorDoSaque > 0 && valorDoSaque <= c.saldo {
		c.saldo -= valorDoSaque
		return "Saque Realizado com Sucesso! saldo:", c.saldo
	} else {
		return "saldo Insuficiente. saldo:", c.saldo
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito Feito com sucesso! saldo:", c.saldo
	} else {
		return "Depósito Mal Sucedido. saldo:", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia <= c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}

// func SomandoNumeros(numeros ...int) int {
// 	resultado := 0

// 	for _, numero := range numeros {
// 		resultado += numero
// 	}
// 	return resultado
// }
