package main

import "fmt"

func main() {
	const IMCMaigreur = 18.5
	const IMCNormal = 25.0
	const IMCSurpoids = 30.0

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

	imc := poids / (taille * taille)

	fmt.Printf("\nBonjour %s !\n", nom)
	fmt.Printf("Poids : %.2f kg\n", poids)
	fmt.Printf("Taille : %.2f m\n", taille)
	fmt.Printf("IMC : %.2f\n", imc)

	if imc < IMCMaigreur {
		fmt.Println("Catégorie : Maigreur")
	} else if imc < IMCNormal {
		fmt.Println("Catégorie : Normal")
	} else if imc < IMCSurpoids {
		fmt.Println("Catégorie : Surpoids")
	} else {
		fmt.Println("Catégorie : Obésité")
	}
}