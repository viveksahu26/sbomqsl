package licenses

import (
	"embed"
	"encoding/json"
	"fmt"
)

var (
	// res variable is embedding the contents of files directory using FileSystem

	// go:embed files
	res      embed.FS
	licenses = map[string]string{
		"spdx":          "files/licenses_spdx.json",
		"spdxException": "files/licenses_spdx_exception.json",
		"aboutCode":     "files/licenses_aboutcode.json",
	}
)

type spdxLicense struct {
	Version    string
	Licenses   []spdxLicenseDetail
	Exceptions []spdxLicenseDetail
}

type spdxLicenseDetail struct {
	Reference          string   `json:"reference"`
	IsDeprecated       bool     `json:"isDeprecatedLicenseId"`
	DetailsURL         string   `json:"detailsUrl"`
	ReferenceNumber    int      `json:"referenceNumber"`
	Name               string   `json:"name"`
	LicenseID          string   `json:"licenseId"`
	LicenseExceptionId string   `json:"licenseExceptionId"`
	SeeAlso            []string `json:"seeAlso"`
	IsOsiApproved      bool     `json:"isOsiApproved"`
	IsFsfLibre         bool     `json:"isFsfLibre"`
}

var LicenseName = map[string]string{}

func LoadSpdxLicenses() error {
	licenseData, err := res.ReadFile(licenses["spdx"])
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return err
	}

	var sl spdxLicense
	if err := json.Unmarshal(licenseData, &sl); err != nil {
		fmt.Printf("error: %v\n", err)
		return err
	}
	fmt.Println("sl: ", sl.Licenses)

	for _, l := range sl.Licenses {
		LicenseName[l.LicenseID] = l.Name
	}

	return nil
}

func init() {
	LoadSpdxLicenses()
}
