package models

//User ...
type User struct {
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
	Profile  string `json:"profile,omitempty"`
}
