package timekeeping

import "fmt"

// FindTask by Name
func FindTask(name string) []Task {
	var tasks []Task
	fmt.Println(name)
	db.Find(&tasks, "name = ?", name)
	return tasks
}

// ListTask list all tasks
func ListTask() []Task {
	var tasks []Task
	db.Find(&tasks)
	return tasks
}

// ListClosedTask list all closed tasks
func ListClosedTask(closed bool) []Task {
	var tasks []Task
	db.Find(&tasks, "is_closed = ?", closed)
	return tasks
}
