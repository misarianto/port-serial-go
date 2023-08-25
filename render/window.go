package render

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/otiai10/gosseract"
	"github.com/signintech/gopdf"
)

func RenderScannerResultWin() {
	cmd := exec.Command("sane-scan-image")
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Println("Error:", err)
		return
	}

	imageText, err := ocrImageWin(output)
	if err != nil {
		log.Println("OCR Error:", err)
		return
	}

	err = createPDFWin("output.pdf", imageText)
	if err != nil {
		log.Println("PDF Creation Error:", err)
		return
	}

	fmt.Println("PDF created successfully.")
}

func ocrImageWin(imageData []byte) (string, error) {
	client := gosseract.NewClient()
	defer client.Close()

	err := client.SetImageFromBytes(imageData)
	if err != nil {
		return "", err
	}

	return client.Text()
}

func createPDFWin(outputPath string, content string) error {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(nil, content)

	err := pdf.WritePdf(outputPath)
	if err != nil {
		return err
	}

	return nil
}
