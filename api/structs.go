package api

type (
	BadResponse struct {
		is_ok   bool
		message string
	}

	SuccessfulResponse struct {
		is_ok   bool
		message string
	}

	CatalogResponse struct {
		Title    string
		CoverUrl string
		ShopID   int
	}

	ItemReponse struct {
		Title    string
		Price    int
		CoverUrl string
		Currency string
	}

	OneItemResponse struct {
		Title       string
		Description string
		Price       int
		CoverUrl    string
		Currency    string
		CatalogID   int
		CategoryID  int
	}
)
