package ui

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func OnActivate(app *gtk.Application) {
	appWindow, err := gtk.ApplicationWindowNew(app)
	if err != nil {
		log.Fatal("I cannot create appWindow: ", err)
	} 
	appWindow.SetDefaultSize(720, 680)
	appWindow.SetPosition(gtk.WIN_POS_CENTER)
	headerbar, _ := headerbar()
	appWindow.SetTitlebar(headerbar)
	applyStyle(true)
	ui(appWindow)
	appWindow.ShowAll()
}

func ui(app *gtk.ApplicationWindow) {
	mainbox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	if err != nil {
		log.Fatal("I cannot create MainBox, because: ", err.Error())
	}
	mainbox.SetName("mainbox")
	app.Add(mainbox)

	//Searchbar
	searchEntry, err := gtk.SearchEntryNew()
	if err != nil {
		log.Fatal("I cannot create SearchEntry, beacause: ", err.Error())
	}
	searchEntry.Widget.SetHAlign(gtk.ALIGN_CENTER)
	mainbox.PackStart(searchEntry, false, false, 5)

	//Label
	label, err := gtk.LabelNew("Cliente")
	if err != nil {
		log.Fatal("I cannot create label, beacause: ", err.Error())
	}
	mainbox.PackStart(label, false, false, 5)

	//Hbox
	hbox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	if err != nil {
		log.Fatal("I cannot create Hbox, beacause: ", err.Error())
	}
	mainbox.PackEnd(hbox, true, true, 5)

	
	//Button backward
	backward, err := gtk.ButtonNewFromIconName("forward-rtl-symbolic", gtk.ICON_SIZE_BUTTON)
	if err != nil {
		log.Fatal("I cannot create Bakckward button, beacause: ", err.Error())
	}
	hbox.PackStart(backward, false, true, 2)

	//grid
	grid, err := gtk.GridNew()
	if err != nil {
		log.Fatal("I cannot create grid button, beacause: ", err.Error())
	}
	grid.SetRowSpacing(5)
	grid.SetHAlign(gtk.ALIGN_CENTER)
	grid.SizeAllocate(app.GetAllocation())
	labels(grid)
	fields(grid)

	//ScrolledWindow
	scroll, err := gtk.ScrolledWindowNew(nil, nil)
	if err != nil {
		log.Fatal("I cannot create ScrolledWindow button, beacause: ", err.Error())
	}
	hbox.PackStart(scroll, true, true, 5)
	scroll.Add(grid)

	//Button forward
	forward, err := gtk.ButtonNewFromIconName("forward-symbolic", gtk.ICON_SIZE_BUTTON)
	if err != nil {
		log.Fatal("I cannot create Forward button, beacause: ", err.Error())
	}
	hbox.PackEnd(forward, false, true, 2)
}

func applyStyle(dark bool) {
	css, err := gtk.CssProviderNew()
	if err != nil {
		log.Println("0 Failed to apply css style: ", err.Error())
		return
	}
	if dark {
		err = css.LoadFromPath("ui/dark.css")
		if err != nil {
			log.Println("1 Failed to apply css style: ", err.Error())
		}
	} else {
		err = css.LoadFromData("ui/light.css")
		if err != nil {
			log.Println("2 Failed to apply css style: ", err.Error())
		}
	}
}