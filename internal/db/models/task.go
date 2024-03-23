package models

type TaskGroup struct {
	id          int
	projectID   int
	name        string
	description string
	createAt    string
	updateAt    string
}

type Task struct {
	id          int
	taskGroupID int
	title       string
	description string
	assigneeID  int
	status      string
	prirority   string
	startDate   string
	endDate     string
	createAt    string
	updateAt    string
}
