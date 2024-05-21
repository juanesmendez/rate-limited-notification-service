package model

import "time"

type RateLimitRule struct {
	Limit      int
	TimeWindow time.Duration
}
