package types

import "time"

// UnixTimestamp is a custom type for Unix timestamps.
type UnixTimestamp int64

// Time converts UnixTimestamp to time.Time.
func (ut UnixTimestamp) Time() time.Time {
	return time.Unix(int64(ut), 0)
}

// String formats UnixTimestamp as a human-readable string.
func (ut UnixTimestamp) String() string {
	return ut.Time().Format(time.RFC3339)
}

// NewUnixTimestamp creates a new UnixTimestamp from time.Time.
func NewUnixTimestamp(t time.Time) UnixTimestamp {
	return UnixTimestamp(t.Unix())
}
