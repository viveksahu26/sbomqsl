package sbom

import (
	"fmt"
	"io"

	cyclonedx "github.com/CycloneDX/cyclonedx-go"
)

type CdxDoc struct {
	CdxSpec *Spec
}

func (c CdxDoc) Spec() SpecInterface {
	return c.CdxSpec
}

func NewCdxDoc(f io.ReadSeeker) (Document, error) {
	bom := new(cyclonedx.BOM)
	decoder := cyclonedx.NewBOMDecoder(f, cyclonedx.BOMFileFormatJSON)
	if err := decoder.Decode(bom); err != nil {
		return nil, err
	}

	doc := &CdxDoc{}

	doc.parse()
	return doc, nil
}

func (c *CdxDoc) parse() {
	c.parseSpec()
}

func (c *CdxDoc) parseSpec() {
	spec := NewSpec()
	spec.Version = "1.0"
	spec.Format = "JSON"
	spec.Name = "Example Spec"
	spec.SpecType = "Type A"
	spec.CreationTimestamp = "2023-10-01T00:00:00Z"
	spec.Namespace = "http://example.namespace"
	spec.Comment = "This is an example spec."
	spec.SpdxID = "SPDX-1234"

	// if err := spec.Validate(); err != nil {
	// 	log.Fatalf("Validation failed: %v", err)
	// }

	fmt.Println("Version:", spec.GetVersion())
	fmt.Println("Format:", spec.GetFormat())
	fmt.Println("Name:", spec.GetName())
	fmt.Println("Spec Type:", spec.GetSpecType())
	fmt.Println("Creation Timestamp:", spec.GetCreationTimestamp())
	fmt.Println("Namespace:", spec.GetNamespace())
	fmt.Println("Comment:", spec.GetComment())
	fmt.Println("SPDX ID:", spec.GetSpdxID())

	c.CdxSpec = spec
}
