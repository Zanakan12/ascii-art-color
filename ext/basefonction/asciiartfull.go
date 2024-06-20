package basefonction

import "strings"

func AsciiArtFull(input1,template,colortext, color string) []string{
	// Vérifie si l'input contient des sauts de ligne.
	linebreak := strings.Contains(input1, "\\n")
	// Initialise un tableau pour stocker les lignes de sortie.
	var outputLines []string
	if linebreak && len(input1) == 2 {
		// Si l'input ne contient qu'un saut de ligne, imprime une ligne vide.
	} else {
		if linebreak {
			// Si l'input contient des sauts de ligne, divise l'input en mots.
			words := strings.Split(input1, "\\n")
			for _, word := range words {
				if word == "" {
					// Ajoute une ligne vide à la sortie.
					outputLines = append(outputLines, "")
				} else {
					// Génère une ligne ASCII pour chaque mot et l'ajoute à la sortie.
					outputLines = append(outputLines, GenerateASCIILine(input1, template,"",""))
				}
			}
		} else {
			// Si l'input ne contient pas de saut de ligne, génère une seule ligne ASCII.
			outputLines = append(outputLines, GenerateASCIILine(input1, template,"",""))
		}
	}
	return outputLines
}
