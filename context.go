package types

import "context"

// Context is simple context.Context wrapper
type Context struct{ context.Context }

// NewContext is context.Background
func NewContext() Context { return Context{context.Background()} }

// Deadline is Context.Deadline
func (c Context) Deadline() (deadline Time, ok bool) { return c.Context.Deadline() }

// Done is Context.Done
func (c Context) Done() <-chan struct{} { return c.Context.Done() }

// Err is Context.Err
func (c Context) Err() error { return c.Context.Err() }

// Value is Context.Value
func (c Context) Value(key Any) Any { return c.Value(key) }

// WithCancel is context.WithCancel(this)
func (c Context) WithCancel() (ctx Context, cancel context.CancelFunc) {
	context, cancel := context.WithCancel(c)
	return Context{context}, cancel
}

// WithDeadline is context.WithDeadline(this)
func (c Context) WithDeadline(d Time) (Context, context.CancelFunc) {
	context, cancel := context.WithDeadline(c, d)
	return Context{context}, cancel
}

// WithTimeout is context.WithTimeout(this)
func (c Context) WithTimeout(d Duration) (Context, context.CancelFunc) {
	context, cancel := context.WithTimeout(c, d)
	return Context{context}, cancel
}

// WithValue is context.WithValue(this)
func (c Context) WithValue(k, v Any) Context { return Context{context.WithValue(c, k, v)} }
