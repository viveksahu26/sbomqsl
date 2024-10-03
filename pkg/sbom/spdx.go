package sbom

import (
	"fmt"
	"io"

	spdx_json "github.com/spdx/tools-golang/json"
	"github.com/spdx/tools-golang/spdx"
)

type SpdxDoc struct {
	SpdxSpec *Spec
}

func (s SpdxDoc) Spec() SpecInterface {
	return s.SpdxSpec
}

func NewSpdxDoc(f io.ReadSeeker) (Document, error) {
	var d *spdx.Document
	var err error

	d, err = spdx_json.Read(f)
	if err != nil {
		fmt.Errorf("unsupported spdx format %s", string("json"))
	}
	fmt.Println("d: ", d)
	doc := &SpdxDoc{}

	doc.parse()
	return doc, nil
}

func (s *SpdxDoc) parse() {
	s.parseSpec()
}

func (s *SpdxDoc) parseSpec() {
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

	s.SpdxSpec = spec
}
