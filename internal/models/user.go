package models

// easyjson:json
type User struct {
	Username string `json:"nickname"`
	Password string `json:"password"`
}
