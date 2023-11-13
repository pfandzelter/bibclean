package merge

import (
	"github.com/pfandzelter/bibclean/pkg/bibtex"
)

const COPY_POSTFIX = " (duplicate)"

func MergeElements(orig []*bibtex.Element) ([]*bibtex.Element, error) {
	// iterate through the elements
	m := make(map[string]*bibtex.Element)

	for _, elem := range orig {
		_, ok := m[elem.ID]
		if !ok {
			// new element! just add it to the map
			m[elem.ID] = elem
			continue
		}

		// uh-oh, a repeated element!
		// merge them tag by tag
		for tag, value := range elem.Tags {
			// check if this tag exists in what we have
			v, ok := m[elem.ID].Tags[tag]

			if !ok || v == "" {
				// nope!
				m[elem.ID].Tags[tag] = value
				continue

			}
			// add as a copy if not in there already

			copyTag := tag + COPY_POSTFIX
			v, ok = m[elem.ID].Tags[copyTag]

			if !ok || v == "" {
				m[elem.ID].Tags[copyTag] = value
				continue
			}

			m[elem.ID].Tags[copyTag] = m[elem.ID].Tags[copyTag] + " - " + value

		}
	}

	// convert the map back into a list
	l := make([]*bibtex.Element, 0, len(m))

	for _, elem := range m {
		l = append(l, elem)
	}

	return l, nil
}
