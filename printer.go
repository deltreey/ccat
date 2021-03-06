package main

import (
	"io"
	"text/scanner"

	"github.com/jingweno/ccat/Godeps/_workspace/src/github.com/sourcegraph/syntaxhighlight"
)

type ColorDefs map[syntaxhighlight.Kind]string

var LightColorDefs = ColorDefs{
	syntaxhighlight.String:        "brown",
	syntaxhighlight.Keyword:       "darkblue",
	syntaxhighlight.Comment:       "lightgrey",
	syntaxhighlight.Type:          "teal",
	syntaxhighlight.Literal:       "teal",
	syntaxhighlight.Punctuation:   "darkred",
	syntaxhighlight.Plaintext:     "darkblue",
	syntaxhighlight.Tag:           "blue",
	syntaxhighlight.HTMLTag:       "lightgreen",
	syntaxhighlight.HTMLAttrName:  "blue",
	syntaxhighlight.HTMLAttrValue: "green",
	syntaxhighlight.Decimal:       "darkblue",
}

var DarkColorDefs = ColorDefs{
	syntaxhighlight.String:        "brown",
	syntaxhighlight.Keyword:       "blue",
	syntaxhighlight.Comment:       "darkgrey",
	syntaxhighlight.Type:          "turquoise",
	syntaxhighlight.Literal:       "turquoise",
	syntaxhighlight.Punctuation:   "red",
	syntaxhighlight.Plaintext:     "blue",
	syntaxhighlight.Tag:           "blue",
	syntaxhighlight.HTMLTag:       "lightgreen",
	syntaxhighlight.HTMLAttrName:  "blue",
	syntaxhighlight.HTMLAttrValue: "green",
	syntaxhighlight.Decimal:       "blue",
}

func AsCCat(r io.Reader, w io.Writer, cdefs ColorDefs) error {
	return syntaxhighlight.Print(newScanner(r), w, Printer{cdefs})
}

func newScanner(r io.Reader) *scanner.Scanner {
	var s scanner.Scanner
	s.Init(r)
	s.Error = func(_ *scanner.Scanner, _ string) {}
	s.Whitespace = 0
	s.Mode = s.Mode ^ scanner.SkipComments

	return &s
}

type Printer struct {
	ColorDefs ColorDefs
}

func (p Printer) Print(w io.Writer, kind syntaxhighlight.Kind, tokText string) error {
	c := p.ColorDefs[kind]
	if c != "" {
		tokText = Colorize(c, tokText)
	}

	_, err := io.WriteString(w, tokText)

	return err
}
