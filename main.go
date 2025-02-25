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

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"

var configPath = os.ExpandEnv("$HOME/.config/fish/config.fish")

func main() {
	// TODO!
}

func setProxy(address string) {
	file, err := os.OpenFile(configPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Couldn't open fish config file", err)
		return
	}
	defer file.Close()

	add := fmt.Sprintf(`
	set -gx http_proxy "%s"
	set -gx https_proxy "%s"
	set -gx ftp_proxy "%s"
	set -gx all_proxy "%s"
	set -gx no_proxy "%s"`, address, address, address, address, address)

	if _, err = file.WriteString(add); err != nil {
		fmt.Println("Couldn't append to config file", err)
		return
	}

	cmd := exec.Command("fish", "-c", "source "+configPath)
	if err = cmd.Run(); err != nil {
		fmt.Println("Couldn't source the config file", err)
		return
	}
	fmt.Println(Green + "Should be good to go!" + Reset)
}

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

	newLengh := len(lines) - 5
	if newLengh < 0 {
		newLengh = 0
	}
	lines = lines[:newLengh]

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
	return writer.Flush()
}
