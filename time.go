package types

import "time"

// Hour = time.Hour
const Hour = time.Hour

// Millisecond = time.Millisecond
const Millisecond = time.Millisecond

// Minute = time.Minute
const Minute = time.Minute

// Second = time.Second
const Second = time.Second

// Duration = time.Duration
type Duration = time.Duration

// Ticker = time.Ticker
type Ticker = time.Ticker

// Time = time.Time
type Time = time.Time

// Timer = time.Timer
type Timer = time.Timer

// Clocker is a clock interface
type Clocker interface {
	// NewTime returns the time
	NewTime() Time
}

// ClockerFunc is a func type for Clocker
type ClockerFunc func() Time

// NewTime returns f()
func (f ClockerFunc) NewTime() Time { return f() }

// DefaultClock creates a Clocker from NewTime
func DefaultClock() Clocker { return ClockerFunc(NewTime) }

// Date returns time.Date()
func Date(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) Time {
	return time.Date(year, month, day, hour, min, sec, nsec, loc)
}

// NewTicker creates a time.Ticker
func NewTicker(d Duration) *Ticker { return time.NewTicker(d) }

// NewTimer creates a time.Timer
func NewTimer(d Duration) *Timer { return time.NewTimer(d) }

// NewTime returns time.Now()
func NewTime() Time { return time.Now() }

// NewTimeUnix return time.Unix()
func NewTimeUnix(sec int64, nsec int64) Time { return time.Unix(sec, nsec) }
