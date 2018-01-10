package util

type JsonData struct {
	Title      string `json:"title"`
	Summary    string `json:"summary"`
	Url        string `json:"url"`
	OriginUrl  string `json:"origin_url"`
	OriginName string `json:"origin_name"`
	ImageUrl   string `json:"image_url"`
	ImageTitle string `json:"image_title"`
	InputDate  string `json:"input_date"`
}
