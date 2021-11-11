package controllers

import (
	"com.github/FelipeAlafy/union/manager"
	"com.github/FelipeAlafy/union/osmanager"
)

func reload() {
	Instances := manager.GetClients(osmanager.GetInstances())
	var nInst []manager.Client
	if !filterArchived {
		//in this case get all instances when it has Archived property equals false
		for _, ins := range Instances {
			filter := ins.FilterByArchived(filterArchived)
			if filter {
				nInst = append(nInst, ins)
			}
		}
		instances = nInst
		clientLimit = len(nInst)
	} else {
		//In this case get all instances
		instances = Instances
		clientLimit = len(instances)
	}
}