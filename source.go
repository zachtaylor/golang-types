package types

import (
	"runtime"
	"strconv"
)

// Source links to a go code instruction
type Source struct {
	file string
	line int
}

// NewSource creates a Source reference to the caller number # higher in the stack frame than the invoking call (0 is default)
func NewSource(history int) Source {
	_, file, line, _ := runtime.Caller(1 + history)

	return Source{
		file: file,
		line: line,
	}
}

// File returns go file reference
func (src Source) File() string {
	return src.file
}

// Line returns source line
func (src Source) Line() int {
	return src.line
}

func (src *Source) String() string {
	msg := ""
	if src != nil {
		if file := src.File(); file != "" {
			if flen := len(file); flen > 4 && file[flen-3:] == ".go" {
				msg = file[:flen-3]
			} else {
				msg = file
			}
		}
		if lno := src.Line(); lno > 0 {
			msg += "#" + strconv.FormatInt(int64(lno), 10)
		}
	}
	return msg
}
