package main

import (
	"banco/src/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)

}

type verificarConta interface {
	Sacar(valorDoBoleto float64) string
}

func main() {
	contaDoMilton := contas.ContaPoupanca{}
	contaDoMilton.Depositar(500)
	PagarBoleto(&contaDoMilton, 200)
	fmt.Println(contaDoMilton.ObterSaldo())

	contaDolucas := contas.ConstaCorrente{}
	contaDolucas.Depositar(500)
	PagarBoleto(&contaDolucas, 200)
	fmt.Println(contaDolucas.ObterSaldo())
}
