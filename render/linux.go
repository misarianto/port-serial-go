package render

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/jung-kurt/gofpdf"
	"github.com/otiai10/gosseract"
)

func RenderScannerResultLinux() {
	cmd := exec.Command("scanimage", "--format=png")
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Println("Error:", err)
		return
	}

	imageText, err := OcrImage(output)
	if err != nil {
		log.Println("OCR Error:", err)
		return
	}

	err = createPDF("output.pdf", imageText)
	if err != nil {
		log.Println("PDF Creation Error:", err)
		return
	}

	fmt.Println("PDF created successfully.")
}

func OcrImage(imageData []byte) (string, error) {
	client := gosseract.NewClient()
	defer client.Close()

	err := client.SetImageFromBytes(imageData)
	if err != nil {
		return "", err
	}

	return client.Text()
}

func createPDF(outputPath string, content string) error {
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
