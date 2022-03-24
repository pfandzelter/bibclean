package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/pfandzelter/bibclean/pkg/bbl"
	"github.com/pfandzelter/bibclean/pkg/bibtex"
)

var types []string = []string{"inproceedings",
	"article",
	"book",
	"incollection",
	"mastersthesis",
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

type additionalFields map[string]map[string]struct{}

func (a *additionalFields) String() string {
	s := ""
	for t, f := range *a {
		s += t
		s += ":"
		for field := range f {
			s += field
			s += ","
		}
		s += "\n"
	}
	return s
}

func (a *additionalFields) Set(v string) error {
	s := strings.Split(v, ":")

	if len(s) != 2 {
		return errors.New("value does not have the correct format: use <type>:<field>")
	}

	fields, ok := (*a)[s[0]]

	if !ok {
		fields = make(map[string]struct{})
		(*a)[s[0]] = fields
	}

	fields[s[1]] = struct{}{}

	return nil
}

func main() {

	var bibfile, newfile, bblfile, shorten *string
	var shortenBooktitle, shortenAll bool
	var additional additionalFields = make(additionalFields)

	bibfile = flag.String("in", "", "input bibliography file")
	newfile = flag.String("out", "", "output bibliography file")
	bblfile = flag.String("bbl", "", "(optional) auxillary .bbl file to check which references have been used in the text")
	shorten = flag.String("shorten", "", "(optional) level of applied title shortening to conform with IEEE citation style, can be \"publication\" (shorten only proceeding and journal titles with some common abbreviations) or \"all\" (aggressive shortening including shortening titles, uses the full list of abbrevations)")
	flag.Var(&additional, "additional", "Additional fields for entries: specify as many as you like in the form \"--additional=article:booktitle --additional=techreport:address\" (this will add a \"booktitle\" field to \"@article\" entries and an \"address\" field to \"@techreport\" entries)")

	flag.Parse()

	incorrectUse := (*bibfile == "") || (*newfile == "")

	switch *shorten {
	case "all":
		shortenAll = true
		fallthrough
	case "publication":
		shortenBooktitle = true
	}

	if incorrectUse {
		flag.PrintDefaults()
		os.Exit(1)
	}

	contents, err := os.ReadFile(*bibfile)

	check(err)

	tmpFile := path.Join(path.Base(*newfile), ".bibclean.tmp")

	outTmp, err := os.Create(tmpFile)

	check(err)

	defer outTmp.Close()
	defer os.Remove(tmpFile)

	usebbl := (*bblfile != "")
	used := make(map[string]struct{})

	if usebbl {
		bblcontents, err := os.ReadFile(*bblfile)

		check(err)

		used, err = bbl.Parse(bblcontents)

		check(err)
	}

	elements, err := bibtex.Parse(contents, shortenBooktitle, shortenAll, additional)

	check(err)

	if usebbl {
		fmt.Fprintf(outTmp, "// --------------------\n// --- %s ---\n// --------------------\n\n", "USED ENTRIES")

		for _, t := range types {

			fmt.Fprintf(outTmp, "// --- %s ---\n\n", strings.ToUpper(t))

			for _, element := range elements {
				if element.Type == t {
					if _, ok := used[element.ID]; ok {
						fmt.Fprintf(outTmp, "%s\n\n", element)
					}
				}
			}
		}

		fmt.Fprintf(outTmp, "// ----------------------\n// --- %s ---\n// ----------------------\n\n", "UNUSED ENTRIES")
	}

	for _, t := range types {

		fmt.Fprintf(outTmp, "// --- %s ---\n\n", strings.ToUpper(t))

		for _, element := range elements {
			if element.Type == t {
				if _, ok := used[element.ID]; !usebbl || !ok {
					fmt.Fprintf(outTmp, "%s\n\n", element)
				}
			}
		}
	}

	outFile, err := os.Create(*newfile)

	check(err)

	defer outFile.Close()

	_, err = io.Copy(outFile, outTmp)

	check(err)
}
