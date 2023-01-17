package main

import "os"

func main() {
	App{os.Args[0], os.Args[1:]}.run()
}
