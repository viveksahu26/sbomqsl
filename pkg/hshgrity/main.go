package main

import (
	"fmt"

	hshgrity "github.com/viveksahu26/sbomqsl/pkg/hshgrity/hsh"
)

func main() {
	// Provided hash from SBOM
	providedHash := "6adc31703e5ecc3486c1ea4854cac89286484363f5c3e9aab9ea8fe25efb2ffc"

	// Purl from SBOM
	// purl1 := "pkg:golang/github.com/CycloneDX/cyclonedx-go@v0.9.0"
	// purl2 := "pkg:pkg.go.dev/github.com/sigstore/cosign@v2.4.0"
	purl := "pkg:pkg.go.dev/github.com/interlynk-io/sbomqs@v0.1.9"
	// Get the calculated hashes from the upstream source
	calculatedHashes, err := hshgrity.GetCalculatedHash(purl)
	if err != nil {
		fmt.Println("Error getting calculated hashes:", err)
		return
	}

	// Check if the provided hash is in the list of calculated hashes
	hashFound := false
	for _, hash := range calculatedHashes {
		if providedHash == hash {
			hashFound = true
			break
		}
	}

	if hashFound {
		fmt.Println("The provided hash is present in the checksum file. The component has not been modified.")
	} else {
		fmt.Println("The provided hash is not present in the checksum file. The component may have been modified.")
	}
}
