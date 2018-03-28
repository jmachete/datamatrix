package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

//RUN ON THE TERMINAL
//go build datamatrix.go && ./datamatrix

func main() {

	skus := []string{
		"AQUA-P-4DM-SUC-1X10LT",
		"AQUA-P-4DM-SUC-1X5LT",
		"AQUA-R-PAVIMENT-SUC-1X5LT",
		"AQUA-T-LLM-SUC-1X5LT",
		"AQUA-D-ELIMINA-SUC-1X5LT",
		"AQUA-D-EGF-SUC-1X5LT",
		"AQUA-S-ELIMINA-SUC-1X7LT",
		"AQUA-S-LCB-SUC-1X5LT",
		"BION-L-DDB-SUC-1X1LT",
		"AQUA-I-LIMPA-IN-SUC-1X1LT",
		"DESI-L-BANCAD-M-SUC-1X5LT",
		"PANO-M-WCM-SUC-1X5UN",
		"AQUA-V-LIMPA-VI-SUC-1X1LT",
		"AQUA-F-DFC-SUC-1X24LT",
		"AQUA-D-DLE-SUC-1X5LT",
		"TENS-L-GMH-SUC-1X5LT",
		"SUAV-A-ROUPA-SUC-1X5LT",
		"EMUL-B-DETERGEN-SUC-1X20KG",
		"DELA-S-DETERGEN-SUC-1X25KG",
		"AQUA-A-BP6-SUC-1X5LT",
	}

	// os.RemoveAll("images/")
	// os.MkdirAll("images/", os.ModePerm)

	for index, element := range skus {

		fmt.Printf("%d Data Matrix Created: %s.png\n", index, element)
		// Create the barcode
		qrCode, _ := qr.Encode(element, qr.M, qr.Auto)
		// Scale the barcode to 800x800 pixels
		qrCode, _ = barcode.Scale(qrCode, 800, 800)
		// create the output file
		file, _ := os.Create("images/" + element + ".png")
		defer file.Close()
		// encode the barcode as png
		png.Encode(file, qrCode)
	}
}
