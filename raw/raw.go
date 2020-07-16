package raw

type ViewProfile struct {
	PositionGroupView         PositionGroupView       `json:"positionGroupView"`
	PatentView                EducationViewClass      `json:"patentView"`
	SummaryTreasuryMediaCount int64                   `json:"summaryTreasuryMediaCount"`
	SummaryTreasuryMedias     []interface{}           `json:"summaryTreasuryMedias"`
	EducationView             EducationViewClass      `json:"educationView"`
	OrganizationView          OrganizationView        `json:"organizationView"`
	ProjectView               ProjectView             `json:"projectView"`
	PositionView              PositionView            `json:"positionView"`
	Profile                   Profile                 `json:"profile"`
	LanguageView              EView                   `json:"languageView"`
	CertificationView         CertificationViewClass  `json:"certificationView"`
	TestScoreView             EView                   `json:"testScoreView"`
	VolunteerCauseView        VolunteerCauseView      `json:"volunteerCauseView"`
	EntityUrn                 string                  `json:"entityUrn"`
	CourseView                CertificationViewClass  `json:"courseView"`
	HonorView                 CertificationViewClass  `json:"honorView"`
	SkillView                 SkillView               `json:"skillView"`
	VolunteerExperienceView   VolunteerExperienceView `json:"volunteerExperienceView"`
	PrimaryLocale             Locale                  `json:"primaryLocale"`
	PublicationView           CertificationViewClass  `json:"publicationView"`
}

type CertificationViewClass struct {
	Paging    Paging                     `json:"paging"`
	EntityUrn string                     `json:"entityUrn"`
	ProfileID string                     `json:"profileId"`
	Elements  []CertificationViewElement `json:"elements"`
}

type CertificationViewElement struct {
	EntityUrn     string           `json:"entityUrn"`
	Authority     string           `json:"authority"`
	Name          string           `json:"name"`
	TimePeriod    PurpleTimePeriod `json:"timePeriod"`
	LicenseNumber string           `json:"licenseNumber"`
	Company       MiniCompanyClass `json:"company"`
	DisplaySource string           `json:"displaySource"`
	CompanyUrn    string           `json:"companyUrn"`
	URL           string           `json:"url"`
}

type MiniCompanyClass struct {
	ObjectUrn     string  `json:"objectUrn"`
	EntityUrn     string  `json:"entityUrn"`
	Name          string  `json:"name"`
	Showcase      bool    `json:"showcase"`
	Active        bool    `json:"active"`
	Logo          Picture `json:"logo"`
	UniversalName string  `json:"universalName"`
	TrackingID    string  `json:"trackingId"`
}

type Picture struct {
	COMLinkedinCommonVectorImage COMLinkedinCommonVectorImage `json:"com.linkedin.common.VectorImage"`
}

type COMLinkedinCommonVectorImage struct {
	Artifacts []Artifact `json:"artifacts"`
	RootURL   string     `json:"rootUrl"`
}

type Artifact struct {
	Width                         int64  `json:"width"`
	FileIdentifyingURLPathSegment string `json:"fileIdentifyingUrlPathSegment"`
	ExpiresAt                     int64  `json:"expiresAt"`
	Height                        int64  `json:"height"`
}

type PurpleTimePeriod struct {
	StartDate PurpleDate `json:"startDate"`
}

type PurpleDate struct {
	Month int64 `json:"month"`
	Year  int64 `json:"year"`
}

type Paging struct {
	Start int64         `json:"start"`
	Count int64         `json:"count"`
	Total int64         `json:"total"`
	Links []interface{} `json:"links"`
}

type EducationViewClass struct {
	Paging    Paging                 `json:"paging"`
	EntityUrn string                 `json:"entityUrn"`
	ProfileID string                 `json:"profileId"`
	Elements  []EducationViewElement `json:"elements"`
}

