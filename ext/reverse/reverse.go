package reverse

import (
	"bufio"
	"fmt"
	"os"
)

func Reverse(filepath, template string) {
	// Ouvrir le fichier ASCII art
	file, err := os.Open(filepath) // Assurez-vous que le chemin d'accès au fichier est correct
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier:", err)
		return
	}
	defer file.Close()

	// Lire le fichier ligne par ligne
	scanner := bufio.NewScanner(file)
	var texteASCII []string
	for scanner.Scan() {
		texteASCII = append(texteASCII, scanner.Text())
	}
	fmt.Println(texteASCII)
	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}

	// Calculer le nombre de colonnes (longueur d'une ligne)
	if len(texteASCII) == 0 {
		return // Quitter si le tableau est vide
	}
	nombreColonnes := len(texteASCII[0])

	// Initialiser un tableau pour stocker les résultats
	resultat := make([][]rune, nombreColonnes)
	for i := range resultat {
		resultat[i] = make([]rune, len(texteASCII))
	}

	// Remplir le tableau avec les caractères verticalement
	for col := 0; col < nombreColonnes; col++ {
		for idx, ligne := range texteASCII {
			if col < len(ligne) {
				resultat[col][idx] = rune(ligne[col])

			} else {
				resultat[col][idx] = ' ' // Remplir avec des espaces si la colonne n'existe pas dans la ligne
			}
		}
	}

	// Afficher le tableau ligne par ligne
	result := ""
	for _, colonne := range resultat {
		count := 0
		for _, caractere := range colonne {
			if caractere == ' ' {
				count++
				result += "*"
			} else {
				result += string(caractere)
			}
			//fmt.Printf("%c", caractere)
		}
		if count == 8 {
			result += "\n"
		}
	}
	fmt.Println(result)
}
