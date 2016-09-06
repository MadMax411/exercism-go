package gigasecond

import "time"

const testVersion = 4

// Adds 1e9 (1 billion) seconds to t
func AddGigasecond(t time.Time) time.Time {
  return t.Add(1e9 * time.Second)
}
