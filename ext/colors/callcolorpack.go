package colors



func CallColorPack(words, selectedChars, color string) string {
    if selectedChars == "" {
        return ChangeTextColor(words, color)
    } else {
        charsToChange := make(map[rune]bool)
        for _, char := range selectedChars {
            charsToChange[char] = true
        }
        return ChangeColorChar(words, charsToChange, color)
    }
}
