// Package bibtex is a quick and dirty BibTeX parser for working with
// a Bibtex citation
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
//
// Copyright (c) 2016, Caltech
// All rights not granted herein are expressly reserved by Caltech.
//
// Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
package bibtex

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"regexp"
	"strings"

	// Caltech Library packages
	"github.com/caltechlibrary/bibtex/tok"
)

const (
	// Version of BibTeX package
	Version = `v0.0.8`

	// LicenseText holds the text for displaying license info
	LicenseText = `
%s %s

Copyright (c) 2016, Caltech
All rights not granted herein are expressly reserved by Caltech.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
`

	// DefaultInclude list
	DefaultInclude = "comment,string,article,book,booklet,inbook,incollection,inproceedings,conference,manual,mastersthesis,misc,phdthesis,proceedings,techreport,unpublished"

	// A template for printing an element
	ElementTmplSrc = `
@{{- .Type -}}{
    {{-range .Keys}}
	{{ . -}},
	{{end}}
	{{-range $key, $val := .Tags}}
		{{- $key -}} = {{- $val -}},
	{{end}}
}
`
)

// Generic Element
type Element struct {
	XMLName      xml.Name          `json:"-"`
	ID           string            `xml:"id" json:"id"`
	Key          string            `xml:"key" json:"key"`
	Type         string            `xml:"type" json:"type"`
	Tags         map[string]string `xml:"tags" json:"tags"`
	RequiredKeys *TagTypes
}

type Elements []*Element

type TagTypes struct {
	Required []string
}

// String renders a single BibTeX element
func (element *Element) String() string {
	var out []string

	keys := element.RequiredKeys.Required

	neededKeys := make(map[string]struct{})

	for _, key := range keys {
		neededKeys[key] = struct{}{}
	}

	for key, tag := range element.Tags {
		// add the keys that we don't need as comments
		if _, ok := neededKeys[key]; !ok {
			out = append(out, fmt.Sprintf("%% %s: %s", key, tag))
		}
	}

	if len(element.ID) > 0 {
		out = append(out, fmt.Sprintf("@%s{%s,", element.Type, element.ID))
	} else {
		out = append(out, fmt.Sprintf("@%s{", element.Type))
	}

	for _, ky := range keys {
		val := element.Tags[ky]
		if len(val) != 0 {
			val := regexp.MustCompile(`\s+`).ReplaceAllString(val, " ")
			out = append(out, fmt.Sprintf("    %s = %s,", ky, val))
		} else {
			out = append(out, fmt.Sprintf("    %s = MISSING,", ky))
		}
	}

	// remove trailing comma
	if last := len(out) - 1; last >= 0 {
		if char := len(out[last]) - 1; char >= 0 && out[last][char] == ',' {
			out[last] = out[last][:char]
		}
	}

	out = append(out, "}")
	return strings.Join(out, "\n")
}

//
// Parser related structures
//

// Bib is a naive BibTeX Tokenizer function
// Note: there is an English bias in the AlphaNumeric check
func Bib(token *tok.Token, buf []byte) (*tok.Token, []byte) {
	switch {
	case token.Type == tok.AtSign || token.Type == "BibElement":
		// Get the next Token
		newTok, newBuf := tok.Tok(buf)
		if newTok.Type != tok.OpenCurlyBracket {
			token.Type = "BibElement"
			token.Value = append(token.Value[:], newTok.Value[:]...)
			token, buf = Bib(token, newBuf)
		}
	case token.Type == tok.Space:
		newTok, newBuf := tok.Tok(buf)
		if newTok.Type == tok.Space {
			token.Value = append(token.Value[:], newTok.Value[:]...)
			token, buf = Bib(token, newBuf)
		}
	case token.Type == tok.Letter || token.Type == tok.Numeral || token.Type == "AlphaNumeric":
		// Convert Letters and Numerals to AlphaNumeric Type.
		token.Type = "AlphaNumeric"
		// Get the next Token
		newTok, newBuf := tok.Tok(buf)
		if newTok.Type == tok.Letter || newTok.Type == tok.Numeral {
			token.Value = append(token.Value[:], newTok.Value[:]...)
			token, buf = Bib(token, newBuf)
		}
	default:
		// Revaluate token for more specific token types.
		token = tok.TokenFromMap(token, map[string][]byte{
			tok.OpenCurlyBracket:  tok.OpenCurlyBrackets,
			tok.CloseCurlyBracket: tok.CloseCurlyBrackets,
			tok.AtSign:            tok.AtSignMark,
			tok.EqualSign:         tok.EqualMark,
			tok.DoubleQuote:       tok.DoubleQuoteMark,
			tok.SingleQuote:       tok.SingleQuoteMark,
			"Comma":               []byte(","),
		})
	}

	return token, buf
}

