package structs

type (
	SimpleResponse struct {
		IsOk       bool   `json:"is_ok"`
		Message    string `json:"message"`
		StatusCode int    `json:"status_code"`
	}

	OneCatalogItemResponse struct {
		ID          uint    `json:"id"`
		Title       string  `json:"title"`
		Price       float32 `json:"price"`
		Description string  `json:"description"`
		CoverUrl    string  `json:"cover_url"`
		Currency    string  `json:"currency"`
		// ShopID      int     `json:"shop_id"`
	}

	CatalogItemResponse struct {
		ID       uint    `json:"id"`
		Title    string  `json:"title"`
		Price    float32 `json:"price"`
		CoverUrl string  `json:"cover_url"`
	}

	CatalogResponse struct {
		ShopTitle string                `json:"shop_title"`
		Currency  string                `json:"currency"`
		Items     []CatalogItemResponse `json:"items"`
	}

	MyShopsResponse struct {
		ID       uint   `json:"id"`
		CoverUrl string `json:"cover_url"`
		Title    string `json:"title" gorm:"not null; unique"`
	}

	ShopData struct {
		Title string `json:"shop_title"`
		Currency  string `json:"currency"`
	}
)
