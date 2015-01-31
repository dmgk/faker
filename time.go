package faker

import "time"

type Time struct {
	Date
}

// Between returns random time in [from, to] interval, with second resolution.
func (t Time) Between(from, to time.Time) time.Time {
	return t.Between(from, to)
}

// Forward returns random time in [time.Now(), time.Now() + duration] interval, with second resolution.
func (t Time) Forward(duration time.Duration) time.Time {
	return t.Forward(duration)
}

// Backward returns random time in [time.Now() - duration, time.Now()] interval, with second resolution.
func (t Time) Backward(duration time.Duration) time.Time {
	return t.Backward(duration)
}
