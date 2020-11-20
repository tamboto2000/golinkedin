package simple

type DateRange struct {
	Start *Date `json:"date,omitempty"`
	End   *Date `json:"end,omitempty"`
}