type EducationViewElement struct {
	Projects     []string         `json:"projects"`
	EntityUrn    string           `json:"entityUrn"`
	School       School           `json:"school"`
	TimePeriod   FluffyTimePeriod `json:"timePeriod"`
	Grade        string           `json:"grade"`
	Description  string           `json:"description"`
	Activities   string           `json:"activities"`
	DegreeName   string           `json:"degreeName"`
	SchoolName   string           `json:"schoolName"`
	FieldOfStudy string           `json:"fieldOfStudy"`
	DegreeUrn    string           `json:"degreeUrn"`
	SchoolUrn    string           `json:"schoolUrn"`
}

type School struct {
	Active     bool   `json:"active"`
	Logo       Logo   `json:"logo"`
	ObjectUrn  string `json:"objectUrn"`
	SchoolName string `json:"schoolName"`
	EntityUrn  string `json:"entityUrn"`
	TrackingID string `json:"trackingId"`
}

type Logo struct {
	COMLinkedinCommonVectorImage COMLinkedinCommonVectorImage `json:"com.linkedin.common.VectorImage"`
}

type FluffyTimePeriod struct {
	EndDate   FluffyDate `json:"endDate"`
	StartDate FluffyDate `json:"startDate"`
}

type FluffyDate struct {
	Year int64 `json:"year"`
}

type EView struct {
	Paging    Paging                `json:"paging"`
	EntityUrn string                `json:"entityUrn"`
	ProfileID string                `json:"profileId"`
	Elements  []LanguageViewElement `json:"elements"`
}

type LanguageViewElement struct {
	Name        string `json:"name"`
	EntityUrn   string `json:"entityUrn"`
	Proficiency string `json:"proficiency"`
}

type OrganizationView struct {
	Paging    Paging                    `json:"paging"`
	EntityUrn string                    `json:"entityUrn"`
	ProfileID string                    `json:"profileId"`
	Elements  []OrganizationViewElement `json:"elements"`
}

type OrganizationViewElement struct {
	Name       string             `json:"name"`
	TimePeriod PositionTimePeriod `json:"timePeriod"`
	Position   string             `json:"position"`
	EntityUrn  string             `json:"entityUrn"`
}

type PositionTimePeriod struct {
	EndDate   *PurpleDate `json:"endDate,omitempty"`
	StartDate PurpleDate  `json:"startDate"`
}

type PositionGroupView struct {
	Paging    Paging                     `json:"paging"`
	EntityUrn string                     `json:"entityUrn"`
	ProfileID string                     `json:"profileId"`
	Elements  []PositionGroupViewElement `json:"elements"`
}

type PositionGroupViewElement struct {
	TimePeriod  PositionTimePeriod `json:"timePeriod"`
	Name        string             `json:"name"`
	Positions   []PositionElement  `json:"positions"`
	Paging      Paging             `json:"paging"`
	MiniCompany *MiniCompanyClass  `json:"miniCompany,omitempty"`
	EntityUrn   string             `json:"entityUrn"`
}

type PositionElement struct {
	TimePeriod  TimePeriod   `json:"timePeriod"`
	Name        string       `json:"name"`
	Positions   []Position   `json:"positions"`
	Paging      Paging       `json:"paging"`
	EntityUrn   string       `json:"entityUrn"`
	MiniCompany *MiniCompany `json:"miniCompany,omitempty"`
}

type Position struct {
	EntityUrn       string     `json:"entityUrn"`
	CompanyName     string     `json:"companyName"`
	TimePeriod      TimePeriod `json:"timePeriod"`
	Company         *Company   `json:"company,omitempty"`
	Title           string     `json:"title"`
	CompanyUrn      string     `json:"companyUrn,omitempty"`
	Recommendations []string   `json:"recommendations"`
	LocationName    string     `json:"locationName,omitempty"`
	GeoLocationName string     `json:"geoLocationName,omitempty"`
	Projects        []string   `json:"projects"`
	GeoUrn          string     `json:"geoUrn,omitempty"`
	Region          string     `json:"region,omitempty"`
	Description     string     `json:"description"`
}

type TimePeriod struct {
	StartDate Date `json:"startDate"`
	EndDate   Date `json:"endDate,omitempty"`
}

