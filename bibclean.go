package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pfandzelter/bibclean/pkg/bibtex"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Print("Usage: bibclean <bibfile>\n")
		os.Exit(1)
	}

	bibfile := os.Args[1]

	contents, err := ioutil.ReadFile(bibfile)

	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	elements, err := bibtex.Parse(contents)

	for _, element := range elements {

		fmt.Printf("%s\n", element)
	}
}
