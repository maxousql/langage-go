package main

import "fmt"

const (
	IMCMaigreur  = 18.5
	IMCNormal   = 25.0
	IMCSurpoids = 30.0
)

func calculerIMC(poids float64, taille float64) float64 {
	return poids / (taille * taille)
}

func determinerCategorie(imc float64) string {
	switch {
	case imc < IMCMaigreur:
		return "Maigreur"
	case imc < IMCNormal:
		return "Normal"
	case imc < IMCSurpoids:
		return "Surpoids"
	default:
		return "Obésité"
	}
}

func afficherResultat(nom string, poids float64, taille float64, imc float64, categorie string) {
	fmt.Printf("\nBonjour %s !\n", nom)
	fmt.Printf("Poids : %.2f kg\n", poids)
	fmt.Printf("Taille : %.2f m\n", taille)
	fmt.Printf("IMC : %.2f\n", imc)
	fmt.Printf("Catégorie : %s\n", categorie)
}

func main() {
	var nom string
	var poids float64
	var taille float64

	fmt.Print("Entrez votre prénom : ")
	fmt.Scanln(&nom)

	fmt.Print("Entrez votre poids en kg : ")
	fmt.Scanln(&poids)

	fmt.Print("Entrez votre taille en mètres : ")
	fmt.Scanln(&taille)

	if poids <= 0 || taille <= 0 {
		fmt.Println("Erreur : le poids et la taille doivent être supérieurs à 0.")
		return
	}

	imc := calculerIMC(poids, taille)
	categorie := determinerCategorie(imc)

	afficherResultat(nom, poids, taille, imc, categorie)
}