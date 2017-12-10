package main

import (
	"TaskTemplate/lib"
	"errors"
	"flag"
	"fmt"
	"log"
	"os/exec"
)

func main() {

	templatePtr := flag.String("template", "foo", "Template task name to replicate")
	startPtr := flag.Int("start", 1, "The number to start at when generating tasks")
	endPtr := flag.Int("end", -1, "The number to stop at when generating tasks")
	projectPtr := flag.String("project", "", "Project assignment for the newly generated tasks")
	flag.Parse()

	// Do some basic validations
	if *endPtr < 0 {
		fmt.Println(errors.New("No valid end provided"))
		return
	}

	if *startPtr > *endPtr {
		fmt.Println(errors.New("Start must be less than end"))
		return
	}

	//Generate all the tasks to insert into TaskWarrior
	taskArray := tasktemplate.BuildTaskCommands(*templatePtr, *projectPtr, *startPtr, *endPtr)

	if len(taskArray) == 0 {
		fmt.Println(errors.New("Nothing to generate"))
		return
	}

	//Run the task command for each command that was generated
	for _, task := range taskArray {
		taskCommand := task.GenerateTaskAddCMD()
		cmd := exec.Command(taskCommand[0], taskCommand[1:]...)
		_, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("%s\n", err)
		}
	}
	fmt.Println(fmt.Sprintf("%d tasks added successfully", len(taskArray)))
}
