package models

import (
	"strings"

	"github.com/jinzhu/gorm"
)

// CurrentUser holds the struct of the currently
// highlighted project.
var CurrentUser User

// User model.
type User struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100);unique;not null"`
	Description string
	Tasks       []Task
}

// AllUsers queries the database for, and returns, all projects
// after scanning them into a slice of structs.
func AllUsers() []User {
	var p []User
	o := Setting.SortBy + " " + Setting.SortOrder
	DB.Order(o).Find(&p)
	return p
}

// GetUsers queries the database for, and returns, one project
// after scanning it into the struct.
func GetUser(n string) User {
	var p User
	n = strings.TrimSpace(n)
	DB.Where(&User{Name: n}).First(&p)
	return p
}

// AddUser queries the database for one project by name (project names are unique).
// If the record exists then it is returned;
// else, it will create the record and return that one.
func AddUser(n string) User {
	var p User
	n = strings.TrimSpace(n)
	DB.FirstOrCreate(&p, User{Name: n})
	return p
}

// Delete one project and its children.
func (p User) Delete() {
	tasks := AllTasks(p)
	for _, t := range tasks {
		t.Delete()
	}
	DB.Delete(&p)
}

// UpdateUser will update the user with values defined outside.
func UpdateUser(p User) User {
	DB.Save(&p)
	return p
}

