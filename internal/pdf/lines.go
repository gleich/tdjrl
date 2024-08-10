package pdf

import "github.com/go-pdf/fpdf"

func Lines(doc fpdf.Pdf) {
	width, height := doc.GetPageSize()
	y := doc.GetY() + 10
	bottomMargin := 10
	for y+float64(bottomMargin) < height {
		doc.Rect(10, y, width-20, 0.1, "F")
		y += 8
	}
}
