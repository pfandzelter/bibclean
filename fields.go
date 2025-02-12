package main

// Entry types

var fields = map[string]map[string][]string{
	"ieee": {
		"article": {
			"author",
			"title",
			"journal",
			"year",
			"volume",
			"number",
			"pages",
			"publisher",
		},

		"book": {
			"author",
			"editor",
			"title",
			"publisher",
			"year",
		},

		"incollection": {
			"author",
			"title",
			"booktitle",
			"publisher",
			"year",
		},

		"inproceedings": {
			"author",
			"title",
			"booktitle",
			"pages",
			"month",
			"year",
		},

		"mastersthesis": {
			"author",
			"title",
			"school",
			"month",
			"year",
		},

		"misc": {
			"author",
			"title",
			"howpublished",
			"month",
			"year",
			"note",
			"publisher",
		},

		"phdthesis": {
			"author",
			"title",
			"school",
			"month",
			"year",
		},

		"techreport": {
			"author",
			"title",
			"institution",
			"booktitle",
			"month",
			"year",
		},

		"unpublished": {
			"author",
			"title",
			"month",
			"year",
			"note",
		},
	},
	// https://www.acm.org/publications/authors/bibtex-formatting
	"acm": {
		"article": {
			"author",
			"title",
			"journal",
			"volume",
			"number",
			"month",
			"year",
			"issn",
			"pages",
			"articleno",
			"numpages",
			"url",
			"doi",
			"acmid",
			"publisher",
			"address",
			"issue_date",
			"eprint",
		},

		"book": {
			"author",
			"title",
			"year",
			"isbn",
			"publisher",
			"address",
			"editor",
		},

		"incollection": {
			"author",
			"title",
			"booktitle",
			"publisher",
			"pages",
			"year",
		},

		"inproceedings": {
			"author",
			"title",
			"booktitle",
			"pages",
			"month",
			"year",
			"acmid",
			"publisher",
			"address",
			"series",
			"location",
			"numpages",
			"url",
			"doi",
		},

		"mastersthesis": {
			"author",
			"title",
			"school",
			"month",
			"year",
		},

		"online": {
			"author",
			"organization",
			"title",
			"url",
			"month",
			"year",
			"lastaccessed",
		},

		"phdthesis": {
			"author",
			"title",
			"advisor",
			"school",
			"address",
			"month",
			"year",
			"url",
		},

		"techreport": {
			"author",
			"title",
			"institution",
			"address",
			"url",
			"number",
			"month",
			"year",
		},
	},
	// http://mirrors.ctan.org/macros/latex/contrib/biblatex/doc/biblatex.pdf
	"biblatex": {
		"article": {
			"author",
			"title",
			"journal",
			"volume",
			"number",
			"month",
			"year",
			"issn",
			"pages",
			"articleno",
			"numpages",
			"url",
			"doi",
			"publisher",
			"address",
			"issue_date",
		},

		"book": {
			"author",
			"title",
			"year",
			"isbn",
			"publisher",
			"address",
			"editor",
		},

		"incollection": {
			"author",
			"title",
			"booktitle",
			"publisher",
			"pages",
			"year",
		},

		"inproceedings": {
			"author",
			"title",
			"booktitle",
			"pages",
			"month",
			"year",
			"publisher",
			"address",
			"series",
			"venue",
			"numpages",
			"url",
			"doi",
			"pubstate",
		},

		"mastersthesis": {
			"author",
			"title",
			"institution",
			"address",
			"month",
			"year",
		},

		"online": {
			"author",
			"organization",
			"title",
			"url",
			"month",
			"year",
			"note",
			"urldate",
		},

		"phdthesis": {
			"author",
			"title",
			"advisor",
			"institution",
			"address",
			"month",
			"year",
		},

		"techreport": {
			"author",
			"title",
			"institution",
			"address",
			"url",
			"number",
			"month",
			"year",
		},

		"unpublished": {
			"author",
			"title",
			"month",
			"year",
			"eprint",
			"pubstate",
		},
	},
}
