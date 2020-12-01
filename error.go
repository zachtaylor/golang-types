package types

// NewErr creates an error
func NewErr(string string) error {
	return Error{
		string: string,
	}
}

// WrapErr creates an error with a source
func WrapErr(error error, string string) error {
	return Error{
		error:  error,
		string: string,
	}
}

// Error is a basic error
type Error struct {
	error  error
	string string
}

func (e Error) Error() string {
	if e.error == nil {
		return e.string
	}
	return e.string + ": " + e.error.Error()
}

func (e Error) String() string { return "types.Error{" + e.Error() + "}" }
