package main

import "fmt"
import "os"
import "strings"

type Saver interface {
}

type Loader interface {
}

type AppI interface {
    run()
}

type App struct {
    AppI
    prog string
    args []string
}

func Explode(x string, y string) []string {
    return strings.Split(x, y)
}

func Last(p []string) string {
    return p[len(p) - 1:][0]
}

func (re App) run() {
    // load data file
    // into data structure
    // select UI
    p := Explode(re.prog, "/")
    fmt.Println("Program Name: ", Last(p))
    // check for script
    // run through modules
}

func main() {
    fmt.Println("BEGIN")
    var appi AppI
    var app = App{appi, os.Args[0], os.Args[1:]}
    app.run()
    fmt.Println("END")
}
