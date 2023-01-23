package main

import "fmt"

type View struct {
	prog string
	args []string
}

func (re View) begin() {
	fmt.Println("BEGIN")
	// select UI
}

func (re View) end() {
	fmt.Println("END")
}

func (re View) output(str string) {
	fmt.Println(str)
}

func (re View) outputTestData(t TestData) {
	t.serialise(Printer{})
}
