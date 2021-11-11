package controllers

import (
	"strconv"

	"com.github/FelipeAlafy/union/manager"
	"github.com/gotk3/gotk3/gtk"
)

var entries []*gtk.Entry

func DataGet(Entries, rateios []*gtk.Entry, zonaRural, pago *gtk.CheckButton) manager.Client {
	var client manager.Client
	entries = Entries

	client.Nome 			= getText(entries[0])
	client.Cpf				= getText(entries[1])
	client.Rg				= getText(entries[2])
	client.Nascimento		= getText(entries[3])
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
	client.Montagem			= toFloat(getText(entries[18]))
	client.Pago				= Pago.GetActive()
	client.Comissionario	= getText(entries[19])
	client.Comissao			= toFloat(getText(entries[20]))
	client.NotaKit			= getText(entries[20])
	client.NotaServico		= getText(entries[21])
	client.EntregaPlacas	= getText(entries[22])
	client.Instalacao		= getText(entries[23])
	client.Finalizacao		= getText(entries[24])
	client.Rateio			= [9]string{}
	client.Geracao			= getText(entries[25])
	client.MDGM				= getText(entries[26])
	client.PKH				= getText(entries[27])
	client.Observacao		= getText(entries[28])
	client.Archived			= false
	//client.Fotos --- Will be implement in soon

	for i := 0; i < len(client.Rateio); i++ {
		client.Rateio[i] = getRat(rateios[i])
	}


	return client
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