package linkedin

const (
	Rank1  = "F"
	Rank2  = "S"
	Rank3  = "T"
	People = "PEOPLE"
)

type SearchFilter struct {
	CurrentCompany []string `json:"currentCompant,omitempty"`
	GeoURN         []string `json:"geoUrn"`
	Network        []string `json:"network"`
	ConnectionOf   string   `json:"connectionOf"`
	ResultType     string   `json:"resultType"`
}
