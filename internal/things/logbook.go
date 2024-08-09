package things

import (
	_ "embed"
	"encoding/json"
	"os/exec"
	"time"

	"github.com/gleich/lumber/v2"
)

//go:embed logbook.js
var logbookScript string

func TodosFromLogbook(today bool) []Todo {
	lumber.Info("Loading todos from Thing's logbook")

	out, err := exec.Command("osascript", "-l", "JavaScript", "-e", logbookScript).CombinedOutput()
	if err != nil {
		lumber.Fatal(err, "loading todos from logbook failed")
	}

	var todos []Todo
	err = json.Unmarshal(out, &todos)
	if err != nil {
		lumber.Fatal(err, "failed to parse json from", string(out))
	}

	var filteredTodos []Todo
	now := time.Now()
	format := "01/02/2006"
	if today {
		for _, t := range todos {
			if t.CompletedAt.In(time.Local).Format(format) == now.Format(format) {
				filteredTodos = append(filteredTodos, t)
			}
		}
	} else {
		yesterday := now.AddDate(0, 0, -1)
		for _, t := range todos {
			if t.CompletedAt.In(time.Local).Format(format) == yesterday.Format(format) {
				filteredTodos = append(filteredTodos, t)
			}
		}
	}

	lumber.Success("Loaded", len(filteredTodos), "tasks from Thing's logbook")
	return filteredTodos
}
