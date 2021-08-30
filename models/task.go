package models

import (
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// CurrentTask holds the struct of the currently
// highlighted task.
var CurrentTask Task

// Task model.
type Task struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	User     User
	UserID   uint `gorm:"index;not null"`
	TotalTime   time.Duration
}

// AllTasks queries the database for, and returns, all tasks after scanning them into a slice.
func AllTasks(p User) []Task {
	var t []Task
	DB.Model(&p).Related(&t)
	return t
}

// GetTask queries the database for, and returns, one task
// after scanning it into the struct.
func GetTask(n string) Task {
	var t Task
	n = strings.TrimSpace(n)
	// DB.Where(&Task{Name: n}).First(&t)
	DB.Model(&CurrentUser).Where(&Task{Name: n}).Related(&t)
	return t
}

// AddTask queries the database for one project by name.
// If the record exists then it is returned;
// else, it will create the record and return that one.
func AddTask(n string, u User) Task {
	var t Task
	n = strings.TrimSpace(n)
	DB.FirstOrCreate(&t, Task{Name: n, UserID: u.ID})
	return t
}

// Delete one task and its children.
func (t Task) Delete() {
	DB.Delete(&t)
}

