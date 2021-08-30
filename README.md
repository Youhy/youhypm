## Overview
I created this in response to my need to record entries for tasks and tasks for projects, and my want to not have to type * Minimal requirements:
** Create user
* Minimal set of properties:
** email
** first name
** last name
** Delete user
** Get user
** Get users
** Create a task
* task has to be assigned to a user
** task can occupy some certain amount of time
** task can’t overlap with other tasks
** reminder period
** Delete task
** Get task
** Get tasks
** get list of the task

## Getting Started
Upon starting the program you will be required to type in the name of a project. Navigate to the right and press Ctrl-A to add a task, and do the same to create and start an entry. Press Ctrl-S to save the entry.

### Installation
Download the [latest release](https://github.com/youhy/youhypm) of the binary or

`go get github.com//youhy/youhypm`
### Libraries
* [gocui](https://github.com/jroimartin/gocui)
* [gorm](https://github.com/jinzhu/gorm)
