package tasktemplate

import "fmt"
import "strings"
import "strconv"

func BuildTaskCommands(template, project string, start, end int) []string {
	var taskStrings = []string{}
	for i := start; i <= end; i++ {
		s := fmt.Sprintf("task add \"%s\" project:\"%s\"", template, project)
		idx := strconv.Itoa(i)
		s = strings.Replace(s, "#", idx, 1)
		taskStrings = append(taskStrings, s)
	}
	return taskStrings
}
