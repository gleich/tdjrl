package main

import (
	"os"
	"time"

	"github.com/gleich/lumber/v2"
	"github.com/gleich/tdjrl/internal/things"
)

func main() {
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
}
