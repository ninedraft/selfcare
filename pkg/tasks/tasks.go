package tasks

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"strings"
)

type Tasks struct {
	rnd   *rand.Rand
	tasks []Task
}

func NewTasks() *Tasks {
	return &Tasks{
		rnd:   newRnd(),
		tasks: []Task{},
	}
}

func NewTasksFromList(list []Task) *Tasks {
	var tasks = NewTasks()
	tasks.tasks = list
	return tasks
}

func ParseTasks(re io.Reader) (*Tasks, error) {
	var taskSet = make(map[string][]string)
	var scanner = bufio.NewScanner(re)
	for scanner.Scan() {
		var line = scanner.Text()
		var tokens = strings.SplitN(line, ":", 2)
		switch len(tokens) {
		case 1:
			var taskText = strings.TrimSpace(tokens[0])
			taskSet[taskText] = DefaultLabels()
		case 2:
			var label = strings.TrimSpace(tokens[0])
			var taskText = strings.TrimSpace(tokens[1])
			taskSet[taskText] = append(taskSet[taskText], label)
		default:
			continue
		}
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	var taskList = make([]Task, 0, len(taskSet))
	for text, labels := range taskSet {
		taskList = append(taskList, Task{
			Text:   text,
			Labels: labels,
		})
	}
	var tasks = NewTasks()
	tasks.tasks = taskList
	return tasks, nil
}

func (tasks *Tasks) String() string {
	return fmt.Sprintf("Tasks{%d self care tasks}", len(tasks.tasks))
}

func (tasks *Tasks) NTasks() int {
	return len(tasks.tasks)
}

func (tasks *Tasks) TaskList() []Task {
	var list = make([]Task, 0, tasks.NTasks())
	for _, task := range tasks.tasks {
		list = append(list, task.Clone())
	}
	return list
}

func (tasks *Tasks) PeekN(n int) []Task {
	var nTasks = len(tasks.tasks)
	if n < 0 || n > nTasks {
		n = nTasks
	}
	var sample = make([]Task, 0, nTasks)
	for _, i := range tasks.rnd.Perm(nTasks)[:n] {
		sample = append(sample, tasks.tasks[i])
	}
	return sample
}

func (tasks *Tasks) PeekNasStrings(n int) []string {
	var sample = make([]string, 0, len(tasks.tasks))
	for _, task := range tasks.PeekN(n) {
		sample = append(sample, task.AsString(tasks.rnd))
	}
	return sample
}

var (
	_ json.Marshaler   = new(Tasks)
	_ json.Unmarshaler = new(Tasks)
)

func (tasks *Tasks) MarshalJSON() ([]byte, error) {
	return json.Marshal(tasks.tasks)
}

func (tasks *Tasks) UnmarshalJSON(data []byte) error {
	var taskList []Task
	if err := json.Unmarshal(data, &taskList); err != nil {
		return err
	}
	tasks.tasks = taskList
	if tasks.rnd == nil {
		tasks.rnd = newRnd()
	}
	return nil
}
