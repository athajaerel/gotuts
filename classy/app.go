package main

import "fmt"

type App struct {
	prog string
	args []string
}

func (re App) run() {
	fmt.Println("BEGIN")
	// load data file
	// into data structure
	// select UI
	p := Explode(re.prog, "/")
	fmt.Println("Program Name: ", Last(p))
	// check for script
	// run through modules
	fmt.Println("END")
}
