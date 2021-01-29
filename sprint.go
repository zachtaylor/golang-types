package types

import "fmt"

// Stringer is fmt.Stringer
type Stringer = fmt.Stringer

// Sprinter is a header for `Sprint(Any) string`
type Sprinter interface {
	Sprint(Any) string
}

type fmtSprint byte

func (fmtSprint) Sprint(any Any) string { return fmt.Sprint(any) }

// FmtSprinter is Sprinter using fmt.Sprint
const FmtSprinter = fmtSprint(0)

// SprinterFunc is a function type that implements Sprinter
type SprinterFunc func(Any) string

// Sprint calls the func
func (f SprinterFunc) Sprint(any Any) string { return f(any) }
