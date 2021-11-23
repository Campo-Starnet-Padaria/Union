package ui

import (
	"log"

	"com.github/FelipeAlafy/union/controllers"
	"github.com/gotk3/gotk3/gtk"
)

var editing bool = false
var NewButton *gtk.Button
var ConfigButton *gtk.Button

func OnActivate(app *gtk.Application) {
	appWindow, err := gtk.ApplicationWindowNew(app)
	Err("AppWindow", err) 
	appWindow.SetDefaultSize(720, 680)
	appWindow.SetPosition(gtk.WIN_POS_CENTER)
	headerbar, headerbarButtons := headerbar()
	NewButton = headerbarButtons[0]
	ConfigButton = headerbarButtons[3]
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
	backward.SetOpacity(0.10)

	//grid
	grid, err := gtk.GridNew()
	Err("Grid", err)
	grid.SetRowSpacing(5)
	grid.SetHAlign(gtk.ALIGN_CENTER)
	grid.SizeAllocate(app.GetAllocation())
	labels(grid)
	entries, calendar, zonaRural, pago := fields(grid)

	//ScrolledWindow
	scroll, err := gtk.ScrolledWindowNew(nil, nil)
	Err("Scroll", err)
	hbox.PackStart(scroll, true, true, 5)
	scroll.Add(grid)

	//Button forward
	forward, err := gtk.ButtonNewFromIconName("forward-symbolic", gtk.ICON_SIZE_BUTTON)
	Err("Forward button", err)
	hbox.PackEnd(forward, false, true, 2)
	forward.SetOpacity(0.10)


	//Functions
	controllers.Init(entries, calendar, rateios, zonaRural, pago)
	InitInteractions(entries, rateios, calendar,  zonaRural, pago)
	searchEntry.Connect("activate", func ()  {
		searchFor(searchEntry)
	})
  
	forward.Connect("clicked", func() { 	nextClient()	})
	backward.Connect("clicked", func() { 	previousClient()	})
	hbb[0].Connect("clicked", func ()  { 	Adicionar()		})
	hbb[1].Connect("clicked", func() { 		EditProject()	})
	hbb[2].Connect("clicked", func() { 		configs()		})
	hbb[3].Connect("clicked", func() { 		ArchiveProject(Window)		})
	hbb[4].Connect("clicked", func ()  {	openProjectFolder()		})

}

func Err(comp string, err error) {
	if err != nil {
		log.Fatal("I cannot create ", comp, ", because: ", err.Error())
	}
	
}