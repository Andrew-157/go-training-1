// Dummy server that does nothing and uses stuff from example.com/utils
package server

import (
	"example.com/utils"
	"fmt"
	"time"
)

func SetupServer(name string) {
	fmt.Printf("Setting up server with name: \"%s\"\n", name)
	fmt.Println("GETTING HOSTNAME...")
	countdown(3)
	fmt.Printf("HOSTNAME: %s", utils.GetHostname())
	fmt.Println("IDENTIFYING INTERFACES...")
	countdown(5)
	printInterfaces(utils.GetInterfaces())
	fmt.Println("FINISHING SERVER SETUP...")
	countdown(2)
	fmt.Printf("All done! Server \"%s\" is ready to be used!\n", name)
}

func printInterfaces(interfacesIPs map[string]string) {
	fmt.Printf("NUMBER OF FOUND INTERFACES: %d\n", len(interfacesIPs))
	for inter, ip := range interfacesIPs {
		fmt.Printf("INTERFACE: [%s] | IPV4 ADDRESS: [%s]\n", inter, ip)
	}
}

func countdown(seconds int) {
	for i := seconds; i > 0; i-- {
		fmt.Printf("%d...", i)
		if i == 1 {
			fmt.Printf("DONE\n")
		}
		time.Sleep(1 * time.Second)
	}
}
