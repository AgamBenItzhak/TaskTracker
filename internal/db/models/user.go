package models

type User struct {
	userID       int
	email        string
	passwordHash string
	passwordSalt string
	firstName    string
	lastName     string
	createAt     string
	updateAt     string
	lastSeen     string
}

type UserProject struct {
	userID    int
	projectID int
	role      string
	createAt  string
	updateAt  string
}
