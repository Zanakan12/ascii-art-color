package basefonction

import (
	"os"
	"strings"
)

// /////////////////////////////////////////////////////CALL FONCTION ///////////////////////////////////////////////////////////////////////////////////
func Func() (string, string, string, string) {
	// Initialisation des variables avec des valeurs par défaut
	filename := "" 
	template := "./Template/standard.txt" // Chemin par défaut du template
	str := os.Args[2]
	args, option := os.Args[1:], os.Args[1] // Arguments passés au programme
	check := true
	// Si seulement un argument est passé
	if len(os.Args) == 1 {
		template = "./Template/standard.txt" 
		str = os.Args[1] 
	} else {
		// Boucle à travers les arguments pour traiter chacun d'eux
		for _, v := range args {
			switch {
			// Sélection du template basé sur l'argument
			case v == "special" :
				check = false
				template = "./Template/special.txt"
			case v == "standard" :
				check = false 
				template = "./Template/standard.txt"
			case v == "shadow":
				check = false 
				template = "./Template/shadow.txt"
			case v == "thinkertoy":
				check = false 
				template = "./Template/thinkertoy.txt"
			// Si l'argument spécifie un fichier de sortie
			case strings.HasPrefix(v, "--output="):
				filename = strings.TrimPrefix(v, "--output=")
			case strings.HasPrefix(v, "--reverse="):
				filename = strings.TrimPrefix(v, "--reverse=")
			case strings.HasPrefix(v, "--color="):
				filename = strings.TrimPrefix(v, "--color=")
			case strings.HasPrefix(v, "--align="):
				filename = strings.TrimPrefix(v, "--align=")
			// Par défaut, l'argument est considéré comme la chaîne de caractères à traiter
			default:
				str = v
			}
			// Gestion des options
			if v != "standard" && v != "thinkertoy" && v != "shadow" {
				if strings.Contains(v, "--output=") || strings.Contains(v, "--align=") || strings.Contains(v, "--color=")||strings.Contains(v, "--reverse=") {
					option = strings.TrimSuffix(v,"="+filename)
					return filename,template,str,option 
				}
			}else if !check&&len(os.Args)>4  {
				return "","","",""
			}
		}
	}
	// Retourne le nom du fichier, le template, la chaîne de caractères et l'option
	return filename, template, str, option
}
