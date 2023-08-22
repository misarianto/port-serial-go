package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/tarm/serial"
)

func main() {

	// Membaca port serial berdasarkan sistem operasi
	var serialPort string
	switch runtime.GOOS {
	case "linux":
		serialPort = "/dev/ttyUSB0" // Ganti dengan port serial yang sesuai di Linux
	case "windows":
		serialPort = "COM45" // Ganti dengan port serial yang sesuai di Windows
	case "darwin":
		serialPort = "/dev/cu.usbserial" // Ganti dengan port serial yang sesuai di macOS
	default:
		log.Fatal("Sistem operasi tidak didukung")
	}

	// Konfigurasi port serial
	config := &serial.Config{
		Name: serialPort, // Sesuaikan dengan port serial Anda
		Baud: 9600,
	}

	// Buka port serial
	port, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	// Baca data dari port serial
	buf := make([]byte, 128)
	n, err := port.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	// Tampilkan data yang dibaca
	fmt.Printf("Received: %s", buf[:n])
}