type Date struct {
	Month int64 `json:"month"`
	Year  int64 `json:"year"`
}

type Company struct {
	MiniCompany        MiniCompany        `json:"miniCompany"`
	EmployeeCountRange EmployeeCountRange `json:"employeeCountRange"`
	Industries         []string           `json:"industries"`
	ObjectUrn          string             `json:"objectUrn"`
	EntityUrn          string             `json:"entityUrn"`
	Name               string             `json:"name"`
	Showcase           bool               `json:"showcase"`
	Active             bool               `json:"active"`
	Logo               Logo               `json:"logo"`
	UniversalName      string             `json:"universalName"`
	TrackingID         string             `json:"trackingId"`
}

type MiniCompany struct {
	ObjectUrn     string `json:"objectUrn"`
	EntityUrn     string `json:"entityUrn"`
	Name          string `json:"name"`
	Showcase      bool   `json:"showcase"`
	Active        bool   `json:"active"`
	Logo          Logo   `json:"logo"`
	UniversalName string `json:"universalName"`
	TrackingID    string `json:"trackingId"`
}

type PositionCompany struct {
	MiniCompany        MiniCompanyClass   `json:"miniCompany"`
	EmployeeCountRange EmployeeCountRange `json:"employeeCountRange"`
	Industries         []string           `json:"industries"`
}

type EmployeeCountRange struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}

type PositionView struct {
	Paging    Paging            `json:"paging"`
	EntityUrn string            `json:"entityUrn"`
	ProfileID string            `json:"profileId"`
	Elements  []PositionElement `json:"elements"`
}

type Locale struct {
	Country  string `json:"country"`
	Language string `json:"language"`
}

type Profile struct {
	Summary                       string      `json:"summary"`
	IndustryName                  string      `json:"industryName"`
	LastName                      string      `json:"lastName"`
	SupportedLocales              []Locale    `json:"supportedLocales"`
	LocationName                  string      `json:"locationName"`
	Student                       bool        `json:"student"`
	GeoCountryName                string      `json:"geoCountryName"`
	GeoCountryUrn                 string      `json:"geoCountryUrn"`
	VersionTag                    string      `json:"versionTag"`
	GeoLocationBackfilled         bool        `json:"geoLocationBackfilled"`
	Elt                           bool        `json:"elt"`
	IndustryUrn                   string      `json:"industryUrn"`
	DefaultLocale                 Locale      `json:"defaultLocale"`
	FirstName                     string      `json:"firstName"`
	ShowEducationOnProfileTopCard bool        `json:"showEducationOnProfileTopCard"`
	EntityUrn                     string      `json:"entityUrn"`
	GeoLocation                   GeoLocation `json:"geoLocation"`
	GeoLocationName               string      `json:"geoLocationName"`
	Location                      Location    `json:"location"`
	MiniProfile                   MiniProfile `json:"miniProfile"`
	Headline                      string      `json:"headline"`
}

type GeoLocation struct {
	GeoUrn string `json:"geoUrn"`
}

type Location struct {
	BasicLocation BasicLocation `json:"basicLocation"`
}

type BasicLocation struct {
	CountryCode string `json:"countryCode"`
}

type MiniProfile struct {
	FirstName        string  `json:"firstName"`
	LastName         string  `json:"lastName"`
	Occupation       string  `json:"occupation"`
	ObjectUrn        string  `json:"objectUrn"`
	EntityUrn        string  `json:"entityUrn"`
	PublicIdentifier string  `json:"publicIdentifier"`
	Picture          Picture `json:"picture"`
	BackgroundImage  Picture `json:"backgroundImage"`
	TrackingID       string  `json:"trackingId"`
}

type ProjectView struct {
	Paging    Paging               `json:"paging"`
	EntityUrn string               `json:"entityUrn"`
	ProfileID string               `json:"profileId"`
	Elements  []ProjectViewElement `json:"elements"`
}

