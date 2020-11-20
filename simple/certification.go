package simple

type Certification struct {
	DateRange       DateRange `json:"dateRange"`
	URL             string    `json:"url"`
	Authority       string    `json:"authority"`
	CertificateName string    `json:"certificateName"`
	LicenseNumber   string    `json:"licenseNumber"`
	Company         *Company  `json:"company"`
	Source          string    `json:"source"`
}
