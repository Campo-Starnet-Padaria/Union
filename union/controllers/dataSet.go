package controllers

import (
	"fmt"

	"com.github/FelipeAlafy/union/manager"
	"github.com/gotk3/gotk3/gtk"
)

func dataSet(c manager.Client) {
	ZonaRural.SetActive(c.ZonaRural)
	Pago.SetActive(c.Pago)
	Entries[0].SetText(c.Nome)
	Entries[1].SetText(c.Cpf)
	Entries[2].SetText(c.Rg)
	Entries[3].SetText(c.Nascimento)
	Entries[4].SetText(c.Telefone)
	Entries[5].SetText(c.TelefoneFixo)
	Entries[6].SetText(c.ClienteCemig)
	Entries[7].SetText(c.InstalacaoCemig)
	Entries[8].SetText(c.EstadoCivil)
	Entries[9].SetText(fmt.Sprint(c.Numero))	
	Entries[10].SetText(fmt.Sprint(c.Complemento))
	Entries[11].SetText(c.Rua)
	Entries[12].SetText(c.Bairro)
	Entries[13].SetText(c.Cidade)
	Entries[14].SetText(c.Estado)
	Entries[15].SetText(c.Cep)
	Entries[16].SetText(fmt.Sprintf("%v", c.ValorDoKit))	
	Entries[17].SetText(fmt.Sprintf("%v",c.ValorTotal))
	Entries[18].SetText(c.Comissionario)
	Entries[19].SetText(fmt.Sprintf("%v", c.Comissao))
	Entries[20].SetText(c.NotaKit)
	Entries[21].SetText(c.NotaServico)
	Entries[22].SetText(c.EntregaPlacas)
	Entries[23].SetText(c.Instalacao)
	Entries[24].SetText(c.Finalizacao)
	Entries[25].SetText(c.Geracao)
	Entries[26].SetText(c.MDGM)
	Entries[27].SetText(c.PKH)
	Entries[28].SetText(c.Observacao)
	insertRateios(c, Rateios)
}

func insertRateios(c manager.Client, entries []*gtk.Entry) {
	for i, rats := range c.Rateio {
		entries[i].SetText(rats)
	}
}