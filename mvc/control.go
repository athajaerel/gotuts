package main

import "os"

type Control struct {
	prog string
	args []string
}

func (re Control) run() {
	model := Model{os.Args[0], os.Args[1:], re}
	view := View{os.Args[0], os.Args[1:], re}
	model.begin()
	view.begin()
	model.load_data_file_into_struct()
	view.output("Program Name: " + model.getProgramName())
	// check for script
	// run through modules
	view.end()
	model.end()
}
