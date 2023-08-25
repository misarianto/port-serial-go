package main

import (
	"Go/scanner/portserial"
	"Go/scanner/printer"
	"Go/scanner/render"
	"log"
	"runtime"
)

func main() {

	switch runtime.GOOS {
	case "linux":
		portserial.PortSerialLinx()
		printer.ConnPrinterLinuxMac()
		render.RenderScannerResultLinux()
	case "windows":
		printer.ConnectionPrinterWin()
		render.RenderScannerResultWin()
	case "darwin":
		printer.ConnPrinterLinuxMac()
		render.RenderScannerResultMac()
	default:
		log.Fatal("Sistem operasi tidak didukung")
	}

}
