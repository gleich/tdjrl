package main

import (
	"os"
	"time"

	"github.com/gleich/lumber/v2"
	"github.com/gleich/tdjrl/internal/config"
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

	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, conf.Name)
	err := pdf.OutputFileAndClose("out.pdf")
	if err != nil {
		lumber.Fatal(err, "failed to write PDF")
	}
}
