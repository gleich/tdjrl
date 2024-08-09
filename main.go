package main

import (
	"os"
	"time"

	"github.com/gleich/lumber/v2"
	"github.com/gleich/tdjrl/internal/config"
	"github.com/gleich/tdjrl/internal/pdf"
	"github.com/gleich/tdjrl/internal/things"
	"github.com/go-pdf/fpdf"
)

func main() {
	conf := config.Read()

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
	for _, todo := range todos {
		lumber.Debug(todo.Name)
	}

	doc := fpdf.New("P", "mm", "A4", "")
	doc.AddPage()
	pdf.AppendHeader(doc, conf)
	pdf.WriteTodos(doc, todos)

	err := doc.OutputFileAndClose("out.pdf")
	if err != nil {
		lumber.Fatal(err, "failed to write PDF")
	}
}
