package controllers

import (
	"log"
	"strings"

	"com.github/FelipeAlafy/union/manager"
	"com.github/FelipeAlafy/union/osmanager"
)


func reload() {
	Instances := manager.GetClients(osmanager.GetInstances())
	var nInst []manager.Client
	log.Println("Search for archived project is now: ", FilterArchived)
	if FilterArchived {
		//in this case get all instances when it has Archived property equals false
		for _, ins := range Instances {
			filter := ins.FilterByArchived(FilterArchived)
			if filter {
				nInst = append(nInst, ins)
			}
		}
		instances = nInst
		clientLimit = len(nInst)
	} else {
		//In this case get all instances
		for _, ins := range Instances {
			filter := ins.FilterByArchived(FilterArchived)
			if filter {
				nInst = append(nInst, ins)
			}
		}
		instances = nInst
		clientLimit = len(instances)
	}
}

func SearchForClients(name string) (int) {
	reload()
	//look for a client
	for i, ins := range instances {
		if strings.EqualFold(name, ins.Nome) {
			return i
		}
	}
	return 0
}