type ProjectViewElement struct {
	Occupation  string             `json:"occupation"`
	EntityUrn   string             `json:"entityUrn"`
	Members     []Member           `json:"members"`
	TimePeriod  PositionTimePeriod `json:"timePeriod"`
	Description string             `json:"description"`
	Title       string             `json:"title"`
}

type Member struct {
	Member     MiniProfile `json:"member"`
	EntityUrn  string      `json:"entityUrn"`
	ProfileUrn string      `json:"profileUrn"`
}

type SkillView struct {
	Paging    Paging             `json:"paging"`
	EntityUrn string             `json:"entityUrn"`
	ProfileID string             `json:"profileId"`
	Elements  []SkillViewElement `json:"elements"`
}

type SkillViewElement struct {
	Name      string `json:"name"`
	EntityUrn string `json:"entityUrn"`
}

type VolunteerCauseView struct {
	Paging    Paging                      `json:"paging"`
	EntityUrn string                      `json:"entityUrn"`
	ProfileID string                      `json:"profileId"`
	Elements  []VolunteerCauseViewElement `json:"elements"`
}

type VolunteerCauseViewElement struct {
	CauseType string `json:"causeType"`
	CauseName string `json:"causeName"`
}

type VolunteerExperienceView struct {
	Paging    Paging                           `json:"paging"`
	EntityUrn string                           `json:"entityUrn"`
	ProfileID string                           `json:"profileId"`
	Elements  []VolunteerExperienceViewElement `json:"elements"`
}

type VolunteerExperienceViewElement struct {
	Role        string             `json:"role"`
	EntityUrn   string             `json:"entityUrn"`
	CompanyName string             `json:"companyName"`
	TimePeriod  PositionTimePeriod `json:"timePeriod"`
	Cause       string             `json:"cause"`
	Description *string            `json:"description,omitempty"`
	Company     *PurpleCompany     `json:"company,omitempty"`
	CompanyUrn  *string            `json:"companyUrn,omitempty"`
}

type PurpleCompany struct {
	MiniCompany MiniCompanyClass `json:"miniCompany"`
}

type ConnectionView struct {
	FollowingInfo    FollowingInfo `json:"followingInfo"`
	Distance         Distance      `json:"distance"`
	EntityUrn        string        `json:"entityUrn"`
	Following        bool          `json:"following"`
	Followable       bool          `json:"followable"`
	FollowersCount   int64         `json:"followersCount"`
	ConnectionsCount int64         `json:"connectionsCount"`
}

type Distance struct {
	Value string `json:"value"`
}

type FollowingInfo struct {
	EntityUrn     string `json:"entityUrn"`
	FollowerCount int64  `json:"followerCount"`
	Following     bool   `json:"following"`
	TrackingUrn   string `json:"trackingUrn"`
}

type EndorsedSkillView struct {
	Metadata Metadata  `json:"metadata"`
	Elements []Element `json:"elements"`
	Paging   Paging    `json:"paging"`
}

type Element struct {
	EndorsedSkills []EndorsedSkill `json:"endorsedSkills"`
	Type           Type            `json:"type"`
	CategoryName   string          `json:"categoryName"`
	EntityUrn      string          `json:"entityUrn"`
}

type EndorsedSkill struct {
	OriginalCategoryType Type          `json:"originalCategoryType"`
	Highlights           []Highlight   `json:"highlights"`
	EntityUrn            string        `json:"entityUrn"`
	EndorsedByViewer     bool          `json:"endorsedByViewer"`
	Endorsements         []Endorsement `json:"endorsements"`
	Skill                Skill         `json:"skill"`
	EndorsementCount     int64         `json:"endorsementCount"`
	ProofPoints          []interface{} `json:"proofPoints"`
}

type Endorsement struct {
	Endorser  Endorser `json:"endorser"`
	EntityUrn string   `json:"entityUrn"`
	Status    Status   `json:"status"`
}

type Endorser struct {
	Distance    Distance    `json:"distance"`
	MiniProfile MiniProfile `json:"miniProfile"`
}

