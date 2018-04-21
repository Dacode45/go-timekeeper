package timekeeping

// Tag is an attribute of a task
type Tag struct {
	Name        string `sql:"size:255;index:name_idx"`
	Description string
}
