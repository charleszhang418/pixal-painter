package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/charleszhang418/pixal-painter/apptype"
	"github.com/charleszhang418/pixal-painter/pxcanvas"
	"github.com/charleszhang418/pixal-painter/swatch"
	"github.com/charleszhang418/pixal-painter/ui"
)

func main() {
	pixlApp := app.New()
	pixlWindow := pixlApp.NewWindow("Pixal Painter")
	state := apptype.State{
		BrushColor:     color.NRGBA{255, 255, 255, 255},
		SwatchSelected: 0,
	}

	var tempX int 
	var tempY int 

	fmt.Print("Please enter the size you want: ")
  	fmt.Scan(&tempX)

	fmt.Print("Please enter the pixal size you want: ")
	fmt.Scan(&tempY)

	pixlCanvasConfig := apptype.PxCanvasConfig{
		DrawingArea:  fyne.NewSize(600, 600),
		CanvasOffset: fyne.NewPos(0, 0),
		PxRows:       tempX,
		PxCols:       tempX,
		PxSize:       tempY,
	}

	pixlCanvas := pxcanvas.NewPxCanvas(&state, pixlCanvasConfig)

	appInit := ui.AppInit{
		PixlCanvas: pixlCanvas,
		PixlWindow: pixlWindow,
		State:      &state,
		Swatches:   make([]*swatch.Swatch, 0, 64),
	}
	ui.Setup(&appInit)
	appInit.PixlWindow.ShowAndRun()
}
