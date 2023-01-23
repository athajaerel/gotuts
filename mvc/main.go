package main

import "os"

func main() {
	Control{os.Args[0], os.Args[1:]}.run()
}
