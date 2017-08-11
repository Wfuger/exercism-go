package gigasecond

import "time"


const (
	testVersion = 4
	gigSec = 1000000000
)

func AddGigasecond(t time.Time) time.Time {
	duration := time.Duration(gigSec * time.Second)
	t = t.Add(duration)
	return t
}
