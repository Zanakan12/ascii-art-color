package basefonction

import (
	"fmt"
	"os"
	"reverse/ext/colors"
	"strings"
)

func Verification() bool {
	// Vérifie si le nombre d'arguments est incorrect
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("Usage: \n$ go run main.go <Printable characters> \n$ go run main.go <Printable characters> <template> ")
		fmt.Println("$ go run main.go <Printable characters> <template>  <Option>\nTemplate available : shadow, thinkertoy\nOption : --cololr=<color>, --output=<banner.txt> --align=<text alignment> ")
		return false
	} else {
		var tabasc []rune
		// Génère la table ASCII pour les caractères imprimables (de 32 à 126)
		for i := 32; i <= 126; i++ {
			tabasc = append(tabasc, rune(i))
		}
		// Vérifie si chaque caractère de l'argument fourni est imprimable
		for _, v := range os.Args[1] {
			if !strings.ContainsRune(string(tabasc), v) {
				// Affiche un message d'erreur si un caractère non imprimable est trouvé
				fmt.Printf("'%c' is an unprintable character\n", v)
				fmt.Println("Your input has unprintable characters.")
				return false
			}
		}
		filenameorcolor, template, input1, option := Func()
		if filenameorcolor == "" && template == "" && input1 == "" && option == "" {
			fmt.Println(colors.ApplyColor("Error: Extra argument detected. Please check the provided parameters and try again.", "red"), "\n", colors.ApplyColor("Manual: $ go run .", "blue"))
			return false
		}
		// Retourne vrai si tous les caractères sont imprimables et que le nombre d'arguments est correct
		return true
	}
}
