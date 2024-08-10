package pdf

import (
	"fmt"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/go-pdf/fpdf"
)

func Header(doc fpdf.Pdf, title string) {
	doc.SetFont("Arial", "B", 17)
	doc.CellFormat(0, 10, title, "", 0, "L", false, 0, "")
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
	doc.SetY(14.5)
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
