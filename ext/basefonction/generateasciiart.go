package basefonction

import (
	"reverse/ext/colors"
	"strings"
)

// //////////////////////////////////////////////COMBINAISON DES ASCIIART////////////////////////////////////////////////////////////////
// Génère une ligne d'art ASCII à partir du texte fourni
func GenerateASCIILine(text, template, color, targetColor string) string {
	result := ""
	for i := 0; i < 8; i++ { // Loop for each line of the ASCII art (8 lines per character)
		j := 0
		for j < len(text) {
			letter := rune(text[j])
			lineIndex := int(letter-' ')*9 + i + 1
			line := GetLineAllFroma(lineIndex, template)
			if color != "" && targetColor != "" {
				if strings.HasPrefix(text[j:], targetColor) {
					for k := 0; k < len(targetColor); k++ {
						letter = rune(text[j+k])
						lineIndex = int(letter-' ')*9 + i + 1
						line = GetLineAllFroma(lineIndex, template)
						result += colors.CallColorPack(line, "", color)
					}
					j += len(targetColor) - 1 // Skip the length of the targetColor string
				} else {
					result += line
				}
			} else if color != "" {
				result += colors.CallColorPack(line, "", color)
			} else {
				result += line
			}
			j++
		}
		result += "\n"
	}
	return result
}
