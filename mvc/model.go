package main

type Model struct {
	prog string
	args []string
	control Control
}

func (re Model) begin() {
}

func (re Model) end() {
}

func (re Model) load_data_file_into_struct() {
}

func (re Model) getProgramName() string {
	p := Explode(re.prog, "/")
	return Last(p)
}
