package clock

import (
	"fmt"
	//"strconv"
)

type Clock struct {
	hour int
	minute int
}

const t = 24 * 60

func New(hour, minute int) Clock {
	hour = (hour + minute/60) % 24
	minute = minute % 60

	if hour < 0 {
		hour += 24
	}

	if minute < 0 {
		hour -= 1
		if hour < 0 {
			hour += 24
		}
		minute += 60
	}

	return Clock{hour, minute}
}

// Add func
func (c Clock) Add(m int) Clock{
	return New(c.hour + (c.minute + m)/60,(c.minute + m) % 60)
}

// Subtract func
func (c Clock) Subtract(m int) Clock{
	return New(c.hour + (c.minute - m)/60, (c.minute - m)% 60)
}

// String func
func (c Clock) String() string{
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}