package controllers

import (
	"fmt"
	"log"
	"runtime"

	"com.github/FelipeAlafy/union/manager"
	"com.github/FelipeAlafy/union/osmanager"
	"github.com/gotk3/gotk3/gtk"
	"github.com/skratchdot/open-golang/open"
)

var clientLimit int
var currentClient int = -1
var instances []manager.Client
var Entries []*gtk.Entry
var Calendar []*gtk.Calendar
var	ZonaRural *gtk.CheckButton
var Pago *gtk.CheckButton
var Rateios []*gtk.Entry
var FilterArchived bool = false
var Empresa *gtk.CheckButton
var so string

func Init(entries []*gtk.Entry, calendar []*gtk.Calendar, rateios []*gtk.Entry, zonaRural *gtk.CheckButton, pago *gtk.CheckButton, empresa *gtk.CheckButton, Archived bool) {
	reload()
	FilterArchived = Archived
	Entries = entries
	Calendar = calendar
	ZonaRural = zonaRural 
	Pago = pago
	Rateios = rateios
	Empresa = empresa

	os := runtime.GOOS
    switch os {
    case "windows":
        fmt.Println("Windows")
		so = "\\"
    case "darwin":
        fmt.Println("MAC operating system")
		so = "/"
    case "linux":
        fmt.Println("Linux")
		so = "/"
    default:
        fmt.Printf("%s.\n", os)
		so = "/"
    }

	if clientLimit > 0 {
		NextClient()
	}
}

func forward() manager.Client {
	reload()
	if clientLimit > 0 {
		currentClient++
		if currentClient >= clientLimit {
			log.Println("I cannot read the next client")
			currentClient--
			return instances[currentClient]
		}
		return instances[currentClient]
	}
	return manager.Client{}
}

func backward() manager.Client {
	reload()
	if clientLimit > 0 {
		currentClient--
		if currentClient <= -1 {
			log.Println("I cannot read the back client")
			currentClient++
			return instances[currentClient]
		}
		return instances[currentClient]
	}
	return manager.Client{}
}

//Insert data of the next client on view
func NextClient() {
	client := forward()
	if client.Nome != "" {
		dataSet(client)
	}
}

//Insert data of the previous client on view
func PreviousClient(){
	client := backward()
	if client.Nome != "" {
		dataSet(client)
	}
}

func ActualClient() {
	reload()
	if clientLimit > 1 {
		if currentClient == -1 {
			currentClient++
		}
		dataSet(instances[currentClient])
	}
}

func GetActualClientInstance() manager.Client {
	reload()
	if currentClient == -1 {
		currentClient++
	}
	return instances[currentClient]
}

func SetClient(i int) {
	currentClient := i
	dataSet(instances[currentClient])
}

func AddClient(entries, Rateios []*gtk.Entry, ZonaRural, Pago, Empresa *gtk.CheckButton) {
	client := DataGet(entries, Rateios, ZonaRural, Pago, Empresa)
	osmanager.CreateInstanceFolder(client)
	instance := osmanager.NewInstance(client)
	json, err := client.ClientToJson()
	if err != nil {
		log.Fatal("I cannot add client. ", err.Error())
	}
	osmanager.InstanceData(json, instance)
	ActualClient()
}

func Archive() bool {
	defer ActualClient()
	name, _ := Entries[0].GetText()
	if name != "" {
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
	} else {
		return false
	}
}

func Edit() {
	client := DataGet(Entries, Rateios, ZonaRural, Pago, Empresa)
	osmanager.ReWriteInstanceData(client)
}

func OpenInstanceFolder() {
	err := open.Start(fmt.Sprint("union", so, "instances", so, instances[currentClient].Nome, so))
	if err != nil {
		log.Println("I cannot open the project folder. Because: ", err.Error())
	}
}