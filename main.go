// Proxet
//
// A simple app to change Fish environment variables so the system uses a
// custom proxy across the board
//
// By Hassan Gh

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

// ANSI color codes for terminal output
var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"

// Path to the Fish shell configuration file
var configPath = os.ExpandEnv("$HOME/.config/fish/config.fish")

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
	// If no arguments are provided, print the usage
	fmt.Println("Usage:")
	fmt.Println("proxet -a <proxy_address> to set the proxy")
	fmt.Println("proxet -r to reset the proxy settings")
}

// setProxy sets the proxy address in the Fish shell configuration file
func setProxy(address string) {
	file, err := os.OpenFile(configPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(Red+"Couldn't open fish config file:", err, Reset)
		return
	}
	defer file.Close()

	// Add proxy settings to the configuration file
	add := fmt.Sprintf(`set -gx http_proxy "%s"
set -gx https_proxy "%s"
set -gx ftp_proxy "%s"
set -gx all_proxy "%s"
set -gx no_proxy "%s"`, address, address, address, address, address)

	if _, err = file.WriteString(add); err != nil {
		fmt.Println(Red+"Couldn't append to config file:", err, Reset)
		return
	}

	// Source the configuration file to apply changes
	if err = sourceConfig(); err != nil {
		fmt.Println(Red+"Couldn't source the config file:", err, Reset)
		return
	}
	fmt.Println(Green + "Proxy settings updated!" + Reset)
}

// resetProxy removes the proxy settings from the Fish shell configuration file
func resetProxy() error {
	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	// Remove the last 5 lines which contain the proxy settings
	newLength := len(lines) - 5
	if newLength < 0 {
		newLength = 0
	}
	lines = lines[:newLength]

	file, err = os.OpenFile(configPath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	// Source the configuration file to apply changes
	if err = sourceConfig(); err != nil {
		fmt.Println(Red+"Couldn't source the config file:", err, Reset)
	}

	fmt.Println(Green + "Proxy settings reset!" + Reset)
	return writer.Flush()
}

// sourceConfig sources the Fish shell configuration file to apply changes
func sourceConfig() error {
	cmd := exec.Command("fish", "-c", "source "+configPath)
	return cmd.Run()
}

// TODO
// - Support for other shells
