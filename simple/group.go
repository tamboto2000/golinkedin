package simple

type Group struct {
	ID                int       `json:"id"`
	URL               string    `json:"url"`
	Picture           *Image    `json:"picture"`
	BackgroundPicture *Image    `json:"backgroundPicture"`
	GroupName         string    `json:"groupName"`
	MemberCount       int       `json:"memberCount"`
	Description       string    `json:"description"`
	Discoverability   string    `json:"discoverability"`
	Created           *Date     `json:"created"`
	Rules             string    `json:"rules"`
	RelatedGroups     []Group   `json:"relatedGroups"`
	Owner             *Profile  `json:"owner"`
	Admins            []Profile `json:"admins"`
}
