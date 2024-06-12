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

	CatalogsResponse struct {
		Title    string
		CoverUrl string
	}

	ItemsReposne struct {
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
