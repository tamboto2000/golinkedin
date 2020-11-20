package simple

type Image struct {
	Width     int    `json:"width,omitempty"`
	Height    int    `json:"height,omitempty"`
	ExpiresAt int64  `json:"expiresAt,omitempty"`
	URL       string `json:"url,omitempty"`
}
