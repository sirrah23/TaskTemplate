package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"tasktemplate"
)

func main() {

	templatePtr := flag.String("template", "foo", "Template task name to replicate")
	startPtr := flag.Int("start", 0, "The number to start at when generating tasks")
	endPtr := flag.Int("end", -1, "The number to stop at when generating tasks")
	projectPtr := flag.String("project", "", "Project assignment for the newly generated tasks")
	flag.Parse()

	if *endPtr < 0 {
		fmt.Println(errors.New("No valid end provided"))
		return
	}

	taskArray := tasktemplate.BuildTaskCommands(*templatePtr, *projectPtr, *startPtr, *endPtr)
	for _, task := range taskArray {
		taskCommand := strings.SplitN(task, " ", 2)
		cmd := exec.Command(taskCommand[0], taskCommand[1])
		_, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("%s\n", err)
		}
	}
	fmt.Println("%d tasks added successfully", len(taskArray))
}
