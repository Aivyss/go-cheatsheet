package time

import (
	"github.com/stretchr/testify/assert"
	tm "go-cheatsheet/time"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	t.Run("same date (count first day)", func(t *testing.T) {
		// given
		from := time.Date(2024, 11, 20, 0, 0, 0, 0, time.Local)
		to := time.Date(2024, 11, 20, 0, 0, 0, 0, time.Local)

		period := tm.Period{
			From: tm.Date{
				Year:  from.Year(),
				Month: from.Month(),
				Day:   from.Day(),
			},
			To: tm.Date{
				Year:  to.Year(),
				Month: to.Month(),
				Day:   to.Day(),
			},
			CountFirstDay: true,
		}

		// when, then
		assert.True(t, period.LargerThanOrEqual(0, 0, 1))
		assert.False(t, period.LargerThan(0, 0, 1))
		assert.False(t, period.LargerThanOrEqual(0, 0, 2))
		assert.True(t, period.SmallerThanOrEqual(0, 0, 1))
		assert.False(t, period.SmallerThan(0, 0, 1))
	})
	t.Run("same date (no count first day)", func(t *testing.T) {
		// given
		from := time.Date(2024, 11, 20, 0, 0, 0, 0, time.Local)
		to := time.Date(2024, 11, 20, 0, 0, 0, 0, time.Local)

		period := tm.Period{
			From: tm.Date{
				Year:  from.Year(),
				Month: from.Month(),
				Day:   from.Day(),
			},
			To: tm.Date{
				Year:  to.Year(),
				Month: to.Month(),
				Day:   to.Day(),
			},
		}

		// when, then
		assert.False(t, period.LargerThanOrEqual(0, 0, 1))
		assert.False(t, period.LargerThan(0, 0, 1))
		assert.False(t, period.LargerThanOrEqual(0, 0, 2))
		assert.True(t, period.SmallerThanOrEqual(0, 0, 1))
		assert.True(t, period.SmallerThan(0, 0, 1))
	})
	t.Run("n days (count first day)", func(t *testing.T) {
		// given
		days := 90
		from := time.Date(2024, 11, 20, 0, 0, 0, 0, time.Local)
		for i := 0; i < days; i++ {
			day := i + 1
			to := from.AddDate(0, 0, day)
			period := tm.Period{
				From: tm.Date{
					Year:  from.Year(),
					Month: from.Month(),
					Day:   from.Day(),
				},
				To: tm.Date{
					Year:  to.Year(),
					Month: to.Month(),
					Day:   to.Day(),
				},
				CountFirstDay: true,
			}

			// when, then
			for j := 0; j <= day; j++ {
				assert.True(t, period.LargerThan(0, 0, j))
				assert.True(t, period.LargerThanOrEqual(0, 0, j))
				assert.False(t, period.SmallerThan(0, 0, j))
				assert.False(t, period.SmallerThanOrEqual(0, 0, j))
			}
			assert.True(t, period.LargerThanOrEqual(0, 0, day+1))
			assert.False(t, period.LargerThan(0, 0, day+1))
			assert.False(t, period.LargerThanOrEqual(0, 0, day+2))
			assert.False(t, period.LargerThan(0, 0, day+2))

			assert.False(t, period.SmallerThan(0, 0, day+1))
			assert.True(t, period.SmallerThanOrEqual(0, 0, day+1))
			assert.True(t, period.SmallerThan(0, 0, day+2))
			assert.True(t, period.SmallerThanOrEqual(0, 0, day+2))
		}
	})
	t.Run("n days (no count first day)", func(t *testing.T) {
		// given
		days := 90
		from := time.Date(2024, 11, 20, 0, 0, 0, 0, time.Local)
		for i := 0; i < days; i++ {
			day := i + 1
			to := from.AddDate(0, 0, day)
			period := tm.Period{
				From: tm.Date{
					Year:  from.Year(),
					Month: from.Month(),
					Day:   from.Day(),
				},
				To: tm.Date{
					Year:  to.Year(),
					Month: to.Month(),
					Day:   to.Day(),
				},
			}

			// when, then
			for j := 0; j < day; j++ {
				assert.True(t, period.LargerThan(0, 0, j))
				assert.True(t, period.LargerThanOrEqual(0, 0, j))
				assert.False(t, period.SmallerThan(0, 0, j))
				assert.False(t, period.SmallerThanOrEqual(0, 0, j))
			}

			assert.False(t, period.LargerThan(0, 0, day))
			assert.True(t, period.LargerThanOrEqual(0, 0, day))
			assert.False(t, period.SmallerThan(0, 0, day))
			assert.True(t, period.SmallerThanOrEqual(0, 0, day))

			assert.False(t, period.LargerThan(0, 0, day+1))
			assert.False(t, period.LargerThanOrEqual(0, 0, day+1))
			assert.True(t, period.SmallerThan(0, 0, day+1))
			assert.True(t, period.SmallerThanOrEqual(0, 0, day+1))
		}
	})
}
