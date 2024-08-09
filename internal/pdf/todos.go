package pdf

import (
	"github.com/gleich/tdjrl/internal/things"
	"github.com/go-pdf/fpdf"
)

func WriteTodos(doc fpdf.Pdf, todos []things.Todo) {
	x := 10.0
	y := 30.0
	gap := 5.0
	doc.SetXY(x, y)
	for i, todo := range todos {
		doc.SetFont("Arial", "B", 10)
		doc.Cell(17.5, 5, todo.CompletedAt.Format("03:04 PM"))
		doc.SetFont("Arial", "", 10)
		doc.Cell(20, 5, todo.Name)
		doc.SetXY(x, (float64(i+1))*gap+y)
	}
}
