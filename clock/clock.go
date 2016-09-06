package clock

import "fmt"

const (
	testVersion     = 4
	minutes_in_day  = 1440
	minutes_in_hour = 60
)

type Clock int

func New(hour, minute int) Clock {
	new_clock := Clock(((hour * minutes_in_hour) + minute) % minutes_in_day)

	return new_clock.CheckForNegativeTime()
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock/minutes_in_hour, clock%minutes_in_hour)
}

func (clock Clock) Add(minutes int) Clock {
	clock = (clock + Clock(minutes)) % minutes_in_day

	return clock.CheckForNegativeTime()
}

func (clock Clock) CheckForNegativeTime() Clock {
	if clock < 0 {
		clock += minutes_in_day
	}

	return clock
}
