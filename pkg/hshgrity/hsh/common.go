package hshgrity

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// CalculateHash calculates the SHA256 hash of a file
func CalculateHash(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", fmt.Errorf("error calculating hash: %v", err)
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// GetCalculatedHash constructs the URL for the checksum file and retrieves the hash
func GetCalculatedHash(purl string) ([]string, error) {
	// Extract information from purl
	purlParts := strings.Split(purl, "/")
	if len(purlParts) < 4 {
		return nil, fmt.Errorf("invalid purl format")
	}
	fmt.Println("purlParts: ", purlParts)

	typeParts := strings.Split(purlParts[0], ":")
	fmt.Println("repoParts: ", typeParts)
	if len(typeParts) < 2 {
		return nil, fmt.Errorf("invalid purl format")
	}
	org := purlParts[2]
	// repo := repoParts[2]

	repoParts := strings.Split(purlParts[3], "@")
	if len(repoParts) < 2 {
		return nil, fmt.Errorf("invalid purl format")
	}
	repo := repoParts[0]
	version := repoParts[1]

	// Remove the 'v' prefix from the version if it exists
	if strings.HasPrefix(version, "v") {
		version = version[1:]
	}

	// Construct the URL for the checksum file
	versionedChecksumURL := fmt.Sprintf("https://github.com/%s/%s/releases/download/v%s/%s_%s_checksums.txt", org, repo, version, repo, version)
	fmt.Println("versionedChecksumURL : ", versionedChecksumURL)

	// Attempt to download the versioned checksum file
	hashes, err := downloadChecksumFile(versionedChecksumURL)
	if err == nil {
		return hashes, nil
	}

	// If the versioned checksum file is not found, construct the URL for the general checksum file
	generalChecksumURL := fmt.Sprintf("https://github.com/%s/%s/releases/download/v%s/%s_checksums.txt", org, repo, version, repo)
	fmt.Println("generalChecksumURL: ", generalChecksumURL)

	// Attempt to download the general checksum file
	hashes, err = downloadChecksumFile(generalChecksumURL)
	if err != nil {
		return nil, fmt.Errorf("error downloading checksum file: %v", err)
	}

	return hashes, nil
}

// downloadChecksumFile downloads the checksum file from the given URL and extracts all hashes
func downloadChecksumFile(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error downloading checksum file: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("checksum file not found at URL: %s", url)
	}

	var hashes []string
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) > 0 {
			hashes = append(hashes, parts[0])
		}
	}
	fmt.Println("hashes: ", hashes)

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading checksum file: %v", err)
	}

	return hashes, nil
}
