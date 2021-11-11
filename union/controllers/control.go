package controllers

import (
	"fmt"
	"log"

	"com.github/FelipeAlafy/union/manager"
	"com.github/FelipeAlafy/union/osmanager"
	"github.com/gotk3/gotk3/gtk"
	"github.com/skratchdot/open-golang/open"
)

var clientLimit int
var currentClient int = -1
var instances []manager.Client
var Entries []*gtk.Entry
var	ZonaRural *gtk.CheckButton
var Pago *gtk.CheckButton
var Rateios []*gtk.Entry
var filterArchived bool

func Init(entries []*gtk.Entry, rateios []*gtk.Entry, zonaRural *gtk.CheckButton, pago *gtk.CheckButton) {
	filterArchived = false
	Entries = entries
	ZonaRural = zonaRural 
	Pago = pago
	Rateios = rateios
	
	NextClient()
}


func forward() manager.Client {
	reload()
	currentClient++
	if currentClient >= clientLimit {
		log.Println("I cannot read the next client")
		currentClient--
		return instances[currentClient]
	}
	return instances[currentClient]
}

func backward() manager.Client {
	reload()
	currentClient--
	if currentClient <= -1 {
		log.Println("I cannot read the back client")
		currentClient++
		return instances[currentClient]
	}
	return instances[currentClient]
}

//Insert data of the next client on view
func NextClient() {
	dataSet(forward())
}

//Insert data of the previous client on view
func PreviousClient(){
	dataSet(backward())
}

func ActualClient() {
	reload()
	dataSet(instances[currentClient])
}

func SetClient(i int) {
	currentClient := i
	dataSet(instances[currentClient])
}

func AddClient(entries, rateios []*gtk.Entry, zonaRural, pago *gtk.CheckButton) {
	client := DataGet(entries, rateios, zonaRural, pago)
	osmanager.CreateInstanceFolder(client)
	instance := osmanager.NewInstance(client)
	json, err := client.ClientToJson()
	if err != nil {
		log.Fatal("I cannot add client")
	}
	osmanager.InstanceData(json, instance)
	ActualClient()
}

func Archive() bool {
	defer ActualClient()
	name, _ := Entries[0].GetText()
	c := manager.GetClient(osmanager.GetSpecificInstances(name))
	if !c.Archived {
		c.Archived = true
		osmanager.ReWriteInstanceData(c)
		return true
	} else {
		c.Archived = false
		osmanager.ReWriteInstanceData(c)
		return false
	}
}

func Edit() {
	client := DataGet(Entries, Rateios, ZonaRural, Pago)
	osmanager.ReWriteInstanceData(client)
}

func OpenInstanceFolder() {
	open.Run(fmt.Sprint("union/instances/", instances[currentClient].Nome, "/"))
}