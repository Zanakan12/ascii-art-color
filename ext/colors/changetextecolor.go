package colors

func ChangeTextColor(words, color string) string {
	str := ""
	for _, line := range words {
		str += ApplyColor(string(line), color)
	}
	return str
}
