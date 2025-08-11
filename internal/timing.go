package internal

import (
	"time"
)

type Timing struct {
	Remaining time.Duration
	Expiry    time.Time
}

func NewTimer(remaining time.Duration) *Timing {
	return &Timing{
		Remaining: remaining,
		Expiry:    time.Now().Add(remaining),
	}
}

var TimeRemaining time.Duration
