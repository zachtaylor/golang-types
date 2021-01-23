package types

import "strings"

// StringBuilder is strings.Builder
type StringBuilder = strings.Builder

// StringGenerator creates a new string
type StringGenerator interface {
	New() string
}

// StringReader is strings.Reader
type StringReader = strings.Reader

// StringReplacer is strings.Replacer
type StringReplacer = strings.Replacer

// NewStringReader returns strings.NewReader
func NewStringReader(s string) *StringReader { return strings.NewReader(s) }

// NewStringReplacer returns strings.NewReplacer
func NewStringReplacer(oldnew ...string) *StringReplacer { return strings.NewReplacer(oldnew...) }
