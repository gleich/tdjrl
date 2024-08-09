package main

import (
	"time"

	"github.com/gleich/lumber/v2"
	"github.com/gleich/tdjrl/internal/things"
)

func main() {
	logger := lumber.NewCustomLogger()
	logger.Timezone = time.Local
	lumber.SetLogger(logger)

	todos := things.TodosFromLogbook(false)
	for _, todo := range todos {
		lumber.Debug(todo.Name)
	}
}
