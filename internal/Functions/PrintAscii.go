package BA

import (
	"errors"
	"strings"
)

func PrintAscii(input string, bannerMap map[rune][]string) (string, error) {
	var result strings.Builder

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			result.WriteString("\n")
			continue
		}

		for row := 0; row < 8; row++ {
			for _, ch := range line {
				if ch == '\r' {
					continue // To handle different operating systems
				}
				if art, ok := bannerMap[ch]; ok {
					result.WriteString(art[row])
				} else {
					return "", errors.New("unsupported character: '" + string(ch) + "'")
				}
			}
			result.WriteString("\n")
		}
	}

	return result.String(), nil
}
