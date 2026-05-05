package model

import (
	"fmt"
	"time"
)

type Pessoa struct {
	Nome           string
	Endereco       Endereco
	DataNascimento time.Time
	Idade          int
}

func (p *Pessoa) CalculaIdade() {
	anoDeNascimento := p.DataNascimento.Year()
	anoAtual := time.Now().Year()
	p.Idade = anoAtual - anoDeNascimento
	fmt.Printf("A idade de %s é %d anos.\n", p.Nome, p.Idade)
}

func (p Pessoa) IdadeAtual() int {
	anoDeNascimento := p.DataNascimento.Year()
	anoAtual := time.Now().Year()
	return anoAtual - anoDeNascimento
}

func CalculaIdade(p Pessoa) int {
	anoAtual := time.Now().Year()
	anoNascimento := p.DataNascimento.Year()
	return anoAtual - anoNascimento
}
