package models

type (
	CatalogItemRequest struct {
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Price       float32 `json:"price"`
		CoverUrl    string  `json:"cover_url"`
		ShopID      int     `json:"shop_id"`
	}

	ShopCreateRequest struct {
		Title          string `json:"title"`
		Currency       string `json:"currency"`
		ChannelUrl     string `json:"channel_url"`
		TelegramUserID int    `json:"telegram_user_id"`
	}

	UnitStruct struct {
		Title string `json:"title"`
		Price int    `json:"price"`
	}

	CreateOrderRequest struct {
		UserId   int          `json:"user_id"`
		DataList []UnitStruct `json:"data_list"`
	}
)
