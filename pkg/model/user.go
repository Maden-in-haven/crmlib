package model

type User struct {
	ID           string
	Username     string
	PasswordHash string
	Role         string
	CreatedAt    string
	UpdatedAt    string
	IsDeleted    bool
}