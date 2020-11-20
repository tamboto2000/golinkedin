package simple

type Image struct {
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	ExpiresAt int64  `json:"expiresAt"`
	URL       string `json:"url"`
}
