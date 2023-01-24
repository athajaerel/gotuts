package main

import "testing"

func TestLastItemOfStringArray(t *testing.T) {
	test := [5]string{"one", "two", "three", "four", "five"}
	testresult := LastItemOfStringArray(test[:])  // [:] converts to slice
	if testresult != "five" {
		t.Errorf("LastItemOfStringArray was incorrect, got: %s, wanted: %s.", testresult, "five")
	}
}
