package portserial

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func PortSerialLinx() {

	// List all devices in /dev
	devices, err := ioutil.ReadDir("/dev")
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through the devices and check if any printer-related names are present
	for _, device := range devices {
		deviceName := device.Name()

		// Check for common printer-related strings in the device name
		if strings.Contains(deviceName, "ttyS") || strings.Contains(deviceName, "ttyUSB") || strings.Contains(deviceName, "lp") {
			fmt.Printf("Possible printer connection: /dev/%s\n", deviceName)
		}
	}

	// Check for other common printer-related directories
	printerDirs := []string{"/var/run/cups", "/var/spool/cups"}
	for _, dir := range printerDirs {
		_, err := os.Stat(dir)
		if err == nil {
			fmt.Printf("Possible printer connection: %s\n", dir)
		}
	}
}