type BackgroundImage struct {
	COMLinkedinCommonVectorImage COMLinkedinCommonVectorImage `json:"com.linkedin.common.VectorImage"`
}

type Highlight struct {
	Signature string `json:"signature"`
	Detail    Detail `json:"detail"`
}

type Detail struct {
	COMLinkedinVoyagerIdentityProfileEndorsedSkillElitesInfo COMLinkedinVoyagerIdentityProfileEndorsedSkillElitesInfo `json:"com.linkedin.voyager.identity.profile.endorsedSkill.ElitesInfo"`
}

type COMLinkedinVoyagerIdentityProfileEndorsedSkillElitesInfo struct {
	TotalCount int64      `json:"totalCount"`
	Elites     []Endorser `json:"elites"`
}

type Skill struct {
	Name      string `json:"name"`
	EntityUrn string `json:"entityUrn"`
}

type Metadata struct {
	VieweeEndorsementsEnabled bool   `json:"vieweeEndorsementsEnabled"`
	TrackingID                string `json:"trackingId"`
	TotalSkills               int64  `json:"totalSkills"`
}

type Value string

const (
	Distance3    Value = "DISTANCE_3"
	OutOfNetwork Value = "OUT_OF_NETWORK"
)

type EntityUrn string

const (
	UrnLiFSMiniProfileACoAABF4K48BxU4A0I2Gf3ZsmaLaGhxw1Tgn5Z8 EntityUrn = "urn:li:fs_miniProfile:ACoAABF4k48BxU4A0I2gf3zsmaLaGhxw1tgn5z8"
	UrnLiFSMiniProfileACoAABsxPCcBUUAcgVjYEc5VUmXL6HbBX3K6Y   EntityUrn = "urn:li:fs_miniProfile:ACoAABsxPCcB-u-UAcgVjYEc5vUmXL6HbBX3K6Y"
	UrnLiFSMiniProfileACoAABud3LcBOMJq3LVJz3VoomvdkMyvRwYACDQ EntityUrn = "urn:li:fs_miniProfile:ACoAABud3lcBOMJq3LVJz3voomvdkMyvRwYACDQ"
	UrnLiFSMiniProfileACoAABzwFYMBq1V4ZNNE6DLGiagnLqS9I73VxU  EntityUrn = "urn:li:fs_miniProfile:ACoAABzwFYMBq1v4-zNNE6DlGiagnLqS9i73vxU"
	UrnLiFSMiniProfileACoAAC6JOO8BATOLKZPUdcCAQRHDjn2SEizM2K  EntityUrn = "urn:li:fs_miniProfile:ACoAAC6JOO8BAT-oLKZPUdcCAQRHDjn2sEizM2k"
)

type FirstName string

const (
	Adang            FirstName = "adang"
	Andi             FirstName = "Andi"
	Payam            FirstName = "Payam "
	Thomas           FirstName = "Thomas"
	WhendyPamungkasP FirstName = "Whendy Pamungkas P."
)

type LastName string

const (
	Jasem     LastName = "Jasem"
	Kartana   LastName = "kartana"
	Prahari   LastName = "Prahari"
	TSutrisno LastName = "T. Sutrisno"
	W         LastName = "W."
)

type ObjectUrn string

const (
	UrnLiMember293114767 ObjectUrn = "urn:li:member:293114767"
	UrnLiMember456211495 ObjectUrn = "urn:li:member:456211495"
	UrnLiMember463330903 ObjectUrn = "urn:li:member:463330903"
	UrnLiMember485496195 ObjectUrn = "urn:li:member:485496195"
	UrnLiMember780744943 ObjectUrn = "urn:li:member:780744943"
)

type PublicIdentifier string

const (
	AdangKartana34794B108           PublicIdentifier = "adang-kartana-34794b108"
	PayamHeadhunterIndonesia        PublicIdentifier = "payam-headhunter-indonesia"
	SutrisnoAndi                    PublicIdentifier = "sutrisno-andi"
	ThomasW38720310B                PublicIdentifier = "thomas-w-38720310b"
	WhendyPamungkasPPrahari7271B682 PublicIdentifier = "whendy-pamungkas-p-prahari-7271b682"
)

