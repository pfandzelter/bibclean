package main

import (
	"bytes"
	"cmp"
	"errors"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"

	"github.com/pfandzelter/bibclean/pkg/bbl"
	"github.com/pfandzelter/bibclean/pkg/bibtex"
	"github.com/pfandzelter/bibclean/pkg/merge"
)

// update this version when making changes by tagging the commit
// compile with make to get all this information automatically
// OR go build -ldflags "-X main.version=$(git describe --tags --always) -X main.commit=$(shell git rev-parse HEAD) -X main.date=$(shell date -u '+%Y-%m-%d_%I:%M:%S%p') -X main.builtBy=goreleaser".
// OR goreleaser will do this automatically

var version = "unknown"
var commit = "unknown"
var date = "unknown"

const terminalWidth = 80

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

func fmtBreak(s string, n int) string {
	if len(s) >= n {
		return s
	}

	return fmt.Sprintf(
		"%s %s %s",
		strings.Repeat("-", int(math.Ceil(float64(n-len(s))/2))-1),
		s,
		strings.Repeat("-", int(n-len(s))/2-1),
	)
}

func main() {

	var printVersion, noMerge *bool
	var bibfile, newfile, bblfile, shorten *string
	var defaults *string
	var shortenBooktitle, shortenAll bool
	var additional additionalFields = make(additionalFields)

	printVersion = flag.Bool("version", false, "print bibclean version and exit")
	bibfile = flag.String("in", "", "input bibliography file")
	newfile = flag.String("out", "", "output bibliography file")
	bblfile = flag.String("bbl", "", "(optional) auxillary .bbl file to check which references have been used in the text")
	defaults = flag.String("defaults", "ieee", "(optional) default data fields, can be \"ieee\" (for IEEEtran.bst), \"acm\" (for ACM-Reference-Format.bst), or \"biblatex\" (for biblatex)")
	shorten = flag.String("shorten", "", "(optional) level of applied title shortening to conform with IEEE citation style, can be \"publication\" (shorten only proceeding and journal titles with some common abbreviations) or \"all\" (aggressive shortening including shortening titles, uses the full list of abbrevations)")
	noMerge = flag.Bool("no-merge", false, "(optional) disable merging repeated entries based on key. redundant values will be added as comments")
	flag.Var(&additional, "additional", "Additional fields for entries: specify as many as you like in the form \"--additional=article:booktitle --additional=techreport:address\" (this will add a \"booktitle\" field to \"@article\" entries and an \"address\" field to \"@techreport\" entries)")

	flag.Parse()

	if *printVersion {
		fmt.Printf("bibclean\nversion %s\nbuilt %s\ncommit %s\n", version, date, commit)
		os.Exit(0)
	}

	incorrectUse := (*bibfile == "") || (*newfile == "")

	switch *shorten {
	case "", "none":
	case "all":
		shortenAll = true
		fallthrough
	case "publication":
		shortenBooktitle = true
	default:
		incorrectUse = true
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

	var e map[string][]string
	switch strings.ToLower(*defaults) {
	case "ieee":
		e = fields["ieee"]
	case "acm":
		e = fields["acm"]
	case "biblatex":
		e = fields["biblatex"]
	default:
		fmt.Printf("unknown default type: %s\n", *defaults)
		os.Exit(1)
	}

	types := make([]string, len(e))

	i := 0
	for t := range e {
		types[i] = t
		i++
	}

	plugins := []func(e bibtex.Element) bibtex.Element{
		bibtex.CleanQuotationMarks,
		bibtex.AddProcOf,
		bibtex.CleanCurly,
		bibtex.CleanDOI,
		bibtex.CleanPages,
		bibtex.AddPublisherAddress,
	}

	if shortenBooktitle {
		plugins = append(plugins, bibtex.ShortenBooktitle)
	}

	if shortenAll {
		plugins = append(plugins, bibtex.ShortenAll)
	}

	elements, err := bibtex.Parse(contents, &e, additional, plugins)

	check(err)

	if !*noMerge {
		elements, err = merge.MergeElements(elements)

		check(err)
	}

	// sort types alphabetically
	sort.Strings(types)

	// sort elements by ID
	elemUsed := make(map[string][]*bibtex.Element)
	elemDefault := make(map[string][]*bibtex.Element)

	for _, element := range elements {
		if _, ok := used[element.ID]; ok && usebbl {
			elemUsed[element.Type] = append(elemUsed[element.Type], element)
			continue
		}
		elemDefault[element.Type] = append(elemDefault[element.Type], element)
	}

	for _, t := range types {
		slices.SortFunc(elemUsed[t], func(a, b *bibtex.Element) int {
			return cmp.Compare(strings.ToLower(a.ID), strings.ToLower(b.ID))
		})

		slices.SortFunc(elemDefault[t], func(a, b *bibtex.Element) int {
			return cmp.Compare(strings.ToLower(a.ID), strings.ToLower(b.ID))
		})
	}

	if usebbl {
		fmt.Fprintf(&buf, "%% %s\n", strings.Repeat("-", terminalWidth-2))
		fmt.Fprintf(&buf, "%% %s\n", fmtBreak("USED ENTRIES", terminalWidth-2))
		fmt.Fprintf(&buf, "%% %s\n\n", strings.Repeat("-", terminalWidth-2))
		for _, t := range types {

			fmt.Fprintf(&buf, "%% %s\n\n", fmtBreak(strings.ToUpper(t), terminalWidth-2))

			for _, element := range elemUsed[t] {
				fmt.Fprintf(&buf, "%s\n\n", element)
			}
		}

		fmt.Fprintf(&buf, "%% %s\n", strings.Repeat("-", terminalWidth-2))
		fmt.Fprintf(&buf, "%% %s\n", fmtBreak("UNUSED ENTRIES", terminalWidth-2))
		fmt.Fprintf(&buf, "%% %s\n\n", strings.Repeat("-", terminalWidth-2))
	}

	for _, t := range types {

		fmt.Fprintf(&buf, "%% %s\n\n", fmtBreak(strings.ToUpper(t), terminalWidth-2))

		for _, element := range elemDefault[t] {
			fmt.Fprintf(&buf, "%s\n\n", element)
		}
	}

	outFile, err := os.Create(newfilePath)

	check(err)

	defer outFile.Close()

	_, err = io.Copy(outFile, &buf)

	check(err)
}
