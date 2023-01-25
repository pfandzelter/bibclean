package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/pfandzelter/bibclean/pkg/bbl"
	"github.com/pfandzelter/bibclean/pkg/bibtex"
)

// update this version when making changes by tagging the commit
// compile with make to get all this information automatically
// OR go build -ldflags "-X main.version=$(git describe --tags --always) -X main.commit=$(shell git rev-parse HEAD) -X main.date=$(shell date -u '+%Y-%m-%d_%I:%M:%S%p') -X main.builtBy=goreleaser".
// OR goreleaser will do this automatically

var version = "unknown"
var commit = "unknown"
var date = "unknown"

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

// Entry types
var (
	ieeeElements = &map[string][]string{
		"article":       {"author", "title", "journal", "year", "volume", "number", "pages", "publisher"},
		"book":          {"author", "editor", "title", "publisher", "year"},
		"incollection":  {"author", "title", "booktitle", "publisher", "year"},
		"inproceedings": {"author", "title", "booktitle", "pages", "month", "year"},

		"mastersthesis": {"author", "title", "school", "month", "year"},

		"misc": {"author", "title", "howpublished", "month", "year", "note", "publisher"},

		"phdthesis": {"author", "title", "school", "month", "year"},

		"techreport": {"author", "title", "institution", "booktitle", "month", "year"},

		"unpublished": {"author", "title", "month", "year", "note"},
	}
)

// https://www.acm.org/publications/authors/bibtex-formatting
var (
	acmElements = &map[string][]string{
		"article":       {"author", "title", "journal", "issue_date", "volume", "number", "month", "year", "issn", "pages", "articleno", "numpages", "url", "doi", "acmid", "publisher", "address", "issue_date", "Eprint"},
		"book":          {"author", "title", "year", "isbn", "publisher", "address", "editor"},
		"incollection":  {"author", "title", "booktitle", "publisher", "pages", "year"},
		"inproceedings": {"author", "title", "booktitle", "pages", "month", "year", "acmid", "publisher", "address", "series", "location", "numpages", "url", "doi"},

		"mastersthesis": {"author", "title", "school", "month", "year"},

		"online": {"author", "title", "url", "month", "year", "lastaccessed"},

		"phdthesis": {"author", "title", "publisher", "address", "month", "year"},

		"techreport": {"author", "title", "publisher", "address", "source", "booktitle", "month", "year"},
	}
)

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

	var printVersion *bool
	var bibfile, newfile, bblfile, shorten *string
	var acmDefaults *bool
	var shortenBooktitle, shortenAll bool
	var additional additionalFields = make(additionalFields)

	printVersion = flag.Bool("version", false, "print bibclean version and exit")
	bibfile = flag.String("in", "", "input bibliography file")
	newfile = flag.String("out", "", "output bibliography file")
	bblfile = flag.String("bbl", "", "(optional) auxillary .bbl file to check which references have been used in the text")
	acmDefaults = flag.Bool("acm-defaults", false, "(optional) use ACM defaults instead of IEEE for default entries and fields")
	shorten = flag.String("shorten", "", "(optional) level of applied title shortening to conform with IEEE citation style, can be \"publication\" (shorten only proceeding and journal titles with some common abbreviations) or \"all\" (aggressive shortening including shortening titles, uses the full list of abbrevations)")
	flag.Var(&additional, "additional", "Additional fields for entries: specify as many as you like in the form \"--additional=article:booktitle --additional=techreport:address\" (this will add a \"booktitle\" field to \"@article\" entries and an \"address\" field to \"@techreport\" entries)")

	flag.Parse()

	if *printVersion {
		fmt.Printf("bibclean\nversion %s\nbuilt %s\ncommit %s\n", version, date, commit)
		os.Exit(0)
	}

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

	newfilePath, err := filepath.Abs(*newfile)

	check(err)

	contents, err := os.ReadFile(*bibfile)

	check(err)

	buf := bytes.Buffer{}

	usebbl := (*bblfile != "")
	used := make(map[string]struct{})

	if usebbl {
		bblcontents, err := os.ReadFile(*bblfile)

		check(err)

		used, err = bbl.Parse(bblcontents)

		check(err)
	}

	defaultElements := ieeeElements
	if *acmDefaults {
		defaultElements = acmElements
	}

	elements, err := bibtex.Parse(contents, shortenBooktitle, shortenAll, defaultElements, additional)

	check(err)

	if usebbl {
		fmt.Fprintf(&buf, "// --------------------\n// --- %s ---\n// --------------------\n\n", "USED ENTRIES")

		for _, t := range types {

			fmt.Fprintf(&buf, "// --- %s ---\n\n", strings.ToUpper(t))

			for _, element := range elements {
				if element.Type == t {
					if _, ok := used[element.ID]; ok {
						fmt.Fprintf(&buf, "%s\n\n", element)
					}
				}
			}
		}

		fmt.Fprintf(&buf, "// ----------------------\n// --- %s ---\n// ----------------------\n\n", "UNUSED ENTRIES")
	}

	for _, t := range types {

		fmt.Fprintf(&buf, "// --- %s ---\n\n", strings.ToUpper(t))

		for _, element := range elements {
			if element.Type == t {
				if _, ok := used[element.ID]; !usebbl || !ok {
					fmt.Fprintf(&buf, "%s\n\n", element)
				}
			}
		}
	}

	outFile, err := os.Create(newfilePath)

	check(err)

	defer outFile.Close()

	_, err = io.Copy(outFile, &buf)

	check(err)
}
