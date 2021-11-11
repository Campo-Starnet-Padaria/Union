package ui

import (
	"log"

	"com.github/FelipeAlafy/union/controllers"
	"github.com/gotk3/gotk3/gtk"
)

var cAdd int = 0
var cEdit int = 0
var entries []*gtk.Entry
var rateios []*gtk.Entry
var zonaRural *gtk.CheckButton
var pago *gtk.CheckButton

func InitInteractions(Entries, Rateios []*gtk.Entry, ZonaRural *gtk.CheckButton, Pago *gtk.CheckButton) {
	entries = Entries
	rateios = Rateios
	zonaRural = ZonaRural
	pago = Pago
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

func Adicionar() {
	if cAdd == 0 {
		clearFields()
		editingMode(true)
		cAdd++
	} else {
		controllers.AddClient(entries, rateios, zonaRural, pago)
		cAdd--
		updateUi()
		editingMode(false)
	}
}

func ArchiveProject(window *gtk.Window) {
	result :=  controllers.Archive()
	if result {
		gtk.MessageDialogNew(window, gtk.DIALOG_USE_HEADER_BAR, gtk.MESSAGE_INFO, gtk.BUTTONS_CLOSE, "O projeto foi arquivado com êxito.", nil)
	} else {
		gtk.MessageDialogNew(window, gtk.DIALOG_USE_HEADER_BAR, gtk.MESSAGE_INFO, gtk.BUTTONS_CLOSE, "O projeto foi desarquivado com êxito.", nil)
	}
}

func openProjectFolder() {
	controllers.OpenInstanceFolder()
}

func EditProject() {
	if cEdit == 0 {
		editingMode(true)
		cEdit++
	} else {
		controllers.Edit()
		editingMode(false)
		updateUi()
		cEdit--
	}
}

func editingMode(editing bool) {
	//Entries
	for _, e := range entries {
		if name, _ := e.GetName(); name != "NameEntry" {
			e.SetSensitive(editing)
		}
	}

	rateios[0].SetSensitive(editing)
	rateios[1].SetSensitive(editing)
	rateios[2].SetSensitive(editing)
	rateios[3].SetSensitive(editing)
	rateios[4].SetSensitive(editing)
	rateios[5].SetSensitive(editing)
	rateios[6].SetSensitive(editing)
	rateios[7].SetSensitive(editing)
	rateios[8].SetSensitive(editing)
	
	zonaRural.SetSensitive(editing)
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
	result := controllers.SearchForClients(text, true)
	clearFields()
	controllers.SetClient(result)
}