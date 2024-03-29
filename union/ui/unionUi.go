package ui

import (
	"log"

	"com.github/FelipeAlafy/union/controllers"
	"github.com/gotk3/gotk3/gtk"
)

var editing 			bool = false
var NewButton 			*gtk.Button
var ConfigButton 		*gtk.Button
var EditButton 			*gtk.Button
var FolderButton 		*gtk.Button
var ArchiveButton 		*gtk.Button
var RateioButton 		*gtk.Button
var ProcuracaoButton	*gtk.Button

func OnActivate(app *gtk.Application) {
	appWindow, err := gtk.ApplicationWindowNew(app)
	Err("AppWindow", err) 
	appWindow.SetDefaultSize(720, 680)
	appWindow.SetPosition(gtk.WIN_POS_CENTER)
	appWindow.SetIconFromFile("union_logo.jpg")
	
	headerbar, headerbarButtons := headerbar()
	NewButton = headerbarButtons[0]
	EditButton = headerbarButtons[1]
	ConfigButton = headerbarButtons[2]
	ArchiveButton = headerbarButtons[3]
	FolderButton = headerbarButtons[4]

	appWindow.SetTitlebar(headerbar)
	Rateios := subUiRateio(editing)
	rateios = Rateios

	ui(appWindow, headerbarButtons)
	appWindow.ShowAll()
}

func ui(app *gtk.ApplicationWindow, hbb []*gtk.Button) {
	mainbox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	Err("MainBox", err)
	mainbox.SetName("mainbox")
	app.Add(mainbox)

	//Searchbar
	searchEntry, err := gtk.SearchEntryNew()
	Err("SearchEntry", err)
	searchEntry.Widget.SetHAlign(gtk.ALIGN_CENTER)
	mainbox.PackStart(searchEntry, false, false, 5)

	//Label
	label, err := gtk.LabelNew("Cliente")
	Err("Label Cliente", err)
	mainbox.PackStart(label, false, false, 5)

	//Hbox
	hbox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	Err("Hbox", err)
	mainbox.PackEnd(hbox, true, true, 5)
	
	//Button backward
	backward, err := gtk.ButtonNewFromIconName("forward-rtl-symbolic", gtk.ICON_SIZE_BUTTON)
	Err("Backward button", err)
	hbox.PackStart(backward, false, true, 2)

	//grid
	grid, err := gtk.GridNew()
	Err("Grid", err)
	grid.SetRowSpacing(5)
	grid.SetHAlign(gtk.ALIGN_CENTER)
	grid.SizeAllocate(app.GetAllocation())
	labels(grid)

	//load Fields to ui
	entries, calendar, zonaRural, pago, empresa, Procuracao, RateioB := fields(grid)
	ProcuracaoButton = Procuracao
	RateioButton = RateioB

	//ScrolledWindow
	scroll, err := gtk.ScrolledWindowNew(nil, nil)
	Err("Scroll", err)
	hbox.PackStart(scroll, true, true, 5)
	scroll.Add(grid)

	//Button forward
	forward, err := gtk.ButtonNewFromIconName("forward-symbolic", gtk.ICON_SIZE_BUTTON)
	Err("Forward button", err)
	hbox.PackEnd(forward, false, true, 2)

	//Functions
	entries[1].Connect("insert-text", func ()  {
		pos := entries[1].GetPosition()
		buff, _ := entries[1].GetBuffer()
		log.Println("Entry CPF está na posição", pos)
		defer entries[1].SetPosition(-1)
		if pos == 2 || pos == 6 {
			i := buff.InsertText(uint(pos + 1), ".")
			log.Println("Entry CPF está na posição", i)
		} else if pos == 10 {
			i := buff.InsertText(uint(pos +1), "-")
			log.Println("Entry CPF está na posição", i)
		}
		text, _ := entries[1].GetText()
		entries[1].SetText("")
		entries[1].SetText(text)
	})
	
	controllers.Init(entries, calendar, rateios, zonaRural, pago, empresa, Archived.GetActive())
	InitInteractions(entries, rateios, calendar,  zonaRural, pago, empresa)
	searchEntry.Connect("activate", func ()  {
		searchFor(searchEntry)
	})
  
	forward.Connect("clicked", func()  { 	nextClient()												})
	backward.Connect("clicked", func() { 	previousClient()											})
	hbb[0].Connect("clicked", func ()  { 	Adicionar(entries, rateios, zonaRural, pago, empresa)		})
	hbb[1].Connect("clicked", func()   { 	EditProject()												})
	hbb[2].Connect("clicked", func()   { 	configs()													})
	hbb[3].Connect("clicked", func()   { 	ArchiveProject(Window)										})
	hbb[4].Connect("clicked", func ()  {	openProjectFolder()											})
}

func Err(comp string, err error) {
	if err != nil {
		log.Fatal("I cannot create ", comp, ", because: ", err.Error())
	}
}