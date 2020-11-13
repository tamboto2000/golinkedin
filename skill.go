package linkedin

// SkillNode contains list of profile skills
type SkillNode struct {
	Paging     Paging  `json:"paging,omitempty"`
	Elements   []Skill `json:"elements,omitempty"`
	RecipeType string  `json:"$recipeType,omitempty"`
}

type Skill struct {
	Name            string       `json:"name,omitempty"`
	MultiLocaleName *MultiLocale `json:"multiLocaleName,omitempty"`
	RecipeType      string       `json:"$recipeType,omitempty"`
	EntityUrn       string       `json:"entityUrn,omitempty"`
}
