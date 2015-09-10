package faker

import (
	"testing"
	"time"
)

func TestDateBetween(t *testing.T) {
	from := time.Now().AddDate(0, -1, 0)
	examples := []struct {
		from time.Time
		to   time.Time
	}{
		{from, from},
		{from, from.Add(time.Microsecond)},
		{from, from.Add(time.Millisecond)},
		{from, from.Add(time.Second)},
		{from, from.Add(time.Minute)},
		{from, from.Add(time.Hour)},
		{from, from.AddDate(0, 0, 1)},
		{from, from.AddDate(0, 1, 0)},
		{from, from.AddDate(1, 0, 0)},
		{from, from.AddDate(10, 0, 0)},
	}

	for _, x := range examples {
		d := Date().Between(x.from, x.to)
		if d.Unix() < x.from.Unix() || d.Unix() > x.to.Unix() {
			t.Errorf("expected %v to be in range [%v, %v]", d, x.from, x.to)
		}
	}
}

func TestDateBetweenInvalidRange(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			// ignore
		}
	}()
	now := time.Now()

	Date().Between(now.Add(time.Hour), now)
	t.Errorf("expected invalid time range panic")
}

func TestDateForward(t *testing.T) {
	now := time.Now()
	examples := []time.Duration{
		0,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		time.Minute,
		time.Hour,
		now.AddDate(0, 0, 1).Sub(now),
		now.AddDate(0, 1, 0).Sub(now),
		now.AddDate(1, 0, 0).Sub(now),
		now.AddDate(10, 0, 0).Sub(now),
	}

	for _, x := range examples {
		d := Date().Forward(x)
		from := now.Unix()
		to := now.Add(x).Unix()
		if d.Unix() < from || d.Unix() > to {
			t.Errorf("expected %v to be in range [%v, %v]", d, time.Unix(from, 0), time.Unix(to, 0))
		}
	}
}

func TestDateForwardInvalidDuration(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			// ignore
		}
	}()

	Date().Forward(-1 * time.Hour)
	t.Errorf("expected invalid duration panic")
}

func TestDateBackward(t *testing.T) {
	now := time.Now()
	examples := []time.Duration{
		0,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		time.Minute,
		time.Hour,
		now.AddDate(0, 0, 1).Sub(now),
		now.AddDate(0, 1, 0).Sub(now),
		now.AddDate(1, 0, 0).Sub(now),
		now.AddDate(10, 0, 0).Sub(now),
	}

	for _, x := range examples {
		d := Date().Backward(x)
		from := now.Add(-x).Unix()
		to := now.Unix()
		if d.Unix() < from || d.Unix() > to {
			t.Errorf("expected %v to be in range [%v, %v]", d, time.Unix(from, 0), time.Unix(to, 0))
		}
	}
}

func TestDateBackwardInvalidDuration(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			// ignore
		}
	}()

	Date().Backward(-1 * time.Second)
	t.Errorf("expected invalid duration panic")
}

func TestBirthday(t *testing.T) {
	now := time.Now()
	examples := []struct {
		minAge int
		maxAge int
	}{
		{42, 42},
		{20, 21},
		{22, 32},
		{21, 65},
		{1, 99},
	}

	for _, x := range examples {
		d := Date().Birthday(x.minAge, x.maxAge)
		from := now.AddDate(-x.maxAge, 0, 0).Unix()
		to := now.AddDate(-x.minAge, 0, 0).Unix()
		if d.Unix() < from || d.Unix() > to {
			t.Errorf("expected %v to be in range [%v, %v]", d, time.Unix(from, 0), time.Unix(to, 0))
		}
		age := now.Year() - d.Year()
		if age < x.minAge || age > x.maxAge {
			t.Errorf("expected age %d to be in range [%d, %d]", age, x.minAge, x.maxAge)
		}
	}
}

func TestDateBirthdayInvalidRange(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			// ignore
		}
	}()

	Date().Birthday(45, 25)
	t.Errorf("expected invalid age range panic")
}
