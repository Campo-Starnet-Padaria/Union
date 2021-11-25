package controllers

import (
	"fmt"
	"strconv"

	"com.github/FelipeAlafy/union/manager"
	"github.com/gotk3/gotk3/gtk"
)

var entries []*gtk.Entry

func DataGet(Entries, rateios []*gtk.Entry, zonaRural, pago *gtk.CheckButton, empresa *gtk.CheckButton) manager.Client {
	var client manager.Client
	entries = Entries

	client.Nome 			= getText(entries[0])
	client.Cpf				= getText(entries[1])
	client.Empresa			= empresa.GetActive()
	client.Rg				= getText(entries[2])
	client.Nascimento		= getDate(*Calendar[0])
	client.Nacionalidade	= getText(entries[3])
	client.Telefone			= getText(entries[4])
	client.TelefoneFixo		= getText(entries[5])
	client.ClienteCemig		= getText(entries[6])
	client.InstalacaoCemig	= getText(entries[7])
	client.EstadoCivil		= getText(entries[8])
	client.Numero			= toInt(getText(entries[9]))
	client.Complemento		= toInt(getText(entries[10]))
	client.Rua				= getText(entries[11])
	client.Bairro			= getText(entries[12])
	client.Cidade			= getText(entries[13])
	client.Estado			= getText(entries[14])
	client.Cep				= getText(entries[15])
	client.ZonaRural		= ZonaRural.GetActive()
	client.ValorDoKit		= toFloat(getText(entries[16]))
	client.ValorTotal		= toFloat(getText(entries[17]))
	client.Montagem			= ValorMontadores(getText(entries[16]))
	client.Pago				= Pago.GetActive()
	client.Comissionario	= getText(entries[18])
	client.Comissao			= ValorComissao(getText(entries[16]))
	client.NotaKit			= getText(entries[20])
	client.NotaServico		= getText(entries[21])
	client.EntregaPlacas	= getDate(*Calendar[1])
	client.Instalacao		= getDate(*Calendar[2])
	client.Finalizacao		= getDate(*Calendar[3])
	client.Rateio			= [9]string{}
	client.Geracao			= getText(entries[22])
	client.MDGM				= getText(entries[23])
	client.PKH				= getText(entries[24])
	client.Observacao		= getText(entries[25])
	client.Archived			= false

	for i := 0; i < len(client.Rateio); i++ {
		client.Rateio[i] = getRat(rateios[i])
	}

	return client
}

func getDate(e gtk.Calendar) string {
	year, month, day := e.GetDate()
	return fmt.Sprintf("%v/%v/%v", day, month, year)
}

func getRat(entry *gtk.Entry) string {
	txt, _ := entry.GetText()
	return txt
}

func getText(e *gtk.Entry) string {
	str, _ := e.GetText()
	return str
}

func toInt(str string) int {
	v, _ := strconv.Atoi(str)
	return v
}

func toFloat(str string) float64 {
	v, _ := strconv.ParseFloat(str, 64)
	return v
}