package sbom

// func (s *Spec) Validate() error {
// 	if err := s.validateVersion(); err != nil {
// 		return err
// 	}
// 	if err := s.validateFormat(); err != nil {
// 		return err
// 	}
// 	if err := s.validateSpecType(); err != nil {
// 		return err
// 	}
// 	if err := s.validateName(); err != nil {
// 		return err
// 	}
// 	if err := s.validateCreationTimestamp(); err != nil {
// 		return err
// 	}
// 	if err := s.validateSpdxID(); err != nil {
// 		return err
// 	}
// 	if err := s.validateNamespace(); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (s *Spec) validateVersion() error {
// 	if s.Version == "" {
// 		return errors.New("version cannot be empty")
// 	}
// 	return nil
// }

// func (s *Spec) validateFormat() error {
// 	if s.Format == "" {
// 		return errors.New("format cannot be empty")
// 	}
// 	return nil
// }

// func (s *Spec) validateSpecType() error {
// 	if s.SpecType == "" {
// 		return errors.New("spec type cannot be empty")
// 	}
// 	return nil
// }

// func (s *Spec) validateName() error {
// 	if s.Name == "" {
// 		return errors.New("name cannot be empty")
// 	}
// 	return nil
// }

// func (s *Spec) validateCreationTimestamp() error {
// 	_, err := time.Parse(time.RFC3339, s.CreationTimestamp)
// 	if err != nil {
// 		return errors.New("creation timestamp must be in RFC3339 format")
// 	}
// 	return nil
// }

// func (s *Spec) validateSpdxID() error {
// 	if s.SpdxID == "" {
// 		return errors.New("SPDX ID cannot be empty")
// 	}
// 	return nil
// }

// func (s *Spec) validateNamespace() error {
// 	if s.Namespace == "" {
// 		return errors.New("namespace cannot be empty")
// 	}
// 	// Example regex for validating a URL-like namespace
// 	regex := `^(http|https):\/\/[^\s]+$`
// 	matched, _ := regexp.MatchString(regex, s.Namespace)
// 	if !matched {
// 		return errors.New("namespace must be a valid URL")
// 	}
// 	return nil
// }
