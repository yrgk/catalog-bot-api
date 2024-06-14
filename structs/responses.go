package structs

type (
	SimpleResponse struct {
		IsOk       bool   `json:"is_ok"`
		Message    string `json:"message"`
		StatusCode int    `json:"status_code"`
	}

	// CatalogResponse struct {
	// 	Title    string
	// 	CoverUrl string
	// 	ShopID   int
	// }

	CatalogItemReponse struct {
		Title    string
		Price    int
		CoverUrl string
		Currency string
	}

	OneCatalogItemResponse struct {
		Title       string
		Description string
		Price       int
		CoverUrl    string
		Currency    string
		CatalogID   int
	}
)
