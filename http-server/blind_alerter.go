package poker

import "time"

// BlindAlerter is a implementation interface to wrap time.AfterFunc
type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}
