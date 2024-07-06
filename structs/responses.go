package structs

type (
	SimpleResponse struct {
		IsOk       bool   `json:"is_ok"`
		Message    string `json:"message"`
		StatusCode int    `json:"status_code"`
	}

	CatalogItemResponse struct {
		ID       uint   `json:"id"`
		Title    string `json:"title"`
		Price    float32    `json:"price"`
		CoverUrl string `json:"cover_url"`
		Currency string `json:"currency"`
	}

	OneCatalogItemResponse struct {
		Title       string `json:"title"`
		Price       float32    `json:"price"`
		Description string `json:"description"`
		CoverUrl    string `json:"cover_url"`
		Currency    string `json:"currency"`
		CatalogID   int    `json:"catalog_id"`
	}
)
