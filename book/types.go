package book

type EntitlementBook struct {
	Id                string            `json:"id"`
	UserId            string            `json:"userId"`
	ProductId         string            `json:"productId"`
	ProductName       string            `json:"productName"`
	ReleaseDate       string            `json:"releaseDate"`
	EntitlementSource string            `json:"entitlementSource"`
	EntitlementLink   string            `json:"entitlementLink"`
	CreatedAt         string            `json:"createdAt"`
	UpdatedAt         string            `json:"updatedAt"`
	InfoURL           map[string]string `json:"infoURL"`
	LocalFilenames    map[string]string `json:"localFilenames"`
	TitlePrinted      bool              `json:"titlePrinted"`
}

type ResponseData struct {
	Data    string `json:"data"`
	Message string `json:"message"`
}
