package main

import "strings"

type Model struct {
	prog string
	args []string
	opts Serialisable
}

func (re Model) begin() {
	// if config file doesn't exist, create it
	re.opts.serialise(Saver{})
	re.opts.serialise(Loader{})
}

func (re Model) end() {
}

func (re Model) getProgramName() string {
	p := strings.Split(re.prog, "/")
	return LastItemOfArray(p)
}
