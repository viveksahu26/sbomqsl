package sbom

// A sbom document has:
// components, author, relationship, tools, spec
type Document interface {
	Spec() SpecInterface
}
