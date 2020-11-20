package simple

type Skill struct {
	Name      string    `json:"name"`
	Endorsers []Profile `json:"endorsers"`
}
