package models

// easyjson:json
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// easyjson:json
type Msg struct {
	Message string `json:"message"`
}

// easyjson:json
type Group struct {
	Name  string `json:"name"`
	About string `json:"about"`
}
