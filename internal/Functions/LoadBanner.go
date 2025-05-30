package BA

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(banner string) (map[rune][]string, error) {
	var bannerPath string
	var lines []string
	var rawURL string

	// Adjust style path according to chosen style
	switch banner {
	case "standard":
		bannerPath = "/internal/banners/standard.txt"
		rawURL = "https://raw.githubusercontent.com/first22basel/ASCII-ART-WEB-BA/main/internal/banners/standard.txt"
	case "shadow":
		bannerPath = "/internal/banners/shadow.txt"
		rawURL = "https://raw.githubusercontent.com/first22basel/ASCII-ART-WEB-BA/main/internal/banners/shadow.txt"
	case "thinkertoy":
		bannerPath = "/internal/banners/thinkertoy.txt"
		rawURL = "https://raw.githubusercontent.com/first22basel/ASCII-ART-WEB-BA/main/internal/banners/thinkertoy.txt"
	default:
		return nil, fmt.Errorf("invalid bannner name: %s", banner)
	}

	// Check if banner file is exist
	err := EnsureFile(bannerPath, rawURL)
	if err != nil {
		return nil, fmt.Errorf("%v.txt file is not found: %v", banner, err)
	}

	// Read data from banner file
	data, err := os.ReadFile(bannerPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read %v.txt", banner)
	}

	lines = strings.Split(string(data), "\n")
	fontMap := make(map[rune][]string)
	startChar := 32
	for i := 1; i+7 < len(lines); i += 9 {
		char := rune(startChar)
		fontMap[char] = lines[i : i+8]
		startChar++
	}

	return fontMap, nil
}
