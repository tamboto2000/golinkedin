package linkedin

type HonorNode struct {
	Paging     Paging  `json:"paging,omitempty"`
	RecipeType string  `json:"$recipeType,omitempty"`
	Elements   []Honor `json:"elements,omitempty"`
}

type Honor struct {
	MultiLocaleTitle       MultiLocale     `json:"multiLocaleTitle,omitempty"`
	Description            string          `json:"description,omitempty"`
	OccupationUnion        OccupationUnion `json:"occupationUnion,omitempty"`
	Title                  string          `json:"title,omitempty"`
	Issuer                 string          `json:"issuer,omitempty"`
	IssuedOn               IssuedOn        `json:"issuedOn,omitempty"`
	EntityUrn              string          `json:"entityUrn,omitempty"`
	RecipeType             string          `json:"$recipeType,omitempty"`
	MultiLocaleIssuer      MultiLocale     `json:"multiLocaleIssuer,omitempty"`
	MultiLocaleDescription MultiLocale     `json:"multiLocaleDescription,omitempty"`
}

type IssuedOn struct {
	Month int64 `json:"month,omitempty"`
	Year  int64 `json:"year,omitempty"`
}

type OccupationUnion struct {
	ProfilePosition string `json:"profilePosition,omitempty"`
}
