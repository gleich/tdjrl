package things

import (
	_ "embed"
	"encoding/json"
	"os/exec"
	"time"

	"github.com/briandowns/spinner"
	"github.com/gleich/lumber/v2"
)

//go:embed logbook.js
var logbookScript string

func TodosFromLogbook() []Todo {
	s := spinner.New(spinner.CharSets[14], 20*time.Millisecond)
	s.Suffix = " Loading tasks from Things's logbook"
	s.Start()
	out, err := exec.Command("osascript", "-l", "JavaScript", "-e", logbookScript).CombinedOutput()
	if err != nil {
		lumber.Fatal(err, "loading todos from logbook failed")
	}

	var todos []Todo
	err = json.Unmarshal(out, &todos)
	if err != nil {
		lumber.Fatal(err, "failed to parse json from", string(out))
	}
	s.Stop()
	return todos
}
