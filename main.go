package main

import (
	"os"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/gleich/lumber/v2"
	"github.com/gleich/tdjrl/internal/pdf"
	"github.com/gleich/tdjrl/internal/things"
	"github.com/go-pdf/fpdf"
)

func main() {
	var title string
	huh.NewInput().Title("What is the title?").Value(&title).Run()

	logger := lumber.NewCustomLogger()
	logger.Timezone = time.Local
	lumber.SetLogger(logger)

	today := true
	for _, a := range os.Args {
		if a == "--yesterday" {
			today = false
		}
	}

	todos := things.TodosFromLogbook(today)

	lumber.Info("Creating PDF")
	doc := fpdf.New("P", "mm", "A4", "")
	doc.AddPage()
	pdf.Header(doc, title)
	pdf.Todos(doc, todos)
	pdf.Lines(doc)

	filename := "out.pdf"
	err := doc.OutputFileAndClose(filename)
	if err != nil {
		lumber.Fatal(err, "failed to write PDF")
	}
	lumber.Success("Output to PDF:", filename)
}
