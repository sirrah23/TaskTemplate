package tasktemplate

import (
	"fmt"
	"strconv"
	"strings"
)

type Task struct {
	description, project string
}

func (t Task) GenerateTaskAddCMD() []string {
	return []string{
		"task",
		"add",
		t.description,
		fmt.Sprintf("project:\"%s\"", t.project),
	}
}

//Generate tasks based on a user-provided template
func BuildTaskCommands(template, project string, start, end int) []Task {
	tasks := []Task{}
	for i := start; i <= end; i++ {
		description := strings.Replace(template, "#", strconv.Itoa(i), 1)
		tasks = append(tasks, Task{description, project})
	}
	return tasks
}
