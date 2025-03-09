package log

import "time"

func GetElapsedTime(runnable func()) time.Duration {
	start := time.Now()
	runnable()
	return time.Since(start)
}
