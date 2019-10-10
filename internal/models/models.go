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
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	About string `json:"about"`
}

// easyjson:json
type GroupArr []*Group

// easyjson:json
type Ingredient struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	About string `json:"about"`
}

// easyjson:json
type IngredientArr []*Ingredient

// easyjson:json
type Product struct {
	ID      int32  `json:"id"`
	Name    string `json:"name"`
	Barcode string `json:"barcode"`
}

// easyjson:json
type ProductArr []*Product
