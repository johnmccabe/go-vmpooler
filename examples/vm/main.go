package main

import (
	"log"

	"github.com/johnmccabe/go-vmpooler/vm"
)

const vmpoolerEndpoint = "https://vmpooler.mydomain.com/api/v1"
const myToken = "randomlygeneratedtoken" // see Generate() in github.com/johnmccabe/go-vmpooler/token
const myTemplate = "centos-5-x86_64"

func main() {
	var virtualmachine *vm.VM
	var err error
	var virtualmachines []vm.VM

	log.Print("Create Client ===")
	c := vm.NewClient(vmpoolerEndpoint, myToken)

	log.Print("Query Templates ===")
	templates, err := c.ListTemplates()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("%#v", templates)

	log.Print("Create VM ===")
	virtualmachine, err = c.Create(myTemplate)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("%#v", *virtualmachine)

	log.Print("SetLifetime VM ===")
	virtualmachine, err = c.SetLifetime(virtualmachine.Hostname, 20)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("%#v", *virtualmachine)

	log.Print("SetTags VM ===")
	tags := map[string]string{
		"john":    "mccabe",
		"belfast": "antrim",
	}
	virtualmachine, err = c.SetTags(virtualmachine.Hostname, tags)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("%#v", *virtualmachine)

	log.Print("Query VM ===")
	virtualmachine, err = c.Get(virtualmachine.Hostname)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("%#v", *virtualmachine)

	log.Print("Delete VM ===")
	if err = c.Delete(virtualmachine.Hostname); err != nil {
		log.Fatal(err.Error())
	} else {
		log.Printf("deleted: %s", virtualmachine.Hostname)
	}

	log.Print("List VMs ===")
	virtualmachines, err = c.GetAll()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("%#v", virtualmachines)

	log.Print("Delete All ===")
	for _, vm := range virtualmachines {
		if err = c.Delete(vm.Hostname); err != nil {
			log.Fatal(err.Error())
		} else {
			log.Printf("deleted: %s", vm.Hostname)
		}
	}
}
