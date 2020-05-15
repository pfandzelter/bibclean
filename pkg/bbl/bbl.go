package bbl

import (
	"regexp"
)

// Parse parses a bbl file for used bibtex items.
func Parse(buf []byte) (map[string]struct{}, error) {
	r, err := regexp.Compile(`\\bibitem{\S*}`)

	if err != nil {
		return nil, err
	}

	found := r.FindAll(buf, -1)

	bblitems := make(map[string]struct{})

	for _, item := range found {

		bblitems[string(item[len(`\bibtex{`)+1:len(item)-len("}")])] = struct{}{}
	}

	return bblitems, nil
}
