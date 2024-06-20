package colors

func ChangeColorChar(word string, charsToChange map[rune]bool, color string) string {
	result := ""
	for _, char := range word {
		if charsToChange[char] {
			result += ApplyColor(string(char), color)
		} else {
			result += string(char)
		}
	}
	return result
}
