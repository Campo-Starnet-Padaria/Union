package ui

import (
	"fmt"
	"log"

	"com.github/FelipeAlafy/union/controllers"
	"github.com/gotk3/gotk3/gtk"
)

var cAdd int = 0
var cEdit int = 0
var entries []*gtk.Entry
var rateios []*gtk.Entry
var calendar []*gtk.Calendar
var zonaRural *gtk.CheckButton
var pago *gtk.CheckButton
var empresa *gtk.CheckButton
var Archived *gtk.Switch

func InitInteractions(Entries, Rateios []*gtk.Entry, Calendar []*gtk.Calendar, ZonaRural *gtk.CheckButton, Pago *gtk.CheckButton, Empresa *gtk.CheckButton) {
	entries = Entries
	rateios = Rateios
	calendar = Calendar
	zonaRural = ZonaRural
	pago = Pago
	empresa = Empresa

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
		NewButton.SetLabel("Adicionar")
		fotos.SetSensitive(false)
		ProcuracaoButton.SetSensitive(false)
		fotos.SetTooltipText("Depois que você adicionar esse projeto.\n Vá passando até chegar nele, e assim poderá adicionar.")
		entries[0].SetSensitive(true)
		entries[0].Connect("activate", func () {
			editingMode(true)
			cAdd++
		})
	} else {
		controllers.AddClient()
		cAdd--
		updateUi()
		editingMode(false)
		ProcuracaoButton.SetSensitive(true)
		fotos.SetSensitive(true)
		fotos.SetTooltipText("")
		NewButton.SetLabel("Novo")
	}
}

func ArchiveProject(window *gtk.Window) {
	result :=  controllers.Archive()
	if result {
		dialog := gtk.MessageDialogNew(window, gtk.DIALOG_USE_HEADER_BAR, gtk.MESSAGE_INFO, gtk.BUTTONS_CLOSE, "O projeto foi arquivado com êxito.", nil)
		dialog.Run()
	} else {
		dialog := gtk.MessageDialogNew(window, gtk.DIALOG_USE_HEADER_BAR, gtk.MESSAGE_INFO, gtk.BUTTONS_CLOSE, "O projeto foi desarquivado com êxito.", nil)
		dialog.Run()
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
			if name == "ComissaoEntry" {
				e.SetSensitive(false)
			}
		} else if (name == "NameEntry" && !editing) {
			e.SetSensitive(editing)
		}
	}

	for _, c := range calendar {
		c.SetSensitive(editing)
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

//Configurações
func configs() {
	popover, _ := gtk.PopoverNew(ConfigButton)
	archived, _ := gtk.SwitchNew()
	lbl, _ := gtk.LabelNew("Habilitar arquivados: ")

	grid, _ := gtk.GridNew()
	grid.Attach(lbl, 1, 1, 1, 1)
	grid.AttachNextTo(archived, lbl, gtk.POS_RIGHT, 10, 1)
	popover.Add(grid)
	
	grid.ShowAll()
	popover.ShowNow()
	controllers.FilterArchived = archived.GetActive()
}

//PDF
func pdfCpf(city string, representante, nacionalidade, estadoCivil, rgRepresentante, cpfRepresentante, numero, comp, rua, bairro, cidade, estado, cep *gtk.Entry) {
	client := controllers.GetActualClientInstance()
	

	if !client.Empresa {
		client.GenProcuracaoCPF(city)
	} else {
		client.GenProcuracaoCNPJ(city, getText(representante), getText(nacionalidade), getText(estadoCivil), getText(rgRepresentante), getText(cpfRepresentante), getText(numero), getText(comp), getText(rua),
		 getText(bairro), getText(cidade), getText(estado), getText(cep))
	}
}

func getText(e *gtk.Entry) string {
	t, _ := e.GetText()
	return  t
}