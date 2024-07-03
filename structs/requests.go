package structs

type (
	CatalogItemRequest struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Price       int    `json:"price"`
		CoverUrl    string `json:"cover_url"`
		Currency    string `json:"currency"`
		// CatalogID   int    `json:"catalog_id"`
	}
)