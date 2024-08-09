package things

import (
	_ "embed"
	"encoding/json"
	"os/exec"

	"github.com/gleich/lumber/v2"
)

//go:embed logbook.js
var logbookScript string

func TodosFromLogbook() []Todo {
	out, err := exec.Command("osascript", "-l", "JavaScript", "-e", logbookScript).CombinedOutput()
	if err != nil {
		lumber.Fatal(err, "loading todos from logbook failed")
	}

	var todos []Todo
	err = json.Unmarshal(out, &todos)
	if err != nil {
		lumber.Fatal(err, "failed to parse json from", string(out))
	}
	return todos
}
