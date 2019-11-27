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
	ID          int32   `json:"id"`
	Name        string  `json:"name"`
	Danger      int     `json:"danger"`
	Description string  `json:"description"`
	WikiLink    string  `json:"wiki_link"`
	Groups      []int64 `json:"groups"`
	// Type string `json:"type"`
}

// easyjson:json
type IngredientArr []*Ingredient

// easyjson:json
type ProductExtended struct {
	Shrinked     bool        `json:"shrinked"`
	Barcode      uint64      `json:"barcode"`
	Name         string      `json:"name"`
	Description  string      `json:"description"`
	Contents     string      `json:"contents"`
	CategoryURL  string      `json:"category_url"`
	Mass         string      `json:"mass"`
	BestBefore   string      `json:"best_before"`
	Nutrition    string      `json:"nutrition"`
	Manufacturer string      `json:"manufacturer"`
	Image        string      `json:"image"`
	Ingredients  interface{} `json:"ingredients"`
}

// easyjson:json
type ProductShrinked struct {
	Shrinked bool   `json:"shrinked"`
	Barcode  uint64 `json:"barcode"`
	Name     string `json:"name"`
}

// easyjson:json
type ProductExtendedArr []*ProductExtended

// easyjson:json
type ProductShrinkedArr []*ProductShrinked

// easyjson:json
type ProductToAdd struct {
	Barcode uint64 `json:"barcode"`
	Name    string `json:"name"`
}
