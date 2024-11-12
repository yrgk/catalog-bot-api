package models

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
		Currency    string  `json:"currency"`
		Description string  `json:"description"`
		CoverUrl    string  `json:"cover_url"`
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
		Banners   []Banner     `json:"banners"`
		Items     []CatalogItemResponse `json:"items"`
	}

	MyShopsResponse struct {
		ID    uint   `json:"id"`
		Title string `json:"title"`
	}

	ShopData struct {
		Title    string `json:"shop_title"`
		Currency string `json:"currency"`
	}
)
