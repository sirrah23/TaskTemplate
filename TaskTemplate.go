package main

import (
	"errors"
	"flag"
	"fmt"
)

func main() {
	//templatePtr := flag.String("template", "foo", "Template task name to replicate")
	//startPtr := flag.Int("start", 0, "The number to start at when generating tasks")
	endPtr := flag.Int("end", -1, "The number to stop at when generating tasks")
	//projectPtr := flag.String("project", "", "Project assignment for the newly generated tasks")
	flag.Parse()

	if *endPtr < 0 {
		fmt.Println(errors.New("No valid end provided"))
		return
	}
}
