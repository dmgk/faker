package faker

import "time"

type Date struct{}

// Between returns random time in [from, to] interval, with second resolution.
func (d Date) Between(from, to time.Time) time.Time {
	if to.Sub(from) < 0 {
		panic("invalid time range")
	}
	return time.Unix(RandomInt64(from.Unix(), to.Unix()), 0)
}

// Forward returns random time in [time.Now(), time.Now() + duration] interval, with second resolution.
func (d Date) Forward(duration time.Duration) time.Time {
	if duration < 0 {
		panic("invalid duration")
	}
	now := time.Now()
	return d.Between(now, now.Add(duration))
}

// Backward returns random time in [time.Now() - duration, time.Now()] interval, with second resolution.
func (d Date) Backward(duration time.Duration) time.Time {
	if duration < 0 {
		panic("invalid duration")
	}
	now := time.Now()
	return d.Between(now.Add(-duration), now)
}

// Birthday returns random time so that age of the person born at that moment would be between minAge and maxAge years.
func (d Date) Birthday(minAge, maxAge int) time.Time {
	if minAge > maxAge {
		panic("invalid age range")
	}
	now := time.Now()
	from := now.AddDate(-maxAge, 0, 0)
	to := now.AddDate(-minAge, 0, 0)
	return d.Between(from, to)
}
