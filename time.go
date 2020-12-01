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

// NewTicker creates a time.Ticker
func NewTicker(d Duration) *Ticker {
	return time.NewTicker(d)
}

// NewTimer creates a time.Timer
func NewTimer(d Duration) *Timer {
	return time.NewTimer(d)
}

// NewTime returns time.Now()
func NewTime() Time {
	return time.Now()
}

// NewTimeUnix return time.Unix()
func NewTimeUnix(sec int64, nsec int64) Time {
	return time.Unix(sec, nsec)
}
