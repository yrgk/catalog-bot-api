package structs

type (
	SimpleResponse struct {
		IsOk       bool   `json:"is_ok"`
		Message    string `json:"message"`
		StatusCode int    `json:"status_code"`
	}

	CatalogItemResponse struct {
		ID          uint    `json:"id"`
		Title       string  `json:"title"`
		Price       float32 `json:"price"`
		Description string  `json:"description"`
		CoverUrl    string  `json:"cover_url"`
		Currency    string  `json:"currency"`
		ShopID      int     `json:"shop_id"`
	}

	CatalogResponse struct {
		ShopTitle string                `json:"shop_title"`
		Items     []CatalogItemResponse `json:"items"`
	}
)
