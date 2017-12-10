package tasktemplate

import "fmt"
import "strings"
import "strconv"

//Generate tasks based on a user-provided template
func BuildTaskCommands(template, project string, start, end int) []string {
	var taskStrings = []string{}
	for i := start; i <= end; i++ {
		//Actual TaskWarrior command
		s := fmt.Sprintf("task add \"%s\" project:\"%s\"", template, project)
		idx := strconv.Itoa(i)
		//# is placeholder for numbers
		//NOTE: Might want to add \# as an escape so users can have # in the task itself
		s = strings.Replace(s, "#", idx, 1)
		taskStrings = append(taskStrings, s)
	}
	return taskStrings
}
