package contas

import "banco/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) (string, float64) {
	if valorDoSaque > 0 && valorDoSaque <= c.saldo {
		c.saldo -= valorDoSaque
		return "Saque Realizado com Sucesso! saldo:", c.saldo
	} else {
		return "saldo Insuficiente. saldo:", c.saldo
	}
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito Feito com sucesso! saldo:", c.saldo
	} else {
		return "Depósito Mal Sucedido. saldo:", c.saldo
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
