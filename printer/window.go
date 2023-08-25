package printer

import (
	"fmt"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

func ConnectionPrinterWin() {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	// Buat objek untuk mengakses Windows Management Instrumentation (WMI).
	unknown, err := oleutil.CreateObject("WbemScripting.SWbemLocator")
	if err != nil {
		fmt.Printf("Gagal membuat objek SWbemLocator: %v", err)
		return
	}
	defer unknown.Release()

	wmi, err := unknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		fmt.Printf("Gagal membuat objek IDispatch: %v", err)
		return
	}
	defer wmi.Release()

	// Buat objek untuk mengakses printer.
	printerSvc, err := oleutil.CallMethod(wmi, "ConnectServer", nil, "ROOT\\CIMv2")
	if err != nil {
		fmt.Printf("Gagal membuat objek printer: %v", err)
		return
	}
	defer printerSvc.ToIDispatch().Release()

	// Eksekusi query WMI untuk mendapatkan daftar printer.
	query := "SELECT * FROM Win32_Printer"
	printers, err := oleutil.CallMethod(printerSvc.ToIDispatch(), "ExecQuery", query)
	if err != nil {
		fmt.Printf("Gagal menjalankan query WMI: %v", err)
		return
	}
	defer printers.ToIDispatch().Release()

	// Iterate melalui daftar printer dan tampilkan informasinya.
	enum := printers.ToIDispatch()
	defer enum.Release()

	count, err := oleutil.GetProperty(enum, "Count")
	if err != nil {
		fmt.Printf("Gagal mendapatkan jumlah printer: %v", err)
		return
	}

	numPrinters := int(count.Val)
	for i := 0; i < numPrinters; i++ {
		printer, err := oleutil.CallMethod(enum, "ItemIndex", i)
		if err != nil {
			fmt.Printf("Gagal mendapatkan printer: %v", err)
			continue
		}
		defer printer.ToIDispatch().Release()

		name, err := oleutil.GetProperty(printer.ToIDispatch(), "Name")
		if err != nil {
			fmt.Printf("Gagal mendapatkan nama printer: %v", err)
			continue
		}

		fmt.Printf("Printer: %s\n", name.ToString())
	}
}
