package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pfandzelter/bibclean/pkg/bbl"
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

func check(err error) {
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}

func main() {

	var bibfile, newfile, bblfile *string

	bibfile = flag.String("in", "", "input bibliography file")
	newfile = flag.String("out", "", "output bibliography file")
	bblfile = flag.String("bbl", "", "(optional) auxillary .bbl file to check which references have been used in the text")

	flag.Parse()

	contents, err := ioutil.ReadFile(*bibfile)

	check(err)

	out, err := os.Create(*newfile)

	check(err)

	defer out.Close()

	usebbl := (*bblfile != "")
	used := make(map[string]struct{})

	if usebbl {
		bblcontents, err := ioutil.ReadFile(*bblfile)

		check(err)

		used, err = bbl.Parse(bblcontents)

		check(err)
	}

	elements, err := bibtex.Parse(contents)

	check(err)

	if usebbl {
		fmt.Fprintf(out, "// --------------------\n// --- %s ---\n// --------------------\n\n", "USED ENTRIES")

		for _, t := range types {

			fmt.Fprintf(out, "// --- %s ---\n\n", strings.ToUpper(t))

			for _, element := range elements {
				if element.Type == t {
					if _, ok := used[element.ID]; ok {
						fmt.Fprintf(out, "%s\n\n", element)
					}
				}
			}
		}

		fmt.Fprintf(out, "// ----------------------\n// --- %s ---\n// ----------------------\n\n", "UNUSED ENTRIES")
	}

	for _, t := range types {

		fmt.Fprintf(out, "// --- %s ---\n\n", strings.ToUpper(t))

		for _, element := range elements {
			if element.Type == t {
				if _, ok := used[element.ID]; !usebbl || !ok {
					fmt.Fprintf(out, "%s\n\n", element)
				}
			}
		}
	}
}
