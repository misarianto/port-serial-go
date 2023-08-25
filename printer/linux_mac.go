package printer

import (
	"fmt"
	"os/exec"
	"strings"
)

func ConnPrinterLinuxMac() {
	cmd := exec.Command("lpstat", "-p")
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	printers := strings.Split(strings.TrimSpace(string(output)), "\n")
	for _, printer := range printers {
		fmt.Printf("Printer: %s\n", printer)
	}
}
