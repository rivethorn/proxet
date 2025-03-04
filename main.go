// Proxet
//
// A simple app to change Fish environment variables so the system uses a
// custom proxy across the board
//
// By Hassan Gh

package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if the argument is -r, reset the proxy settings
	if len(os.Args) > 1 && os.Args[1] == "-r" {
		if err := resetProxy(); err != nil {
			fmt.Println(Red+"Couldn't reset proxy settings:", err, Reset)
		}
		return
	}
	// If the argument is -a, receive the proxy address and set it
	if len(os.Args) > 1 && os.Args[1] == "-a" {
		if len(os.Args) < 3 {
			fmt.Println(Red + "Please provide a proxy address" + Reset)
			return
		}
		setProxy(os.Args[2])
		return
	}
	// If the argument is -c, check if the proxy is set
	if len(os.Args) > 1 && os.Args[1] == "-c" {
		if isProxySet() {
			fmt.Println(Green + "Proxy is set" + Reset)
		} else {
			fmt.Println(Red + "Proxy is not set" + Reset)
		}
		return
	}
	// If no arguments are provided, print the usage
	fmt.Println("Usage:")
	fmt.Println("proxet -a <proxy_address> to set the proxy")
	fmt.Println("proxet -r to reset the proxy settings")
	fmt.Println("proxet -c to check if the proxy is set")
}
