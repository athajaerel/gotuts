package main

import "fmt"
import "strconv"

type Serialiser interface {
	ioi(int)
	ios(string)
}

type Serialisable interface {
	serialise(Serialiser)
}

// type Sizer?
type Saver struct {
}

func (re Saver) ioi(i int) {
}

func (re Saver) ios(s string) {
}

type Loader struct {
}

func (re Loader) ioi(i int) {
}

func (re Loader) ios(s string) {
}

type Printer struct {
}

func (re Printer) ioi(i int) {
	fmt.Println(strconv.Itoa(i))
}

func (re Printer) ios(s string) {
	fmt.Println(s)
}
