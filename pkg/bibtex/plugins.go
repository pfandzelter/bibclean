// this file provides small functions that can be plugged in to
// bibtex parsing and cleaning. every function receives a map of keys
// and values and can set new values (keys should be changed). this way, we can
// easily add new minor cleaning functionality to bibclean.
package bibtex

import (
	"regexp"
	"strings"
)

// CleanPages removes single dashes from page numbers and replaces
// them with an em-dash ("--").
func CleanPages(e Element) Element {
	for key, val := range e.Tags {
		if key != "pages" {
			continue
		}

		// check if there is something that has two numbers seperated by something
		r := regexp.MustCompile(`"\d+[^\d]+\d+"`)
		if !r.MatchString(val) {
			// if not, it's probably something weird like Elsevier
			continue
		}

		// replace the non-numbers with an em-dash
		r = regexp.MustCompile(`[^\d"]+`)
		e.Tags[key] = r.ReplaceAllString(val, "--")
	}

	return e
}

// CleanCurly removes the useless escaped curly braces from USENIX
// conference names that Google Scholer adds.
func CleanCurly(e Element) Element {
	for key, val := range e.Tags {
		if key != "booktitle" {
			continue
		}

		// replace "$\{$" with "{"
		val = strings.Replace(val, "$\\{$", "{", -1)
		// replace "$\}$" with "}"
		val = strings.Replace(val, "$\\}$", "}", -1)

		e.Tags[key] = val
	}

	return e
}

// CleanQuotationMarks makes sure that only double quotation marks are
// used in the start and end of the value. Unless the value is only a number.
// Or a date. Mostly, this is about removing curly braces.
// Also, we need to change umlaut escapes: \"{a} does not work, it should be {\"a} so the quotes are in braces.
func CleanQuotationMarks(e Element) Element {
	for key, val := range e.Tags {
		if key == "month" {
			r := regexp.MustCompile(`[^a-z]`)
			e.Tags[key] = r.ReplaceAllString(val, "")
			continue
		}

		// this is probably a bad idea
		// in case the number of volume does have a letter in it for some reason
		if key == "year" || key == "volume" {
			r := regexp.MustCompile(`[^\d]`)
			e.Tags[key] = r.ReplaceAllString(val, "")
			continue
		}

		// remove curly braces
		// only at the start and end
		r := regexp.MustCompile(`^{|}$`)
		val = r.ReplaceAllString(val, "\"")

		// replace umlaut escapes
		r = regexp.MustCompile(`\\"\{([a-zA-Z])\}`)
		// now use a capture group to get the letter and put it in braces
		val = r.ReplaceAllString(val, `{\"$1}`)

		e.Tags[key] = val
	}

	return e
}

// AddProcOf adds "Proceedings of the" to the start of the booktitle if it is missing.
func AddProcOf(e Element) Element {
	// only for inproceedings
	if e.Type != "inproceedings" {
		return e
	}

	for key, val := range e.Tags {
		if key != "booktitle" {
			continue
		}

		// check if the booktitle is empty or MISSING
		if val == "" || val == "MISSING" {
			break
		}

		// check if the booktitle starts with "Proceedings of the"
		r := regexp.MustCompile(`^"Proceedings of the`)
		if r.MatchString(val) {
			continue
		}

		e.Tags[key] = "\"Proceedings of the " + val[1:]
		break
	}

	return e
}

// CleanDOI checks doi and url fields and sets the other one if it is missing.
func CleanDOI(e Element) Element {
	// check if doi is missing
	if doi, ok := e.Tags["doi"]; !ok || doi == "" {
		// check if url is a doi
		if _, ok = e.Tags["url"]; !ok {
			// no doi, no url
			return e
		}

		r := regexp.MustCompile(`^"https?://(dx\.)?doi\.org/`)
		if r.MatchString(e.Tags["url"]) {
			// remove the doi.org prefix
			e.Tags["doi"] = "\"" + r.ReplaceAllString(e.Tags["url"], "")
		}
		return e
	}

	// check if url is missing
	if url, ok := e.Tags["url"]; !ok || url == "" {
		// check if doi is a doi!
		if _, ok = e.Tags["doi"]; !ok {
			// no doi, no url
			return e
		}

		// assuming dois are characterised by having a slash in them
		r := regexp.MustCompile(`^.*/.*$`)
		if r.MatchString(e.Tags["doi"]) {
			e.Tags["url"] = "\"https://doi.org/" + e.Tags["doi"][1:]
		}
	}

	return e
}

// AddPublisherLocation adds the location of the publisher to the
// entry. They're all based in New York for some reason.
func AddPublisherAddress(e Element) Element {
	// check if address exists but is empty
	addresses := map[string]string{
		"ACM":                                 "New York, NY, USA",
		"Association for Computing Machinery": "New York, NY, USA",
		"IEEE":                                "New York, NY, USA",
		"USENIX":                              "Berkeley, CA, USA",
		// Springer is difficult, they have multiple locations
		// some are Berlin, Heidelberg, or Cham
		//	"Springer":  "Berlin, Germany",
		"Elsevier": "Amsterdam, The Netherlands",
	}

	// if address is not a required field, we can't do anything
	addressRequired := false
	for _, req := range e.RequiredKeys.Required {
		if req == "address" {
			addressRequired = true
			break
		}
	}

	if !addressRequired {
		return e
	}

	// if there is already an address, better not touch it

	if _, ok := e.Tags["address"]; !ok && e.Tags["address"] != "" {
		return e
	}

	// if there is no publisher, we can't do anything
	if _, ok := e.Tags["publisher"]; !ok {
		return e
	}

	// check if the publisher is in the list
	if addr, ok := addresses[e.Tags["publisher"][1:len(e.Tags["publisher"])-1]]; ok {
		e.Tags["address"] = "\"" + addr + "\""
	}

	return e
}

// ShortenBooktitle replaces long conference names with approved short forms from IEEE.
func ShortenBooktitle(e Element) Element {
	for tag := range e.Tags {
		if tag == "booktitle" || tag == "journal" {
			for old, new := range *ieeeTitleShortforms {
				//log.Printf("replacing %s with %s in %s", old, new, element.Tags[tag])
				e.Tags[tag] = strings.Replace(e.Tags[tag], old, new, -1)
				e.Tags[tag] = strings.Replace(e.Tags[tag], strings.ToLower(old), strings.ToLower(new), -1)
			}
		}
	}

	return e
}

// ShortenAuthors truncates the author list for more than three authors.
func ShortenAuthors(e Element) Element {
	for key, val := range e.Tags {
		if key != "author" {
			continue
		}

		// check if there are more than three authors
		r := regexp.MustCompile(`\sand.*and`)
		if !r.MatchString(val) {
			continue
		}

		// replace everything after the first author with "et al."
		r = regexp.MustCompile(`\sand.*$`)
		e.Tags[key] = r.ReplaceAllString(val, " and others\"")
	}

	return e
}

// ShortenAll replaces long words with approved short forms from IEEE.
func ShortenAll(e Element) Element {
	for tag := range e.Tags {
		if tag == "title" || tag == "booktitle" || tag == "journal" {
			for old, new := range *ieeeShortforms {
				//log.Printf("replacing %s with %s in %s", old, new, element.Tags[tag])
				e.Tags[tag] = strings.Replace(e.Tags[tag], old, new, -1)
				e.Tags[tag] = strings.Replace(e.Tags[tag], strings.ToLower(old), strings.ToLower(new), -1)
			}
		}
	}

	return ShortenAuthors(e)
}
