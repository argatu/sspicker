package main

import (
	"fmt"
	"os/exec"
	"strings"
)

const ruleName = "sspicker"

func addRule(blockedIps []string) error {
	args := []string{"advfirewall", "firewall", "add", "rule", "name=" + ruleName, "dir=out", "action=block", "remoteip=" + strings.Join(blockedIps, ",")}
	_, err := exec.Command("netsh", args...).Output()
	if err != nil {
		return err
	}

	fmt.Println("Firewall rule has been added")
	return nil
}

func removeRule() error {
	cmd := exec.Command("netsh", "advfirewall", "firewall", "delete", "rule", "name="+ruleName)
	output, err := cmd.Output()
	if err != nil {
		return err
	}

	fmt.Print("\nFirewall rule has been removed")
	fmt.Println(string(output))

	return nil
}
