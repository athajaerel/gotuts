package main

import "strings"

func Explode(x string, y string) []string {
	return strings.Split(x, y)
}

func Last(p []string) string {
	return p[len(p) - 1:][0]
}
