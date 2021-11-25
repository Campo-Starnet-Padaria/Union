package ui

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

var fotos *gtk.Button

func labels (grid *gtk.Grid) {
	//thirty one labels
	nome, err := gtk.LabelNew("Nome: ")
	fatal("nome ", err)
	grid.Attach(nome, 1, 0, 3, 1)

	cpf, err := gtk.LabelNew("CPF/CNPJ: ")
	fatal("", err)
	grid.AttachNextTo(cpf, nome, gtk.POS_BOTTOM, 2, 2)

	empresa, err := gtk.LabelNew("Marque para empresas: ")
	fatal("LabelNew empresa", err)
	grid.AttachNextTo(empresa, cpf, gtk.POS_BOTTOM, 2, 2)

	rg, err := gtk.LabelNew("RG: ")
	fatal("", err)
	grid.AttachNextTo(rg, empresa, gtk.POS_BOTTOM, 2, 2)

	nascimento, err := gtk.LabelNew("Nascimento: ")
	fatal("", err)
	grid.AttachNextTo(nascimento, rg, gtk.POS_BOTTOM, 2, 2)

	nacionalidade, err := gtk.LabelNew("Nacionalidade: ")
	fatal("Nacionalidade EntryNew", err)
	grid.AttachNextTo(nacionalidade, nascimento, gtk.POS_BOTTOM, 2, 2)

	telefone, err := gtk.LabelNew("Telefone: ")
	fatal("", err)
	grid.AttachNextTo(telefone, nacionalidade, gtk.POS_BOTTOM, 2, 2)

	telfixo, err := gtk.LabelNew("Telefone fixo: ")
	fatal("", err)
	grid.AttachNextTo(telfixo, telefone, gtk.POS_BOTTOM, 2, 2)

	clienteCemig, err := gtk.LabelNew("Número do cliente da cemig: ")
	fatal("", err)
	grid.AttachNextTo(clienteCemig, telfixo, gtk.POS_BOTTOM, 2, 2)

	clienteInstalacao, err := gtk.LabelNew("Número da instalação da cemig: ")
	fatal("", err)
	grid.AttachNextTo(clienteInstalacao, clienteCemig, gtk.POS_BOTTOM, 2, 2)

	estadoCivil, err := gtk.LabelNew("Estado Civil")
	fatal("", err)
	grid.AttachNextTo(estadoCivil, clienteInstalacao, gtk.POS_BOTTOM, 2, 2)

	numero, err := gtk.LabelNew("Número da residência: ")
	fatal("", err)
	grid.AttachNextTo(numero, estadoCivil, gtk.POS_BOTTOM, 2, 2)

	complemento, err := gtk.LabelNew("Complemento: ")
	fatal("", err)
	grid.AttachNextTo(complemento, numero, gtk.POS_BOTTOM, 2, 2)

	rua, err := gtk.LabelNew("Rua: ")
	fatal("", err)
	grid.AttachNextTo(rua, complemento, gtk.POS_BOTTOM, 2, 2)

	bairro, err := gtk.LabelNew("Bairro: ")
	fatal("", err)
	grid.AttachNextTo(bairro, rua, gtk.POS_BOTTOM, 2, 2)

	cidade, err := gtk.LabelNew("Cidade: ")
	fatal("", err)
	grid.AttachNextTo(cidade, bairro, gtk.POS_BOTTOM, 2, 2)

	estado, err := gtk.LabelNew("Estado (MG): ")
	fatal("", err)
	grid.AttachNextTo(estado, cidade, gtk.POS_BOTTOM, 2, 2)

	cep, err := gtk.LabelNew("CEP: ")
	fatal("", err)
	grid.AttachNextTo(cep, estado, gtk.POS_BOTTOM, 2, 2)

	zonaRural, err := gtk.LabelNew("É zona rural: ")
	fatal("", err)
	grid.AttachNextTo(zonaRural, cep, gtk.POS_BOTTOM, 2, 2)

	kit, err := gtk.LabelNew("Valor do kit: ")
	fatal("", err)
	grid.AttachNextTo(kit, zonaRural, gtk.POS_BOTTOM, 2, 2)

	total, err := gtk.LabelNew("Valor total: ")
	fatal("", err)
	grid.AttachNextTo(total, kit, gtk.POS_BOTTOM, 2, 2)

	pago, err := gtk.LabelNew("Os montadores foram pagos: ")
	fatal("", err)
	grid.AttachNextTo(pago, total, gtk.POS_BOTTOM, 2, 2)

	comissionario, err := gtk.LabelNew("Comissionario: ")
	fatal("", err)
	grid.AttachNextTo(comissionario, pago, gtk.POS_BOTTOM, 2, 2)

	comissao, err := gtk.LabelNew("Comissão: ")
	fatal("", err)
	grid.AttachNextTo(comissao, comissionario, gtk.POS_BOTTOM, 2, 2)

	notaKit, err := gtk.LabelNew("Nota do kit: ")
	fatal("", err)
	grid.AttachNextTo(notaKit, comissao, gtk.POS_BOTTOM, 2, 2)

	notaServico, err := gtk.LabelNew("Nota do serviço: ")
	fatal("", err)
	grid.AttachNextTo(notaServico, notaKit, gtk.POS_BOTTOM, 2, 2)

	entrega, err := gtk.LabelNew("Entrega das placas: ")
	fatal("", err)
	grid.AttachNextTo(entrega, notaServico, gtk.POS_BOTTOM, 2, 2)

	instalacao, err := gtk.LabelNew("Instalação das placas: ")
	fatal("", err)
	grid.AttachNextTo(instalacao, entrega, gtk.POS_BOTTOM, 2, 2)

	finalizacao, err := gtk.LabelNew("Finalização dos serviços: ")
	fatal("", err)
	grid.AttachNextTo(finalizacao, instalacao, gtk.POS_BOTTOM, 2, 2)

	rateio, err := gtk.LabelNew("Rateio: ")
	fatal("", err)
	grid.AttachNextTo(rateio, finalizacao, gtk.POS_BOTTOM, 2, 2)

	geracao, err := gtk.LabelNew("Geração: ")
	fatal("", err)
	grid.AttachNextTo(geracao, rateio, gtk.POS_BOTTOM, 2, 2)

	mdgm, err := gtk.LabelNew("Média de Geração mensal: ")
	fatal("", err)
	grid.AttachNextTo(mdgm, geracao, gtk.POS_BOTTOM, 2, 2)

	ppkw, err := gtk.LabelNew("Preço por Kw/h: ")
	fatal("", err)
	grid.AttachNextTo(ppkw, mdgm, gtk.POS_BOTTOM, 2, 2)

	observacoes, err := gtk.LabelNew("Observações: ")
	fatal("", err)
	grid.AttachNextTo(observacoes, ppkw, gtk.POS_BOTTOM, 2, 2)
}

