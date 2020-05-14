package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pfandzelter/bibclean/pkg/bibtex"
)

var types []string = []string{"inproceedings",
	"article",
	"book",
	"incollection",
	"masterthesis",
	"misc",
	"phdthesis",
	"techreport",
	"unpublished",
}

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

	for _, t := range types {

		fmt.Printf("// --- %s ---\n\n", strings.ToUpper(t))

		for _, element := range elements {
			if element.Type == t {
				fmt.Printf("%s\n\n", element)
			}
		}
	}
}