type TrackingID string

const (
	HDqQ78LJSs2Avn9N3YrNYg TrackingID = "hDqQ78LJSs2avn9N3yrNYg=="
	LHiDRAIUSKCwF65NdxyLlw TrackingID = "lHiDrAiUSKCwF65NdxyLlw=="
	TuFuYmqbRGmpG7IXuo7YBg TrackingID = "TuFuYmqbRGmpG7IXuo7YBg=="
	WhjUnsFkTLm5CAx6LozFSA TrackingID = "WhjUnsFkTLm5CAx6LozFSA=="
	YQDaj5CNQ7S6VNYgXeQFxA TrackingID = "YQDaj5cNQ7S6vNYgXeQFxA=="
)

type Status string

const (
	Accepted Status = "ACCEPTED"
)

type Type string

const (
	IndustryKnowledge Type = "INDUSTRY_KNOWLEDGE"
	Interpersonal     Type = "INTERPERSONAL"
	None              Type = "NONE"
	ToolsTechnologies Type = "TOOLS_TECHNOLOGIES"
	Top               Type = "TOP"
)

type ContactView struct {
	BirthDateOn struct {
		Year  int `json:"year"`
		Month int `json:"month"`
		Day   int `json:"day"`
	} `json:"birthDateOn"`
	EmailAddress string `json:"emailAddress"`
	Address      string `json:"address"`
	EntityUrn    string `json:"entityUrn"`
	ConnectedAt  int64  `json:"connectedAt"`
	Websites     []struct {
		Type struct {
			MetaData struct {
				Category string `json:"category"`
			} `json:"com.linkedin.voyager.identity.profile.StandardWebsite"`
		} `json:"type"`
		URL string `json:"url"`
	} `json:"websites"`
	TwitterHandles []struct {
		Name         string `json:"name"`
		CredentialID string `json:"credentialId"`
	} `json:"twitterHandles"`
	PhoneNumbers []struct {
		Number string `json:"number"`
		Type   string `json:"type"`
	} `json:"phoneNumbers"`
}

type RecommendationView struct {
	Metadata struct {
		NumVisible int `json:"numVisible"`
	} `json:"metadata"`
	Elements []struct {
		EntityUrn   string `json:"entityUrn"`
		Created     int64  `json:"created"`
		Recommendee struct {
			FirstName        string `json:"firstName"`
			LastName         string `json:"lastName"`
			Occupation       string `json:"occupation"`
			ObjectUrn        string `json:"objectUrn"`
			EntityUrn        string `json:"entityUrn"`
			PublicIdentifier string `json:"publicIdentifier"`
			Picture          struct {
				ComLinkedinCommonVectorImage struct {
					Artifacts []struct {
						Width                         int    `json:"width"`
						FileIdentifyingURLPathSegment string `json:"fileIdentifyingUrlPathSegment"`
						ExpiresAt                     int64  `json:"expiresAt"`
						Height                        int    `json:"height"`
					} `json:"artifacts"`
					RootURL string `json:"rootUrl"`
				} `json:"com.linkedin.common.VectorImage"`
			} `json:"picture"`
			TrackingID string `json:"trackingId"`
		} `json:"recommendee"`
		RecommendationText             string `json:"recommendationText"`
		RecommendeeEntity              string `json:"recommendeeEntity"`
		VisibilityOnRecommenderProfile string `json:"visibilityOnRecommenderProfile"`
		LastModified                   int64  `json:"lastModified"`
		Relationship                   string `json:"relationship"`
		Recommender                    struct {
			FirstName        string `json:"firstName"`
			LastName         string `json:"lastName"`
			Occupation       string `json:"occupation"`
			ObjectUrn        string `json:"objectUrn"`
			EntityUrn        string `json:"entityUrn"`
			PublicIdentifier string `json:"publicIdentifier"`
			Picture          struct {
				ComLinkedinCommonVectorImage struct {
					Artifacts []struct {
						Width                         int    `json:"width"`
						FileIdentifyingURLPathSegment string `json:"fileIdentifyingUrlPathSegment"`
						ExpiresAt                     int64  `json:"expiresAt"`
						Height                        int    `json:"height"`
					} `json:"artifacts"`
					RootURL string `json:"rootUrl"`
				} `json:"com.linkedin.common.VectorImage"`
			} `json:"picture"`
			TrackingID string `json:"trackingId"`
		} `json:"recommender"`
		Status string `json:"status"`
	} `json:"elements"`
	Paging struct {
		Count int           `json:"count"`
		Start int           `json:"start"`
		Total int           `json:"total"`
		Links []interface{} `json:"links"`
	} `json:"paging"`
}

