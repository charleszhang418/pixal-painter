package ui

import (
	"fyne.io/fyne/v2"
	"github.com/charleszhang418/pixal-painter/apptype"
	"github.com/charleszhang418/pixal-painter/pxcanvas"
	"github.com/charleszhang418/pixal-painter/swatch"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
