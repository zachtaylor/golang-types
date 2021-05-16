package types

// NewErr creates an error with call stack
func NewErr(string string) Error {
	return Error{
		source: NewSource(1),
		string: string,
	}
}

// TraceErr wraps an error with call stack
func TraceErr(error error) Error {
	return Error{
		error:  error,
		source: NewSource(1),
	}
}

// WrapErr wraps an error with call stack and comment
func WrapErr(error error, string string) Error {
	return Error{
		error:  error,
		source: NewSource(1),
		string: string,
	}
}

// Error is a custom error
type Error struct {
	error  error
	source Source
	string string
}

func (e Error) Unwrap() error { return e.error }

// Source returns the source of the error if available
func (e Error) Source() Source { return e.source }

func (e Error) Error() string {
	if e.error == nil {
		return e.string
	} else if len(e.string) < 1 {
		return e.error.Error()
	}
	return e.string + ": " + e.error.Error()
}

func (e Error) String() string { return "types.Error{" + e.Error() + "}" }

// ErrorUnwrap returns the cause of the error, if available
func ErrorUnwrap(err error) (error error) {
	if wrapper, _ := err.(ErrorWrapper); wrapper != nil {
		error = wrapper.Unwrap()
	}
	return
}

// ErrorWrapper is an error sub-type which returns wrapped error
type ErrorWrapper interface {
	// Unwrap returns the cause of the given error
	Unwrap() error
}

// ErrorSource returns the source of the error, if available
func ErrorSource(err error) (source Source) {
	if sourcer, _ := err.(ErrorSourcer); sourcer != nil {
		source = sourcer.Source()
	}
	return
}

// ErrorSourcer is an error sub-type which returns error source
type ErrorSourcer interface {
	// Source returns the source of the error
	Source() Source
}

// ErrorChain returns the stack of errors extracted by ErrorCause
func ErrorChain(err error) (causes []error) {
	for ; err != nil; err = ErrorUnwrap(err) {
		causes = append(causes, err)
	}
	return
}
