package main

import "testing"

func TestGetProgramName(t *testing.T) {
	model := Model{}
	testresult := model.getProgramName()
	myname := "app"
	if testresult != myname {
		t.Errorf("Model.getProgramName was incorrect, got: %s, wanted: %s.", testresult, myname)
	}
}
