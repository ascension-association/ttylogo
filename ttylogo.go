// Display a logo on a terminal screen
package main

import (
	"bytes"
	_ "embed"

	"fortio.org/log"
	"fortio.org/terminal/ansipixels"
)

//go:embed logo.png
var logo []byte

func main() {
    // initialize canvas
	cm := ansipixels.DetectColorMode()
	ap := ansipixels.NewAnsiPixels(0)
	if err := ap.Open(); err != nil {
		log.Fatalf("Not a terminal: %v", err)
	}
	ap.Color256 = cm.Color256
	ap.TrueColor = cm.TrueColor
	ap.Margin = 0
	ap.HideCursor()
	ap.ClearScreen()
	ap.SyncBackgroundColor()

	// display logo
	img, err := ap.DecodeImage(bytes.NewReader(logo))
	if err != nil {
		log.Fatalf("Image Load Error: %v", err)
	}
	ap.StartSyncMode()
	ap.ClearScreen()
	ap.ShowImages(img, 1.0, 0, 0)
	ap.EndSyncMode()

	// run indefinitely
	select{}
}


