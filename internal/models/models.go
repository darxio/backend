package models

// easyjson:json
type User struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
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
	ID   int32  `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// easyjson:json
type IngredientArr []*Ingredient

// easyjson:json
type ProductExtended struct {
	Barcode      uint64 `json:"barcode"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Contents     string `json:"contents"`
	CategoryURL  string `json:"category_url"`
	Mass         string `json:"mass"`
	BestBefore   string `json:"best_before"`
	Nutrition    string `json:"nutrition"`
	Manufacturer string `json:"manufacturer"`
	Image        string `json:"image"`
}

// easyjson:json
type ProductShrinked struct {
	Barcode uint64 `json:"barcode"`
	Name    string `json:"name"`
}

// easyjson:json
type ProductExtendedArr []*ProductExtended
