package models

type Project struct {
	projectID   int
	projectName string
	description string
	ownerID     int
	status      string
	startDate   string
	endDate     string
	createAt    string
	updateAt    string
}