func fields(grid *gtk.Grid) ([]*gtk.Entry, []*gtk.Calendar, *gtk.CheckButton, *gtk.CheckButton, *gtk.CheckButton) {
	widgets := make([]*gtk.Entry, 26)
	calendar := make([]*gtk.Calendar, 4)
	//thirty one fields
	nome, err := gtk.EntryNew()
	nome.SetName("NameEntry")
	fatal("field Nome", err)
	grid.Attach(nome, 5, 0, 100, 1)
	widgets[0] = nome

	cpf, err := gtk.EntryNew()
	fatal("field cpf", err)
	grid.AttachNextTo(cpf, nome, gtk.POS_BOTTOM, 2, 2)
	widgets[1] = cpf

	empresa, err := gtk.CheckButtonNew()
	fatal("CheckButtonNew empresa", err)
	grid.AttachNextTo(empresa, cpf, gtk.POS_BOTTOM, 2, 2)

	rg, err := gtk.EntryNew()
	fatal("field RG", err)
	grid.AttachNextTo(rg, empresa, gtk.POS_BOTTOM, 2, 2)
	widgets[2] = rg

	nascimento, err := gtk.CalendarNew()
	fatal("field Nascimento", err)
	grid.AttachNextTo(nascimento, rg, gtk.POS_BOTTOM, 2, 2)
	calendar[0] = nascimento

	nacionalidade, err := gtk.EntryNew()
	fatal("Nacionalidade EntryNew", err)
	grid.AttachNextTo(nacionalidade, nascimento, gtk.POS_BOTTOM,  2, 2)
	widgets[3] = nacionalidade 

	telefone, err := gtk.EntryNew()
	fatal("", err)
	grid.AttachNextTo(telefone, nacionalidade, gtk.POS_BOTTOM, 2, 2)
	widgets[4] = telefone

	telfixo, err := gtk.EntryNew()
	fatal("Entry Telefone fixo", err)
	grid.AttachNextTo(telfixo, telefone, gtk.POS_BOTTOM, 2, 2)
	widgets[5] = telfixo

	clienteCemig, err := gtk.EntryNew()
	fatal("Entry cliente cemig", err)
	grid.AttachNextTo(clienteCemig, telfixo, gtk.POS_BOTTOM, 2, 2)
	widgets[6] = clienteCemig

	clienteInstalacao, err := gtk.EntryNew()
	fatal("Entry cliente instalação", err)
	grid.AttachNextTo(clienteInstalacao, clienteCemig, gtk.POS_BOTTOM, 2, 2)
	widgets[7] = clienteInstalacao

	estadoCivil, err := gtk.EntryNew()
	fatal("Entry estado civil", err)
	grid.AttachNextTo(estadoCivil, clienteInstalacao, gtk.POS_BOTTOM, 2, 2)
	widgets[8] = estadoCivil

	numero, err := gtk.EntryNew()
	fatal("Entry número", err)
	grid.AttachNextTo(numero, estadoCivil, gtk.POS_BOTTOM, 2, 2)
	widgets[9] = numero

	complemento, err := gtk.EntryNew()
	fatal("Entry Complemento", err)
	grid.AttachNextTo(complemento, numero, gtk.POS_BOTTOM, 2, 2)
	widgets[10] = complemento

	rua, err := gtk.EntryNew()
	fatal("Entry rua", err)
	grid.AttachNextTo(rua, complemento, gtk.POS_BOTTOM, 2, 2)
	widgets[11] = rua

	bairro, err := gtk.EntryNew()
	fatal("Entry Bairro", err)
	grid.AttachNextTo(bairro, rua, gtk.POS_BOTTOM, 2, 2)
	widgets[12] = bairro

	cidade, err := gtk.EntryNew()
	fatal("Entry cidade", err)
	grid.AttachNextTo(cidade, bairro, gtk.POS_BOTTOM, 2, 2)
	widgets[13] = cidade

	estado, err := gtk.EntryNew()
	fatal("Entry Estado", err)
	grid.AttachNextTo(estado, cidade, gtk.POS_BOTTOM, 2, 2)
	widgets[14] = estado

	cep, err := gtk.EntryNew()
	fatal("Entry Cep", err)
	grid.AttachNextTo(cep, estado, gtk.POS_BOTTOM, 2, 2)
	widgets[15] = cep

	zonaRural, err := gtk.CheckButtonNew()
	fatal("CheckButton Zona Rural", err)
	zonaRural.SetHAlign(gtk.ALIGN_CENTER)
	grid.AttachNextTo(zonaRural, cep, gtk.POS_BOTTOM, 2, 2)

	kit, err := gtk.EntryNew()
	fatal("Entry kit", err)
	grid.AttachNextTo(kit, zonaRural, gtk.POS_BOTTOM, 2, 2)
	widgets[16] = kit

	total, err := gtk.EntryNew()
	fatal("Entry total", err)
	grid.AttachNextTo(total, kit, gtk.POS_BOTTOM, 2, 2)
	widgets[17] = total

	pago, err := gtk.CheckButtonNew()
	fatal("CheckButton Pago ", err)
	pago.SetLabel("Valor: ")
	pago.Widget.SetHAlign(gtk.ALIGN_CENTER)
	grid.AttachNextTo(pago, total, gtk.POS_BOTTOM, 2, 2)
	
	comissionario, err := gtk.EntryNew()
	fatal("Entry comissionario", err)
	grid.AttachNextTo(comissionario, pago, gtk.POS_BOTTOM, 2, 2)
	widgets[18] = comissionario

	comissao, err := gtk.EntryNew()
	fatal("Entry comissão", err)
	grid.AttachNextTo(comissao, comissionario, gtk.POS_BOTTOM, 2, 2)
	comissao.SetSensitive(false)
	comissao.SetName("ComissaoEntry")
	widgets[19] = comissao

	notaKit, err := gtk.EntryNew()
	fatal("Entry nota kit", err)
	grid.AttachNextTo(notaKit, comissao, gtk.POS_BOTTOM, 2, 2)
	widgets[20] = notaKit

	notaServico, err := gtk.EntryNew()
	fatal("Entry nota serviço", err)
	grid.AttachNextTo(notaServico, notaKit, gtk.POS_BOTTOM, 2, 2)
	widgets[21] = notaServico

	
	entrega, err := gtk.CalendarNew()
	fatal("entrega Calendar", err)
	grid.AttachNextTo(entrega, notaServico, gtk.POS_BOTTOM, 2, 2)
	calendar[1] = entrega

	instalacao, err := gtk.CalendarNew()
	fatal("Instalação Calendar", err)
	grid.AttachNextTo(instalacao, entrega, gtk.POS_BOTTOM, 2, 2)
	calendar[2] = instalacao
	
	finalizacao, err := gtk.CalendarNew()
	fatal("Finalização Calendar", err)
	grid.AttachNextTo(finalizacao, instalacao, gtk.POS_BOTTOM, 2, 2)
	calendar[3] = finalizacao
	
	rateio, err := gtk.ButtonNewWithLabel("Clique para definir os rateios")
	fatal("Rateio button", err)
	grid.AttachNextTo(rateio, finalizacao, gtk.POS_BOTTOM, 2, 2)

	geracao, err := gtk.EntryNew()
	fatal("Geração Entry", err)
	grid.AttachNextTo(geracao, rateio, gtk.POS_BOTTOM, 2, 2)
	widgets[22] = geracao

	mdgm, err := gtk.EntryNew()
	fatal("MGDM Entry", err)
	grid.AttachNextTo(mdgm, geracao, gtk.POS_BOTTOM, 2, 2)
	widgets[23] = mdgm

	ppkw, err := gtk.EntryNew()
	fatal("PPKW entry", err)
	grid.AttachNextTo(ppkw, mdgm, gtk.POS_BOTTOM, 2, 2)
	widgets[24] = ppkw

	observacoes, err := gtk.EntryNew()
	fatal("Observações TextView", err)
	grid.AttachNextTo(observacoes, ppkw, gtk.POS_BOTTOM, 2, 20)
	widgets[25] = observacoes

	fotos, err = gtk.ButtonNewWithLabel("Fotos")
	fatal("ButtonNew Fotos", err)
	grid.AttachNextTo(fotos, observacoes, gtk.POS_BOTTOM, 2, 2)

	procuracaoCPF, err := gtk.ButtonNewWithLabel("Gerar procuração")
	fatal("ButtonNew procuracaoCPF", err)
	grid.AttachNextTo(procuracaoCPF, fotos, gtk.POS_BOTTOM, 2, 2)

	for _, entry := range widgets {
		entry.SetSensitive(false)
	}

	for _, cal := range calendar {
		cal.SetSensitive(false)
	}

	zonaRural.SetSensitive(false)
	pago.SetSensitive(false)

	rateio.Connect("clicked", func ()  {	initRat()	})
	fotos.Connect("clicked", func ()  {		photoView()		})
	procuracaoCPF.Connect("clicked", func () { 		
		w := procuracaoUi()
		w.ShowAll()
	})

	return widgets, calendar,  zonaRural, pago, empresa
}

func fatal(comp string, err error) {
	if err != nil {
		log.Fatal("I cannot create ", comp, ", because: ", err.Error())
	}
}