package main

import (
	"github.com/gleich/lumber/v2"
	"github.com/gleich/tdjrl/internal/things"
)

func main() {
	todos := things.TodosFromLogbook()
	for _, todo := range todos {
		lumber.Debug(todo.Name)
	}
}
