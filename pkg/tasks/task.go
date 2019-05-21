package tasks

import (
	"fmt"
	"math/rand"
)

type Task struct {
	Labels []string `json:"labels"`
	Text   string   `json:"text"`
}

func (task Task) String() string {
	return task.AsString(rnd)
}

func (task Task) AsString(rnd *rand.Rand) string {
	var labels = task.Labels
	if len(labels) == 0 {
		labels = DefaultLabels()
	}
	var label = labels[rnd.Intn(len(labels))]
	return fmt.Sprintf("%s: %s", label, task.Text)
}

func (task Task) Clone() Task {
	return Task{
		Labels: append([]string{}, task.Labels...),
		Text:   task.Text,
	}
}
