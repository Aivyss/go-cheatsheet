package time

import "time"

type Date struct {
	Year  int        `json:"year"`
	Month time.Month `json:"month"`
	Day   int        `json:"day"`
	Loc   *time.Location
}

func (d *Date) GetDate() time.Time {
	loc := d.Loc
	if loc == nil {
		loc = time.Local
	}

	return time.Date(d.Year, d.Month, d.Day, 0, 0, 0, 0, loc)
}

func (d *Date) Format(format string) string {
	return d.GetDate().Format(format)
}

type Period struct {
	From          Date `json:"from"`
	To            Date `json:"to"`
	CountFirstDay bool
}

func (p *Period) LargerThan(year, month, day int) bool {
	if p.CountFirstDay {
		day -= 1
	}

	shiftDate := p.From.GetDate().AddDate(year, month, day)
	return shiftDate.Before(p.To.GetDate())
}
func (p *Period) LargerThanOrEqual(year, month, day int) bool {
	if p.CountFirstDay {
		day -= 1
	}

	shiftDate := p.From.GetDate().AddDate(0, 0, day)
	to := p.To.GetDate()
	return shiftDate.Before(to) || shiftDate.Equal(to)
}
func (p *Period) SmallerThan(year, month, day int) bool {
	if p.CountFirstDay {
		day -= 1
	}

	shiftDate := p.From.GetDate().AddDate(year, month, day)
	return shiftDate.After(p.To.GetDate())
}
func (p *Period) SmallerThanOrEqual(year, month, day int) bool {
	if p.CountFirstDay {
		day -= 1
	}

	to := p.To.GetDate()
	shiftDate := p.From.GetDate().AddDate(0, 0, day)
	return shiftDate.After(to) || shiftDate.Equal(to)
}
