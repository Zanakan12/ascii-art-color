package main

import (
	"fmt"
	"os"

	"reverse/ext/basefonction"
	"reverse/ext/output"
	"reverse/ext/reverse"
	"strings"
)

func main() {
	/////////////////////////////////////////////VERIFICATION OF ARGUMENT/////////////////////////////////////////////////////////
	// Vérifie les arguments fournis au programme. Si la vérification échoue, le programme s'arrête.
	if !basefonction.Verification() {
		return
	}
	filenameorcolor, template, input1, option := basefonction.Func()
	/////////////////////////////////////////////TRAITEMENT DES ARGUMENTS/////////////////////////////////////////////////////////
	// Récupère les arguments nécessaires pour le traitement.
	outputLines := basefonction.AsciiArtFull(input1, template, "", "")
	// Si un fichier de sortie est spécifié, écrit le résultat dans ce fichier.
	if option == "--output" {
		output.OutPut(filenameorcolor, outputLines)
	} else if option == "--reverse" {
		reverse.Reverse(filenameorcolor, template)
		fmt.Println("Fonction reverse en cours de maintenance")
	} else if option == "--color" {
		tocolor := ""
		if len(os.Args) == 4 {
			tocolor = os.Args[2]
			input1 = os.Args[3]
		} else {
			tocolor = ""
		}
		if !strings.Contains(input1,tocolor ) {
			fmt.Println("Your letters are not in the character string to be colored orpoorly written command.\n Usage : $ go run . ")
			return
		}
		fmt.Println(basefonction.GenerateASCIILine(input1, template, filenameorcolor, tocolor))
	} else if option == "--align" {
		fmt.Println("Fonction align en cours de maintenance")
	} else {
		// Imprime chaque ligne de sortie.
		for _, line := range outputLines {
			fmt.Println(line)
		}
	}
}
