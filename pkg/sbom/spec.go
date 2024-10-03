package sbom

// Spec interface defines a set of methods that any type implementing this interface must provide
type SpecInterface interface {
	GetVersion() string
	GetFormat() string
	GetSpecType() string
	GetName() string
	GetCreationTimestamp() string
	GetSpdxID() string
	GetComment() string
	GetNamespace() string
}

// The Specs struct is a concrete implementation of the Spec interface.
// It contains fields that store the data for each method defined in the Spec interface.
type Spec struct {
	Version           string
	Format            string
	SpecType          string
	Name              string
	CreationTimestamp string
	SpdxID            string
	Comment           string
	Namespace         string
}

func NewSpec() *Spec {
	return &Spec{}
}


func (s Spec) GetVersion() string {
	return s.Version
}

func (s Spec) GetFormat() string {
	return s.Format
}

func (s Spec) GetSpecType() string {
	return s.SpecType
}

func (s Spec) GetName() string {
	return s.Name
}

func (s Spec) GetCreationTimestamp() string {
	return s.CreationTimestamp
}

func (s Spec) GetSpdxID() string {
	return s.SpdxID
}

func (s Spec) GetComment() string {
	return s.Comment
}

func (s Spec) GetNamespace() string {
	return s.Namespace
}

