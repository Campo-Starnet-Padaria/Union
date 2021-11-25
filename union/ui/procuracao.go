package ui

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func procuracaoUi() *gtk.Window {
	w, e := gtk.WindowNew(gtk.WINDOW_TOPLEVEL) 
	if e != nil {
		log.Println("Could not open Procuração Window, because, ", e.Error())
	}
	w.SetDefaultSize(400, 500)
	w.SetTitle("Insira os dados para a procuração")
	hbox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	grid, _ := gtk.GridNew()
	lbl, _ := gtk.LabelNew("Insira a cidade em que a procuração será assinada")
	grid.Attach(lbl, 5, 5, 40, 10)
	
	entry, _ := gtk.EntryNew()
	btn, _ := gtk.ButtonNewWithLabel("Gerar Procuração")
	grid.AttachNextTo(entry, lbl, gtk.POS_RIGHT, 40, 10)
	
	rep, _ := gtk.LabelNew("Representante")
	representante, _ := gtk.EntryNew()
	grid.AttachNextTo(rep, lbl, gtk.POS_BOTTOM, 40, 10)
	grid.AttachNextTo(representante, rep, gtk.POS_RIGHT, 40, 10)
	
	nacRep, _ := gtk.LabelNew("Nacionalidade do representante")
	nacionalidade, _ := gtk.EntryNew()
	grid.AttachNextTo(nacRep, rep, gtk.POS_BOTTOM, 40, 10)
	grid.AttachNextTo(nacionalidade, nacRep, gtk.POS_RIGHT, 40, 10)
	
	ec, _ := gtk.LabelNew("Estado Civil")
	estadoCivil, _ := gtk.EntryNew()
	grid.AttachNextTo(ec, nacRep, gtk.POS_BOTTOM, 40, 10)
	grid.AttachNextTo(estadoCivil, ec, gtk.POS_RIGHT, 40, 10)

	rgrep, _ := gtk.LabelNew("Rg do representante")
	rgRepresentante, _ := gtk.EntryNew()
	grid.AttachNextTo(rgrep, ec, gtk.POS_BOTTOM, 40, 10)
	grid.AttachNextTo(rgRepresentante, rgrep, gtk.POS_RIGHT, 40, 10)
	
	cpfrep, _ := gtk.LabelNew("Cpf do representante")
	cpfRepresentante, _ := gtk.EntryNew()
	grid.AttachNextTo(cpfrep, rgrep, gtk.POS_BOTTOM, 40, 10)
	grid.AttachNextTo(cpfRepresentante, cpfrep, gtk.POS_RIGHT, 40, 10)

	num, _ := gtk.LabelNew("Numero da residência do representante")
	numero, _ := gtk.EntryNew()
	grid.AttachNextTo(num, cpfrep, gtk.POS_BOTTOM, 40, 10)
	grid.AttachNextTo(numero, num, gtk.POS_RIGHT, 40, 10)
	
	complemento, _ := gtk.LabelNew("complemento da residência do representante")
	comp, _ := gtk.EntryNew()
	grid.AttachNextTo(complemento, num, gtk.POS_BOTTOM, 40, 10)
	grid.AttachNextTo(comp, complemento, gtk.POS_RIGHT, 40, 10)
	
	ru, _ := gtk.LabelNew("Rua do representante")
	rua, _ := gtk.EntryNew()
	grid.AttachNextTo(ru, complemento, gtk.POS_BOTTOM, 40, 10)
	grid.AttachNextTo(rua, ru, gtk.POS_RIGHT, 40, 10)
	
	bai, _ := gtk.LabelNew("Bairro do representante")
	bairro, _ := gtk.EntryNew()
	grid.AttachNextTo(bai, ru, gtk.POS_BOTTOM, 40, 10)
	grid.AttachNextTo(bairro, bai, gtk.POS_RIGHT, 40, 10)
	
	cid, _ := gtk.LabelNew("Cidade do representante")
	cidade, _ := gtk.EntryNew()
	grid.AttachNextTo(cid, bai, gtk.POS_BOTTOM, 40, 10)
	grid.AttachNextTo(cidade, cid, gtk.POS_RIGHT, 40, 10)
	
	est, _ := gtk.LabelNew("Estado do representante")
	estado, _ := gtk.EntryNew()
	grid.AttachNextTo(est, cid, gtk.POS_BOTTOM, 40, 10)
	grid.AttachNextTo(estado, est, gtk.POS_RIGHT, 40, 10)
	
	cepp, _ := gtk.LabelNew("Cep do representante")
	cep, _ := gtk.EntryNew()
	grid.AttachNextTo(cepp, est, gtk.POS_BOTTOM, 40, 10)
	grid.AttachNextTo(cep, cepp, gtk.POS_RIGHT, 40, 10)
	
	hbox.PackStart(grid, true, false, 10)
	hbox.PackEnd(btn, true, false, 10)
	
	hbox.ShowAll()
	btn.Connect("clicked", func(){
		city,_ := entry.GetText()
		w.Hide()
		pdfCpf(city, representante, nacionalidade, estadoCivil, rgRepresentante, cpfRepresentante, numero, comp, rua, bairro, cidade, estado, cep)
	})
	w.Add(hbox)
	return w
}