type InterestInfluencerView struct {
	Elements []struct {
		Entity struct {
			ComLinkedinVoyagerIdentitySharedMiniProfile struct {
				FirstName       string `json:"firstName"`
				LastName        string `json:"lastName"`
				Occupation      string `json:"occupation"`
				ObjectUrn       string `json:"objectUrn"`
				EntityUrn       string `json:"entityUrn"`
				BackgroundImage struct {
					ComLinkedinCommonVectorImage struct {
						Artifacts []struct {
							Width                         int    `json:"width"`
							FileIdentifyingURLPathSegment string `json:"fileIdentifyingUrlPathSegment"`
							ExpiresAt                     int64  `json:"expiresAt"`
							Height                        int    `json:"height"`
						} `json:"artifacts"`
						RootURL string `json:"rootUrl"`
					} `json:"com.linkedin.common.VectorImage"`
				} `json:"backgroundImage"`
				PublicIdentifier string `json:"publicIdentifier"`
				Picture          struct {
					ComLinkedinCommonVectorImage struct {
						Artifacts []struct {
							Width                         int    `json:"width"`
							FileIdentifyingURLPathSegment string `json:"fileIdentifyingUrlPathSegment"`
							ExpiresAt                     int64  `json:"expiresAt"`
							Height                        int    `json:"height"`
						} `json:"artifacts"`
						RootURL string `json:"rootUrl"`
					} `json:"com.linkedin.common.VectorImage"`
				} `json:"picture"`
				TrackingID string `json:"trackingId"`
			} `json:"com.linkedin.voyager.identity.shared.MiniProfile"`
		} `json:"entity"`
		FollowingInfo struct {
			EntityUrn     string `json:"entityUrn"`
			FollowerCount int    `json:"followerCount"`
			Following     bool   `json:"following"`
			TrackingUrn   string `json:"trackingUrn"`
		} `json:"followingInfo"`
	} `json:"elements"`
	Paging struct {
		Count int           `json:"count"`
		Start int           `json:"start"`
		Total int           `json:"total"`
		Links []interface{} `json:"links"`
	} `json:"paging"`
}

type InterestCompanyLoad struct {
	Elements []struct {
		Entity struct {
			ComLinkedinVoyagerEntitiesSharedMiniCompany struct {
				ObjectUrn string `json:"objectUrn"`
				EntityUrn string `json:"entityUrn"`
				Name      string `json:"name"`
				Showcase  bool   `json:"showcase"`
				Active    bool   `json:"active"`
				Logo      struct {
					ComLinkedinCommonVectorImage struct {
						Artifacts []struct {
							Width                         int    `json:"width"`
							FileIdentifyingURLPathSegment string `json:"fileIdentifyingUrlPathSegment"`
							ExpiresAt                     int64  `json:"expiresAt"`
							Height                        int    `json:"height"`
						} `json:"artifacts"`
						RootURL string `json:"rootUrl"`
					} `json:"com.linkedin.common.VectorImage"`
				} `json:"logo"`
				UniversalName string `json:"universalName"`
				TrackingID    string `json:"trackingId"`
			} `json:"com.linkedin.voyager.entities.shared.MiniCompany"`
		} `json:"entity"`
		FollowingInfo struct {
			EntityUrn     string `json:"entityUrn"`
			FollowerCount int    `json:"followerCount"`
			Following     bool   `json:"following"`
			TrackingUrn   string `json:"trackingUrn"`
		} `json:"followingInfo"`
	} `json:"elements"`
	Paging struct {
		Count int           `json:"count"`
		Start int           `json:"start"`
		Total int           `json:"total"`
		Links []interface{} `json:"links"`
	} `json:"paging"`
}

