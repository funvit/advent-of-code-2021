// +build gen

package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed soltpl.gotext
var tpl string

var errLog = log.New(os.Stderr, "ERR ", log.LstdFlags)
var warnLog = log.New(os.Stderr, "WRN ", log.LstdFlags)

func main() {
	fmt.Println("Solution file generator")

	t, err := template.New("").Parse(tpl)
	if err != nil {
		errLog.Println("template parse:", err)
		os.Exit(1)
	}

	var day int
	flag.IntVar(&day, "day", 0, "day number (ex: 1)")
	flag.Parse()

	if day <= 0 {
		errLog.Println("wrong day number")
	}

	folder := filepath.Join(".", fmt.Sprintf("day%2d", day))

	_, err = os.Stat(folder)
	if err == nil {
		warnLog.Println("folder already exists, wrong day?")
		os.Exit(1)
	}

	err = os.Mkdir(folder, 0755)
	if err != nil {
		errLog.Println(err)
		os.Exit(1)
	}

	f, err := os.Create(filepath.Join(folder, "sol.go"))
	if err != nil {
		errLog.Println("create solution go-file:", err)
		os.Exit(1)
	}
	defer f.Close()

	type data struct {
		Day int
	}

	err = t.Execute(f, data{Day: day})
	if err != nil {
		errLog.Println("process template:", err)
		os.Exit(1)
	}

	log.Println("folder crated:", folder)

	fmt.Println("Done.")
}
