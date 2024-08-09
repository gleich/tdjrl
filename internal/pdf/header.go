package pdf

import (
	"fmt"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/gleich/tdjrl/internal/config"
	"github.com/go-pdf/fpdf"
)

func AppendHeader(doc fpdf.Pdf, conf config.Config) {
	doc.SetFont("Arial", "B", 20)
	doc.CellFormat(0, 10, conf.Name, "", 0, "L", false, 0, "")
	now := time.Now()
	doc.CellFormat(
		0,
		10,
		fmt.Sprintf(
			"%s, %s %s, %d",
			now.Weekday(),
			now.Month(),
			humanize.Ordinal(now.Day()),
			now.Year(),
		),
		"",
		0,
		"R",
		false,
		0,
		"",
	)
	doc.SetFont("Arial", "I", 9)
	doc.SetY(16)
	doc.CellFormat(
		0,
		10,
		fmt.Sprintf(
			"Generated %s",
			now.Format("01/02/06 @ 03:04 PM"),
		),
		"",
		0,
		"R",
		false,
		0,
		"",
	)
}
