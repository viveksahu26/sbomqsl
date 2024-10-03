package main

import (
	"fmt"
	"os"

	"github.com/caarlos0/log"
	"github.com/viveksahu26/sbomqsl/pkg/sbom"
)

func main() {
	// // if err := spec.Validate(); err != nil {
	// // 	log.Fatalf("Validation failed: %v", err)
	// // }

	// path := "/home/linuzz/sbom/sbomqs-fossa.cyclonedx.json"
	path := "/home/linuzz/sbom/sbomqs-fossa.spdx.json"
	f, err := os.Open(path)
	if err != nil {
		log.Debugf("os.Open failed for file :%s\n", path)
		fmt.Printf("failed to open %s\n", path)
		// return nil, nil, err
	}
	defer f.Close()

	// doc, err := sbom.NewCdxDoc(f)
	// if err != nil {
	// 	fmt.Println("failed to create new cdx doc: ", err)
	// }
	// fmt.Println("doc : ", doc)

	doc, err := sbom.NewSpdxDoc(f)
	if err != nil {
		fmt.Println("failed to create new cdx doc: ", err)
	}
	fmt.Println("doc : ", doc)
}
