package basefonction

import (
	"bufio"
	"fmt"
	"os"
)

// /////////////////////////////////////////////////FORMAT DE TEXTE A UTILSER//////////////////////////////////////////////////////////////
func GetLineAllFroma(lineNumber int, template string) string {
	file, err := os.Open(template)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier:", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan() && i < lineNumber; i++ {
		// Avance jusqu'à la ligne désirée
	}
	return scanner.Text()

}
