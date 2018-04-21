package timekeeping

import "time"

// Task contains the intervals a task was worked on, a description, and some tags
type Task struct {
	TaskID      int `gorm:"primary_key"`
	Name        string
	IsClosed    bool
	Intervals   []*Interval
	Description string
	Tags        []*Tag
}

// NewTask creates a new task
func NewTask(name, description string, startTime time.Time) *Task {
	// Create a new interval and adds it to the list
	interval := NewInterval(startTime)
	task := &Task{
		Name:        name,
		Description: description,
		Intervals:   []*Interval{interval},
	}
	db.Create(task)
	return task
}

// EndTask Closes the last interval
func (t *Task) EndTask(endTime time.Time) {
	i := t.Intervals[len(t.Intervals)-1]
	i.EndInterval(endTime)
	t.IsClosed = true
	db.Update(t)
}
