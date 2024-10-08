package things

import (
	_ "embed"
	"encoding/json"
	"os/exec"
	"sort"
	"time"

	"github.com/gleich/lumber/v2"
)

//go:embed logbook.js
var logbookScript string

// Loads todos from Thing's logbook.
// Either loads the tasks from today or yesterday
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
	format := "01/02/2006"
	now := time.Now()
	if !today {
		now = now.AddDate(0, 0, -1)
	}
	for _, t := range todos {
		t.CompletedAt = t.CompletedAt.In(time.Local)
		if t.CompletedAt.Format(format) == now.Format(format) {
			filteredTodos = append(filteredTodos, t)
		}
	}

	sort.Slice(
		filteredTodos,
		func(i, j int) bool { return filteredTodos[j].CompletedAt.After(filteredTodos[i].CompletedAt) },
	)

	lumber.Success("Loaded", len(filteredTodos), "tasks from Thing's logbook")
	return filteredTodos
}
