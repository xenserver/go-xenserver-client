package main

// Execute "go run list-all-vms.go" to run the code

import (
	"log"

	"github.com/xenserver/go-xenserver-client"
)
import "fmt"

func main() {
	//log.SetLevel(log.DebugLevel)

	// Create a new client
	client := client.NewXenAPIClient("127.0.0.1", "username", "password")

	// Login, this will create a session that is used by all other methods
	err := client.Login()
	if err != nil {
		log.Panic(err)
	}

	// Fetch a list of all VMs, unfortunaly this will return templates and the domain controller as well
	vms, err := client.GetVMs()
	if err != nil {
		log.Panic(err)
	}

	// Loop thru all VMs
	for _, vm := range vms {
		// Fetch information about this vm. This method will also return a list with all available records
		vm.GetRecord()

		// Ignore all that is not a VM (templates and domain controller)
		if !vm.IsVM {
			continue
		}

		// Print out information about the VM
		fmt.Printf("%s (%s) - %s\n", vm.Name, vm.PowerState, vm.Description)
	}

}
