package ui

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
)

var Window *gtk.Window

func initRat()  {
	Window.ShowAll()
}


//Call before initRat
func subUiRateio(editing bool) ([]*gtk.Entry) {
	window, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	hb, save := rateioHeaderbar()
	window.SetTitlebar(hb)
	window.SetDefaultSize(600, 250)

	fatal("WindowNew window", err)
	grid, err := gtk.GridNew()
	fatal("GridNew grid", err)
	grid.InsertColumn(1)
	grid.SetColumnSpacing(10)
	grid.SetRowSpacing(10)
	grid.SetColumnHomogeneous(true)
	grid.SetRowHomogeneous(true)

	entries := make([]*gtk.Entry, 10)

	lbl1, _ := gtk.LabelNew("Rateio 1: ")
	entry1, _ := gtk.EntryNew()
	entry1.SetSensitive(editing)

	grid.Attach(lbl1, 5, 10, 10, 1)
	grid.Attach(entry1, 15, 10, 40, 1)

	//For labels
	var sibling *gtk.Label = lbl1
	for i := 2; i < 10; i++ {
		lbl, err := gtk.LabelNew(fmt.Sprint("Rateio ", i, ": "))
		fatal(fmt.Sprint("LabelNew Rateio ", i), err)
		grid.AttachNextTo(lbl, sibling, gtk.POS_BOTTOM, 10, 1)
		sibling = lbl
	}

	//For entries
	var sEntry *gtk.Entry = entry1
	entries[0] = entry1
	for i := 1; i < 9; i++ {
		entry, err := gtk.EntryNew()
		fatal(fmt.Sprint("EntryNew Entry", i), err)
		grid.AttachNextTo(entry, sEntry, gtk.POS_BOTTOM, 40, 1)
		entry.SetSensitive(editing)
		sEntry = entry
		entries[i] = entry
	}
	window.Add(grid)

	save.Connect("clicked", func () {
		window.Hide()
	})
	Window = window
	return entries
}

