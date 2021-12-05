package ui

import (
	"fmt"
	"log"
	"time"

	"com.github/FelipeAlafy/union/controllers"
	"github.com/gotk3/gotk3/gtk"
)

var cAdd 		int = 0
var cEdit 		int = 0
var entries 	[]*gtk.Entry
var rateios 	[]*gtk.Entry
var calendar 	[]*gtk.Calendar
var zonaRural 	*gtk.CheckButton
var empresa		*gtk.CheckButton
var pago 		*gtk.CheckButton
var Archived 	*gtk.Switch
var filter 		bool = false
var calendars 	[]*gtk.Calendar

func InitInteractions(Entries, Rateios []*gtk.Entry, Calendar []*gtk.Calendar, ZonaRural *gtk.CheckButton, Pago *gtk.CheckButton, Empresa *gtk.CheckButton) {
	entries = Entries
	rateios = Rateios
	calendar = Calendar
	zonaRural = ZonaRural
	pago = Pago
	calendars = calendar
	empresa = Empresa
	controllers.FilterArchived = false

	//Entry: Valor do Kit
	entries[16].Connect("activate", func ()  {
		v, _ := entries[16].GetText()
		m := controllers.CashFormat(controllers.ValorMontadores(v))
		pago.SetLabel(fmt.Sprint("Valor: R$ ", m))
	})

	//Entry: Valor 
	entries[18].Connect("activate", func ()  {
		v, _ := entries[16].GetText()
		m := controllers.CashFormat(controllers.ValorComissao(v))
		entries[19].SetText(fmt.Sprint("Valor: R$ ", m))
	})
	empresa.SetSensitive(false)

	if entries[0].GetTextLength() > 0 {
		hbEditMode(true)
		RateioButton.SetSensitive(true)
	} else {
		hbEditMode(false)
		RateioButton.SetSensitive(false)
	}
}

func hbEditMode(e bool) {
	EditButton.SetSensitive(e)
	FolderButton.SetSensitive(e)
	ArchiveButton.SetSensitive(e)
	ProcuracaoButton.SetSensitive(e)
}

func nextClient() {
	controllers.NextClient()
}

func previousClient() {
	controllers.PreviousClient()
}

func clearFields() {
	for _, v := range entries {
		v.SetText("")
	}

	for _, v := range rateios {
		v.SetText("")
	} 

	zonaRural.SetActive(false)
	pago.SetActive(false)
}

func Adicionar(entries, Rateios []*gtk.Entry, ZonaRural, Pago, Empresa *gtk.CheckButton) {
	if cAdd == 0 {
		clearFields()
		NewButton.SetLabel("Adicionar")
		ProcuracaoButton.SetSensitive(false)
		entries[0].SetSensitive(true)
		hbEditMode(false)
		entries[0].Connect("activate", func () {
			editingMode(true, false, true)
			entries[0].SetSensitive(false)
			cAdd++
		})
	} else {
		controllers.AddClient(entries, Rateios, ZonaRural, Pago, Empresa)
		cAdd--
		updateUi()
		editingMode(false, true, false)
		hbEditMode(true)
		ProcuracaoButton.SetSensitive(true)
		NewButton.SetLabel("Novo")
	}
}

func ArchiveProject(window *gtk.Window) {
	result :=  controllers.Archive()
	if result {
		log.Println(">>> Project was archived.")
	} else {
		log.Println(">>> Project was unarchived.")
	}
}

func openProjectFolder() {
	controllers.OpenInstanceFolder()
}

func EditProject() {
	if cEdit == 0 {
		editingMode(true, true, true)
		cEdit++
	} else {
		controllers.Edit()
		editingMode(false, true, false)
		updateUi()
		cEdit--
	}
}

func rateiosEditMode(e bool) {
	rateios[0].SetSensitive(e)
	rateios[1].SetSensitive(e)
	rateios[2].SetSensitive(e)
	rateios[3].SetSensitive(e)
	rateios[4].SetSensitive(e)
	rateios[5].SetSensitive(e)
	rateios[6].SetSensitive(e)
	rateios[7].SetSensitive(e)
	rateios[8].SetSensitive(e)
}

func editingMode(editing bool, hb bool, r bool) {
	//Entries
	for _, e := range entries {
		if name, _ := e.GetName(); name != "NameEntry" {
			e.SetSensitive(editing)
			if name == "ComissaoEntry" {
				e.SetSensitive(false)
			}
		} else if (name == "NameEntry" && !editing) {
			e.SetSensitive(editing)
		}
	}

	if entries[0].GetTextLength() > 0 {
		hbEditMode(hb)
		RateioButton.SetSensitive(true)
	} else {
		hbEditMode(hb)
		RateioButton.SetSensitive(false)
	}

	for _, c := range calendars {
		t := time.Now()
		c.SelectDay(uint(t.Day()))
		c.SelectMonth(uint(t.Month()), uint(t.Year()))
	}

	for _, c := range calendar {
		c.SetSensitive(editing)
	}

	rateiosEditMode(r)
	
	zonaRural.SetSensitive(editing)
	empresa.SetSensitive(editing)
	pago.SetSensitive(editing)
}

func updateUi() {
	clearFields()
	controllers.ActualClient()
}

//Search interaction
func searchFor(Entry *gtk.SearchEntry) {
	text, _ := Entry.GetText()
	log.Println("Searching for ", text)
	result := controllers.SearchForClients(text)
	clearFields()
	controllers.SetClient(result)
}

//Configurações
func configs() {
	popover, _ := gtk.PopoverNew(ConfigButton)
	archived, _ := gtk.SwitchNew()
	lbl, _ := gtk.LabelNew("Habilitar arquivados: ")

	grid, _ := gtk.GridNew()
	grid.Attach(lbl, 1, 1, 1, 1)
	grid.AttachNextTo(archived, lbl, gtk.POS_RIGHT, 10, 1)
	popover.Add(grid)

	archived.SetActive(filter)
	
	grid.ShowAll()
	popover.ShowNow()
	filter = archived.GetActive()
	controllers.FilterArchived = filter
	archived.Connect("state-set", func ()  {
		filter = !filter
		controllers.FilterArchived = filter
	})
}

//PDF
func pdfCpf(city string) {
	client := controllers.GetActualClientInstance()
	log.Println("Gen CPF Procuração for ", client.Nome)
	client.GenProcuracaoCPF(city)
}

func pdfCNPJ(city string, representante, nacionalidade, estadoCivil, rgRepresentante, cpfRepresentante, numero, comp, rua, bairro, cidade, estado, cep *gtk.Entry) {
	client := controllers.GetActualClientInstance()
	log.Println("Gen CNPJ procuração for ", client.Nome)
	client.GenProcuracaoCNPJ(city, getText(representante), getText(nacionalidade), getText(estadoCivil), getText(rgRepresentante), getText(cpfRepresentante), getText(numero), getText(comp), getText(rua),
		 getText(bairro), getText(cidade), getText(estado), getText(cep))
}

func getText(e *gtk.Entry) string {
	t, _ := e.GetText()
	return  t
}