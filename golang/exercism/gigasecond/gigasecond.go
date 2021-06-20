// Package gigasecond calculates the time at which someone has lived for a gigasecond.
package gigasecond

import "time"

// AddGigasecond returns the time t + 10^9 seconds
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9 * time.Second)
}
