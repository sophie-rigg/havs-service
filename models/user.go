package models

// User represents a user of the system.
type User struct {
	ID   string `json:"id"` // uuid
	Name string `json:"name"`
}

// NewUser creates a new user.
func NewUser(id string) *User {
	return &User{
		ID: id,
	}
}

// SetName sets the name of the user.
func (u *User) SetName(name string) {
	u.Name = name
}