func mkElement(elementType string, defaultElements *TagTypes, additionalFields map[string]struct{}, buf []byte) (*Element, error) {
	var (
		key     []byte
		val     []byte
		between []byte
		token   *tok.Token
		err     error
		tags    map[string]string
		id      string
	)

	element := new(Element)
	element.Type = elementType

	element.RequiredKeys = &TagTypes{
		Required: make([]string, len(defaultElements.Required), len(defaultElements.Required)+len(additionalFields)),
	}

	for i, field := range defaultElements.Required {
		element.RequiredKeys.Required[i] = field
	}

	for f := range additionalFields {
		element.RequiredKeys.Required = append(element.RequiredKeys.Required, f)
	}

	tags = make(map[string]string)

	for {
		if len(buf) == 0 {
			if len(key) > 0 {
				// We have a trailing key/value pair to save.
				tags[strings.ToLower(string(key))] = string(val)
			}
			break
		}
		_, token, buf = tok.Skip2(tok.Space, buf, Bib)
		switch {
		case token.Type == tok.OpenCurlyBracket:
			buf = tok.Backup(token, buf)
			between, buf, err = tok.Between([]byte("{"), []byte("}"), []byte(""), buf)
			if err != nil {
				return element, err
			}
			// Non-destructively copy the quote into val
			val = append(val, []byte("{")[0])
			val = append(val[:], between[:]...)
			val = append(val, []byte("}")[0])
		case token.Type == tok.DoubleQuote:
			buf = tok.Backup(token, buf)
			between, buf, err = tok.Between([]byte("\""), []byte("\""), []byte("\\"), buf)
			if err != nil {
				return element, err
			}
			// Non-destructively copy the quote into val
			val = append(val, []byte("\"")[0])
			val = append(val[:], between[:]...)
			val = append(val, []byte("\"")[0])
		case token.Type == tok.EqualSign:
			key = val
			val = nil
		case token.Type == "Comma" || len(buf) == 0:
			if len(key) > 0 {
				//make a map entry
				tags[strings.ToLower(string(key))] = string(val)
			} else if len(val) > 0 && len(id) == 0 {
				// this is our element id^
				id = string(val)
			}
			key = nil
			val = nil
		case token.Type == tok.Punctuation && bytes.Equal(token.Value, []byte("#")):
			val = append(val[:], []byte(" # ")[:]...)
		default:
			val = append(val[:], token.Value[:]...)
		}
	}
	if len(tags) > 0 {
		element.Tags = tags
	}

	element.ID = id

	return element, nil
}

// Parse a BibTeX file into appropriate structures
func Parse(buf []byte, defaultElements *map[string][]string, additionalFields map[string]map[string]struct{}, plugins []func(Element) Element) ([]*Element, error) {

	var (
		lineNo      int
		token       *tok.Token
		elements    []*Element
		err         error
		skipped     []byte
		entrySource []byte
		LF          = []byte("\n")
	)

	// convert the default elements map to a map of TagTypes
	defaultFields := make(map[string]*TagTypes)
	for elementType, fields := range *defaultElements {
		defaultFields[elementType] = &TagTypes{Required: fields}
	}

	lineNo = 1
	for {
		if len(buf) == 0 {
			break
		}
		skipped, token, buf = tok.Skip2(tok.Space, buf, Bib)
		lineNo = lineNo + bytes.Count(skipped, LF)
		if token.Type == tok.AtSign {
			// We may have a entry key
			token, buf = tok.Tok2(buf, Bib)
			if token.Type == "AlphaNumeric" {
				elementType := token.Value[:]
				skipped, token, buf = tok.Skip2(tok.Space, buf, Bib)
				lineNo = lineNo + bytes.Count(skipped, LF)
				if token.Type == tok.OpenCurlyBracket {
					// Ok it looks like we have a Bib entry now.
					buf = tok.Backup(token, buf)
					entrySource, buf, err = tok.Between([]byte("{"), []byte("}"), []byte(""), buf)
					if err != nil {
						return elements, fmt.Errorf("problem parsing entry at %d", lineNo)
					}
					// OK, we have an entry, let's process it.
					et := strings.ToLower(string(elementType))

					if _, ok := defaultFields[et]; !ok {
						return nil, fmt.Errorf("element type %s is unknown (line %d)", et, lineNo)
					}

					element, err := mkElement(et, defaultFields[et], additionalFields[et], entrySource)
					if err != nil {
						return elements, fmt.Errorf("error parsing element at l. %d, %s", lineNo, err)
					}
					lineNo = lineNo + bytes.Count(entrySource, LF)
					// OK, we have an element, let's append to our array...

					elements = append(elements, element)
				}
			}
		}
	}
	if len(elements) == 0 {
		return nil, fmt.Errorf("no elements found")
	}

	// run plugins
	for _, plugin := range plugins {
		for _, element := range elements {
			*element = plugin(*element)
		}
	}

	return elements, nil
}
