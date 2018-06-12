package model

type AddNewsRequest struct {
	News   *News    `json:"news"`
	Topics []string `json:"topics"`
}

type GetByTopicsRequest struct {
	Topics []string `json:"topics"`
}
