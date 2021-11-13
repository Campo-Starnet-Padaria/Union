package controllers

import (
	"fmt"
	"strconv"
	"strings"

	"com.github/FelipeAlafy/union/manager"
	"github.com/gotk3/gotk3/gtk"
)

func dataSet(c manager.Client) {
	ZonaRural.SetActive(c.ZonaRural)
	Pago.SetActive(c.Pago)
	Entries[0].SetText(c.Nome)
	Entries[1].SetText(c.Cpf)
	Entries[2].SetText(c.Rg)
	setDate(Calendar[0], c.Nascimento)
	Entries[3].SetText(c.Nacionalidade)
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
	setDate(Calendar[1], c.EntregaPlacas)
	setDate(Calendar[2], c.Instalacao)	
	setDate(Calendar[3], c.Finalizacao)
	Entries[22].SetText(c.Geracao)
	Entries[23].SetText(c.MDGM)
	Entries[24].SetText(c.PKH)
	Entries[25].SetText(c.Observacao)
	insertRateios(c, Rateios)
}

func setDate(e *gtk.Calendar, date string) {
	if date != "" {
		fdate := strings.Split(date, "/")
		day, _ := strconv.ParseUint(fdate[0], 10, 64)
		month, _ := strconv.ParseUint(fdate[1], 10, 64)
		year, _ := strconv.ParseUint(fdate[2], 10, 64)
		e.SelectDay(uint(day))
		e.SelectMonth(uint(month), uint(year))
	}
} 

func insertRateios(c manager.Client, entries []*gtk.Entry) {
	for i, rats := range c.Rateio {
		entries[i].SetText(rats)
	}
}