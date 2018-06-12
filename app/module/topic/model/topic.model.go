package model

type Topic struct {
	TopicID int64  `json:"topic_id" db:"topic_id"`
	Name    string `json:"name" db:"name"`
	IsValid int64  `json:"is_valid" db:"is_valid"`
}

type TopicNews struct {
	TopicID int64 `json:"topic_id" db:"topic_id"`
	NewsID  int64 `json:"news_id" db:"news_id"`
	IsValid int64 `json:"is_valid" db:"is_valid"`
}
