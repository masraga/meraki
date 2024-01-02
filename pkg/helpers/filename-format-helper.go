package helpers

import (
	"fmt"
	"strings"
)

func FilenameFormatHelper(file string, extenstion string) string {
	spellText := strings.Split(file, "")
	var fileName []string
	for i, r := range spellText {
		if r >= "A" && r <= "Z" {
			if i == 0 {
				fileName = append(fileName, strings.ToLower(r))
			} else {
				fileName = append(fileName, fmt.Sprintf("-%s", strings.ToLower(r)))
			}
		} else {
			fileName = append(fileName, strings.ToLower(r))
		}
	}
	return fmt.Sprintf("%s.%s", strings.Join(fileName, ""), extenstion)
}
