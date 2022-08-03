package articles

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type ErrorModel struct {
	StatusCode int    `json:"StatusCode"`
	Message    string `json:"Message"`
}
