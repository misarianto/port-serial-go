package render

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/jung-kurt/gofpdf"
	"github.com/otiai10/gosseract"
)

func RenderScannerResultMac() {
	cmd := exec.Command("sane-scan-image")
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Println("Error:", err)
		return
	}

	imageText, err := ocrImageMac(output)
	if err != nil {
		log.Println("OCR Error:", err)
		return
	}

	err = createPDFMac("output.pdf", imageText)
	if err != nil {
		log.Println("PDF Creation Error:", err)
		return
	}

	fmt.Println("PDF created successfully.")
}

func ocrImageMac(imageData []byte) (string, error) {
	client := gosseract.NewClient()
	defer client.Close()

	err := client.SetImageFromBytes(imageData)
	if err != nil {
		return "", err
	}

	return client.Text()
}

func createPDFMac(outputPath string, content string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.MultiCell(0, 10, content, "", "", false)

	err := pdf.OutputFileAndClose(outputPath)
	if err != nil {
		return err
	}

	return nil
}
