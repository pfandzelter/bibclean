package bbl

import (
	"regexp"
)

const (
	beginr = `\\bibitem(.*\%\n\s*)?{`
	endr   = `}`
)

// Parse parses a bbl file for used bibtex items.
func Parse(buf []byte) (map[string]struct{}, error) {

	r, err := regexp.Compile(beginr + `\S*}` + endr)

	if err != nil {
		return nil, err
	}

	found := r.FindAll(buf, -1)

	bblitems := make(map[string]struct{})

	beginr, err := regexp.Compile(beginr)

	if err != nil {
		return nil, err
	}

	endr, err := regexp.Compile(endr)
	if err != nil {
		return nil, err
	}

	for _, item := range found {

		item = beginr.ReplaceAll(item, []byte(""))
		item = endr.ReplaceAll(item, []byte(""))

		bblitems[string(item)] = struct{}{}
	}

	return bblitems, nil
}
