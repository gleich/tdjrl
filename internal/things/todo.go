package things

import "time"

type Todo struct {
	Name        string    `json:"name"`
	Status      string    `json:"status"`
	Notes       string    `json:"notes"`
	Tags        string    `json:"tags"`
	ID          string    `json:"id"`
	CompletedAt time.Time `json:"completed_date"`
	Project     struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Tags string `json:"tags"`
		Area *Area  `json:"area"`
	} `json:"project"`
	Area *Area `json:"area"`
}

type Area struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Tags string `json:"tags"`
}
