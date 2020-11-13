package linkedin

type OrganizationNode struct {
	Paging     Paging         `json:"paging,omitempty"`
	RecipeType string         `json:"$recipeType,omitempty"`
	Elements   []Organization `json:"elements,omitempty"`
}

type Organization struct {
	EntityUrn               string      `json:"entityUrn,omitempty"`
	DateRange               DateRange   `json:"dateRange,omitempty"`
	Name                    string      `json:"name,omitempty"`
	MultiLocaleName         MultiLocale `json:"multiLocaleName,omitempty"`
	RecipeType              string      `json:"$recipeType,omitempty"`
	PositionHeld            string      `json:"positionHeld,omitempty"`
	MultiLocalePositionHeld MultiLocale `json:"multiLocalePositionHeld,omitempty"`
}
