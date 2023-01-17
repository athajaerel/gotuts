package main

import "fmt"
import "os"

func main() {
	fmt.Println("BEGIN")
	App{os.Args[0], os.Args[1:]}.run()
	fmt.Println("END")
}
