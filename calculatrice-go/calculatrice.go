package main

import (
	"errors"
	"fmt"
)

func creerOperation(op string) func(float64, float64) float64 {
	switch op {
	case "+":
		return func(a, b float64) float64 {
			return a + b
		}
	case "-":
		return func(a, b float64) float64 {
			return a - b
		}
	case "*":
		return func(a, b float64) float64 {
			return a * b
		}
	case "/":
		return func(a, b float64) float64 {
			return a / b
		}
	default:
		return nil
	}
}

func operer(a, b float64, op string) (float64, error) {
	if op == "/" && b == 0 {
		return 0, errors.New("division par zéro impossible")
	}

	operation := creerOperation(op)

	if operation == nil {
		return 0, errors.New("opération inconnue")
	}

	resultat := operation(a, b)

	return resultat, nil
}

func main() {
	var a float64
	var b float64
	var op string

	fmt.Println("=== Calculatrice CLI ===")
	fmt.Println("Opérations disponibles : +, -, *, /")
	fmt.Println("Pour quitter, tapez : 0 0 quit")
	fmt.Println()

	for {
		fmt.Print("Entrez deux nombres et une opération : ")

		_, err := fmt.Scan(&a, &b, &op)
		if err != nil {
			fmt.Println("Erreur : saisie invalide.")
			return
		}

		if op == "quit" {
			fmt.Println("Fin du programme.")
			break
		}

		resultat, err := operer(a, b, op)
		if err != nil {
			fmt.Println("Erreur :", err)
		} else {
			fmt.Printf("Résultat : %.2f\n", resultat)
		}

		fmt.Println()
	}
}