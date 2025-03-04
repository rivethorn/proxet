package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// ANSI color codes for terminal output
var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"

// Path to the Fish shell configuration file
var configPath = os.ExpandEnv("$HOME/.config/fish/config.fish")

// setProxy sets the proxy address in the Fish shell configuration file
func setProxy(address string) {
	if isProxySet() {
		fmt.Println(Yellow + "Proxy is already set" + Reset)
		return
	}

	file, err := os.OpenFile(configPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(Red+"Couldn't open fish config file:", err, Reset)
		return
	}
	defer file.Close()

	// Add proxy settings to the configuration file
	add := fmt.Sprintf(`
set -gx http_proxy "%s"
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
	if !isProxySet() {
		fmt.Println(Yellow + "No proxy settings to reset" + Reset)
		return nil
	}

	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Skip lines that contain proxy settings
		if strings.Contains(line, "_proxy") {
			continue
		}
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	file, err = os.OpenFile(configPath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i, line := range lines {
		if i > 0 {
			_, err := writer.WriteString("\n")
			if err != nil {
				return err
			}
		}
		_, err := writer.WriteString(line)
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

// isProxySet checks if the proxy is set in the Fish shell configuration file
func isProxySet() bool {
	file, err := os.Open(configPath)
	if err != nil {
		fmt.Println(Red+"Couldn't open fish config file:", err, Reset)
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "http_proxy") || strings.Contains(line, "https_proxy") ||
			strings.Contains(line, "ftp_proxy") || strings.Contains(line, "all_proxy") ||
			strings.Contains(line, "no_proxy") {
			return true
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(Red+"Error reading fish config file:", err, Reset)
		return false
	}
	return false
}

// sourceConfig sources the Fish shell configuration file to apply changes
func sourceConfig() error {
	cmd := exec.Command("fish", "-c", "source "+configPath)
	return cmd.Run()
}

// !TODO
// - Support for other shells
// - Support for User-level and System-level
