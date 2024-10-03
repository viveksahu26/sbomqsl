package hshgrity

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDownloadChecksumFile(t *testing.T) {
	x := "6998cd617b528dd594a0336bf2c202250b45ea7f06c2be1fe6a5c102ef440e1e  cyclonedx-go-0.9.1.tar.gz"
	y := "3616b75ed31fb61cc0d88b21be052d87f13cb8e93e4bad53dd7edb9a2bac06c9  cyclonedx-go-0.9.1.tar.gz.cdx.sbom"
	tests := []struct {
		name           string
		responseStatus int
		responseBody   string
		expectedHashes []string
		expectError    bool
	}{
		{
			name:           "validChecksumFile",
			responseStatus: http.StatusOK,
			responseBody:   x + "\n" + y + "\n",
			expectedHashes: []string{"6998cd617b528dd594a0336bf2c202250b45ea7f06c2be1fe6a5c102ef440e1e", "3616b75ed31fb61cc0d88b21be052d87f13cb8e93e4bad53dd7edb9a2bac06c9"},
			expectError:    false,
		},
		{
			name:           "checkSumFileNotFound",
			responseStatus: http.StatusNotFound,
			responseBody:   "",
			expectedHashes: nil,
			expectError:    true,
		},
		{
			name:           "emptyChecksumFile",
			responseStatus: http.StatusOK,
			responseBody:   "",
			expectedHashes: []string{},
			expectError:    false,
		},
		{
			name:           "malformedChecksumFile",
			responseStatus: http.StatusOK,
			responseBody:   "abc123\nxyz789  file2\n",
			expectedHashes: []string{"abc123", "xyz789"},
			expectError:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock server
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.responseStatus)
				w.Write([]byte(tt.responseBody))
			}))
			defer server.Close()

			// Call the function with the mock server URL
			hashes, err := downloadChecksumFile(server.URL)

			// Check for expected error
			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			}

			// Check for expected hashes
			if !equal(hashes, tt.expectedHashes) {
				t.Errorf("expected hashes %v, but got %v", tt.expectedHashes, hashes)
			}
		})
	}
}

// Helper function to compare two slices
func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// func TestGetCalculatedHash(t *testing.T) {
// 	// Define checksum strings
// 	x := "6998cd617b528dd594a0336bf2c202250b45ea7f06c2be1fe6a5c102ef440e1e  cyclonedx-go-0.9.1.tar.gz"
// 	y := "3616b75ed31fb61cc0d88b21be052d87f13cb8e93e4bad53dd7edb9a2bac06c9  cyclonedx-go-0.9.1.tar.gz.cdx.sbom"

// 	tests := []struct {
// 		name           string
// 		purl           string
// 		responseStatus int
// 		responseBody   string
// 		expectedHashes []string
// 		expectError    bool
// 	}{
// 		{
// 			name:           "Valid versioned checksum file",
// 			purl:           "pkg:golang/github.com/CycloneDX/cyclonedx-go@v0.9.0",
// 			responseStatus: http.StatusOK,
// 			responseBody:   x + "\n" + y + "\n",
// 			expectedHashes: []string{"6998cd617b528dd594a0336bf2c202250b45ea7f06c2be1fe6a5c102ef440e1e", "3616b75ed31fb61cc0d88b21be052d87f13cb8e93e4bad53dd7edb9a2bac06c9"},
// 			expectError:    false,
// 		},
// 		{
// 			name:           "Valid general checksum file",
// 			purl:           "pkg:golang/github.com/CycloneDX/cyclonedx-go@v0.9.0",
// 			responseStatus: http.StatusNotFound,
// 			responseBody:   "",
// 			expectedHashes: nil,
// 			expectError:    true,
// 		},
// 		{
// 			name:           "Checksum file not found",
// 			purl:           "pkg:golang/github.com/CycloneDX/cyclonedx-go@v0.9.0",
// 			responseStatus: http.StatusNotFound,
// 			responseBody:   "",
// 			expectedHashes: nil,
// 			expectError:    true,
// 		},
// 		{
// 			name:           "Empty checksum file",
// 			purl:           "pkg:golang/github.com/CycloneDX/cyclonedx-go@v0.9.0",
// 			responseStatus: http.StatusOK,
// 			responseBody:   "",
// 			expectedHashes: []string{},
// 			expectError:    false,
// 		},
// 		{
// 			name:           "Malformed checksum file",
// 			purl:           "pkg:golang/github.com/CycloneDX/cyclonedx-go@v0.9.0",
// 			responseStatus: http.StatusOK,
// 			responseBody:   "abc123\nxyz789  file2\n",
// 			expectedHashes: []string{"abc123", "xyz789"},
// 			expectError:    false,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// Create a mock server
// 			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 				w.WriteHeader(tt.responseStatus)
// 				w.Write([]byte(tt.responseBody))
// 			}))
// 			defer server.Close()

// 			// Replace the URL construction in GetCalculatedHash to use the mock server URL
// 			// originalDownloadChecksumFile := downloadChecksumFile
// 			// downloadChecksumFile = func(url string) ([]string, error) {
// 			// 	return originalDownloadChecksumFile(server.URL)
// 			// }
// 			// defer func() { downloadChecksumFile = originalDownloadChecksumFile }()

// 			// Call the function with the test purl
// 			hashes, err := GetCalculatedHash(tt.purl)

// 			// Check for expected error
// 			if tt.expectError {
// 				if err == nil {
// 					t.Errorf("expected error but got none")
// 				}
// 			} else {
// 				if err != nil {
// 					t.Errorf("unexpected error: %v", err)
// 				}
// 			}

// 			// Check for expected hashes
// 			if !equal(hashes, tt.expectedHashes) {
// 				t.Errorf("expected hashes %v, but got %v", tt.expectedHashes, hashes)
// 			}
// 		})
// 	}
// }
