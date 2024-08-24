package contas

import "banco/src/clientes"

type ConstaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ConstaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "saque realizado conm sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ConstaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "O valor do deposito não pode ser negativo", c.saldo
	}

}

func (c *ConstaCorrente) Tranferir(valorDaTransferencia float64, constaDestino *ConstaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		constaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func (c *ConstaCorrente) ObterSaldo() float64 {
	return c.saldo
}