type InterestSchoolLoad struct {
	Elements []struct {
		Entity struct {
			ComLinkedinVoyagerEntitiesSharedMiniSchool struct {
				Active bool `json:"active"`
				Logo   struct {
					ComLinkedinCommonVectorImage struct {
						Artifacts []struct {
							Width                         int    `json:"width"`
							FileIdentifyingURLPathSegment string `json:"fileIdentifyingUrlPathSegment"`
							ExpiresAt                     int64  `json:"expiresAt"`
							Height                        int    `json:"height"`
						} `json:"artifacts"`
						RootURL string `json:"rootUrl"`
					} `json:"com.linkedin.common.VectorImage"`
				} `json:"logo"`
				ObjectUrn  string `json:"objectUrn"`
				SchoolName string `json:"schoolName"`
				EntityUrn  string `json:"entityUrn"`
				TrackingID string `json:"trackingId"`
			} `json:"com.linkedin.voyager.entities.shared.MiniSchool"`
		} `json:"entity"`
		FollowingInfo struct {
			EntityUrn     string `json:"entityUrn"`
			FollowerCount int    `json:"followerCount"`
			Following     bool   `json:"following"`
			TrackingUrn   string `json:"trackingUrn"`
		} `json:"followingInfo"`
	} `json:"elements"`
	Paging struct {
		Count int           `json:"count"`
		Start int           `json:"start"`
		Total int           `json:"total"`
		Links []interface{} `json:"links"`
	} `json:"paging"`
}

type InterestGroupLoad struct {
	Elements []struct {
		Entity struct {
			ComLinkedinVoyagerEntitiesSharedMiniGroup struct {
				GroupName        string `json:"groupName"`
				ObjectUrn        string `json:"objectUrn"`
				GroupDescription string `json:"groupDescription"`
				EntityUrn        string `json:"entityUrn"`
				Logo             struct {
					ComLinkedinCommonVectorImage struct {
						Artifacts []struct {
							Width                         int    `json:"width"`
							FileIdentifyingURLPathSegment string `json:"fileIdentifyingUrlPathSegment"`
							ExpiresAt                     int64  `json:"expiresAt"`
							Height                        int    `json:"height"`
						} `json:"artifacts"`
						RootURL string `json:"rootUrl"`
					} `json:"com.linkedin.common.VectorImage"`
				} `json:"logo"`
				TrackingID string `json:"trackingId"`
			} `json:"com.linkedin.voyager.entities.shared.MiniGroup"`
		} `json:"entity"`
		FollowingInfo struct {
			EntityUrn     string `json:"entityUrn"`
			FollowerCount int    `json:"followerCount"`
			Following     bool   `json:"following"`
			TrackingUrn   string `json:"trackingUrn"`
		} `json:"followingInfo"`
	} `json:"elements"`
	Paging struct {
		Count int           `json:"count"`
		Start int           `json:"start"`
		Total int           `json:"total"`
		Links []interface{} `json:"links"`
	} `json:"paging"`
}

type CertificationsView struct {
	Elements []CertificationElement `json:"elements"`
	Paging   Paging                 `json:"paging"`
}

type CertificationElement struct {
	EntityUrn     string     `json:"entityUrn"`
	Authority     string     `json:"authority"`
	Name          string     `json:"name"`
	TimePeriod    TimePeriod `json:"timePeriod"`
	LicenseNumber string     `json:"licenseNumber"`
	Company       Company    `json:"company"`
	DisplaySource string     `json:"displaySource"`
	CompanyUrn    string     `json:"companyUrn"`
	URL           string     `json:"url"`
}
