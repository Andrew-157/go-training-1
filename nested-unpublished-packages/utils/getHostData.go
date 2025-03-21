// Package that has some dummy util functions for dummy server
package utils

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

// `execute` runs the provided command and returns its output as a string.
// If the command returns a non-zero exit code, the function logs the error and exits the program.
//
// Parameters:
//   - cmd: The command to execute, including arguments.
//
// Returns:
//
//	The output of the command if it executes successfully. Exits the program if an error occurs.
func execute(cmd string) string {
	splittedCmd := strings.Split(cmd, " ")
	out, err := exec.Command(splittedCmd[0], splittedCmd[1:]...).Output()
	if err != nil {
		log.Fatalf("Error happened during execution of `%v`: %v\n", cmd, err)
		os.Exit(1)
	}
	return string(out)
}

// `GetHostname` returns hostname of the host
func GetHostname() string {
	return execute("hostname")
}

// `GetInterfaces` returns map of all interfaces on host as keys
// and their IP(in IPv4 format) addresses as values using `ip addr` command
func GetInterfaces() map[string]string {
	interfaceIP := make(map[string]string)
	interfaceSectionLen := 6 // Number of lines for one interface in `ip a` output
	splitted := strings.Split(execute("ip addr"), "\n")
	numberOfInterfaces := len(splitted) / interfaceSectionLen
	j := 0
	for i := 1; i <= numberOfInterfaces; i++ {
		currentInterface := splitted[j : i*interfaceSectionLen]
		interfaceName := strings.Split(currentInterface[0], " ")[1]
		interfaceName = interfaceName[:len(interfaceName)-1] // omit `:` from interface name, e.g. "lo:" -> "lo"
		interfaceIPv4Address := strings.Split(currentInterface[2], " ")[5]
		interfaceIP[interfaceName] = interfaceIPv4Address
		j += interfaceSectionLen
	}
	return interfaceIP
}
