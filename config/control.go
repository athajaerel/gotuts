package main

type Control struct {
	prog string
	args []string
}

type TestData struct {
	a string
	b int32
	c string
	d int32
}

func (re TestData) serialise(s Serialiser) {
	s.ios(&re.a)
	s.ioi32(&re.b)
	s.ios(&re.c)
	s.ioi32(&re.d)
}

func (re Control) run() {
	t := TestData{"hello", 45, "world", -12}
	model := Model{re.prog, re.args, t}
	view := View{re.prog, re.args}
	model.begin()
	view.begin()
	view.output("Program Name: " + model.getProgramName())
	view.outputTestData(t)
	// check for script
	// run through modules
	view.end()
	model.end()
}
