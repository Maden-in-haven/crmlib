package model

type User struct {
	ID           string
	Username     string
	PasswordHash string
	Role         string
	CreatedAt    string
	UpdatedAt    string
}

type Admin struct {
	ID          string
	Username    string
	Permissions map[string]interface{}
	CreatedAt   string
}

type Client struct {
	ID          string
	Username    string
	FullName    string
	PhoneNumber string
	CreatedAt   string
}

type Manager struct {
	ID        string
	Username  string
	FullName  string
	HireDate  string
	CreatedAt string
}

type UserLog struct {
	ID        string
	UserID    string
	Action    string
	Timestamp string
}
