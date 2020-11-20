package simple

type Company struct {
	UniversalName        string `json:"universalName,omitempty"`
	URL                  string `json:"url,omitempty"`
	Logo                 *Image `json:"logo,omitempty"`
	BackgroundCoverImage *Image `json:"BackgroundCoverImage,omitempty"`
	CompanyName          string `json:"companyName,omitempty"`
	Industry             string `json:"Industry,omitempty"`
	FollowerCount        int    `json:"followerCount,omitempty"`
	Tagline              string `json:"tagline,omitempty"`
	Description          string `json:"description,omitempty"`
	// StaffCountRange is staff count range based on company's admin input
	StaffCountRange int `json:"staffCountRange,omitempty"`
	// StaffCount is staff count based on people on linkedin worked in the company
	StaffCount   int        `json:"staffCount,omitempty"`
	CompanyType  string     `json:"companyType,omitempty"`
	Founded      int        `json:"founded,omitempty"`
	Specialities []string   `json:"specialities,omitempty"`
	Locations    []Location `json:"locations,omitempty"`
}
