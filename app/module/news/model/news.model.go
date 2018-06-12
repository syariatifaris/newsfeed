package model

type News struct {
	NewsID      int64  `json:"news_id,omitempty" db:"news_id"`
	Title       string `json:"title" db:"title"`
	HtmlContent string `json:"html_content" db:"html_content"`
	Status      string `json:"status" db:"status"`
	IsValid     int64  `json:"is_valid" db:"is_valid"`
}
