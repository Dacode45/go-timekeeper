package timekeeping

import "time"

// Interval is a block of time something was worked on
type Interval struct {
	IntervalID int `gorm:"primary_key"`
	IsClosed   bool
	Start      time.Time
	End        time.Time
}

// NewInterval creates a new interval
func NewInterval(startTime time.Time) *Interval {
	i := &Interval{
		Start:    startTime,
		IsClosed: false,
	}
	db.Create(i)
	return i
}

// EndInterval closes the interval.
func (i *Interval) EndInterval(endTime time.Time) {
	i.IsClosed = true
	i.End = endTime
	db.Update(i)
}
