package common

type Game struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ImgUrl      string `json:"img_url"`
	Description string `json:"description"`
	Link        string `json:"link"`
}
