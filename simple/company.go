package simple

type Company struct {
	UniversalName        string `json:"universalName"`
	URL                  string `json:"url"`
	Logo                 *Image `json:"logo,omitempty"`
	BackgroundCoverImage *Image `json:"BackgroundCoverImage"`
	CompanyName          string `json:"companyName,omitempty"`
	Industry             string `json:"Industry,omitempty"`
	FollowerCount        int    `json:"followerCount"`
	Tagline              string `json:"tagline"`
	Description          string `json:"description"`
	// StaffCountRange is staff count range based on company's admin input
	StaffCountRange int `json:"staffCountRange"`
	// StaffCount is staff count based on people on linkedin worked in the company
	StaffCount   int        `json:"staffCount"`
	CompanyType  string     `json:"companyType"`
	Founded      int        `json:"founded"`
	Specialities []string   `json:"specialities"`
	Locations    []Location `json:"locations"`
}
