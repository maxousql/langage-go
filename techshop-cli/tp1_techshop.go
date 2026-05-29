package main

import (
	"errors"
	"fmt"
	"strings"
)

type Produit struct {
	ID        int
	Nom       string
	Marque    string
	Prix      float64
	Stock     int
	Categorie string
	Actif     bool
}

type Catalogue struct {
	Produits []Produit
}

func (c *Catalogue) AjouterProduit(p Produit) error {
	if p.ID <= 0 {
		return errors.New("l'ID doit être supérieur à 0")
	}

	if p.Prix < 0 {
		return errors.New("le prix ne peut pas être négatif")
	}

	if p.Stock < 0 {
		return errors.New("le stock ne peut pas être négatif")
	}

	for _, produit := range c.Produits {
		if produit.ID == p.ID {
			return fmt.Errorf("un produit avec l'ID %d existe déjà", p.ID)
		}
	}

	c.Produits = append(c.Produits, p)
	return nil
}

func (c Catalogue) TrouverParID(id int) (Produit, error) {
	for _, produit := range c.Produits {
		if produit.ID == id {
			return produit, nil
		}
	}

	return Produit{}, fmt.Errorf("aucun produit trouvé avec l'ID %d", id)
}

func (c Catalogue) TrouverParCategorie(cat string) []Produit {
	var resultats []Produit

	for _, produit := range c.Produits {
		if produit.Actif && strings.EqualFold(produit.Categorie, cat) {
			resultats = append(resultats, produit)
		}
	}

	return resultats
}

func (c *Catalogue) AppliquerReduction(categorie string, pct float64) int {
	if pct <= 0 || pct > 100 {
		return 0
	}

	nbModifies := 0

	for i, produit := range c.Produits {
		if produit.Actif && strings.EqualFold(produit.Categorie, categorie) {
			reduction := c.Produits[i].Prix * pct / 100
			c.Produits[i].Prix -= reduction
			nbModifies++
		}
	}

	return nbModifies
}

func (c *Catalogue) Vendre(id int, qte int) error {
	if qte <= 0 {
		return errors.New("la quantité doit être supérieure à 0")
	}

	for i, produit := range c.Produits {
		if produit.ID == id {
			if !produit.Actif {
				return errors.New("ce produit n'est pas actif")
			}

			if produit.Stock < qte {
				return fmt.Errorf("stock insuffisant : stock disponible = %d", produit.Stock)
			}

			c.Produits[i].Stock -= qte
			return nil
		}
	}

	return fmt.Errorf("aucun produit trouvé avec l'ID %d", id)
}

func (c Catalogue) Rapport() string {
	nbProduits := 0
	stockTotal := 0
	valeurTotale := 0.0

	for _, produit := range c.Produits {
		if produit.Actif {
			nbProduits++
			stockTotal += produit.Stock
			valeurTotale += produit.Prix * float64(produit.Stock)
		}
	}

	return fmt.Sprintf(
		"Nombre de produits actifs : %d\nStock total : %d unités\nValeur totale du stock : %.2f €",
		nbProduits,
		stockTotal,
		valeurTotale,
	)
}

func afficherProduit(p Produit) {
	etat := "Inactif"
	if p.Actif {
		etat = "Actif"
	}

	fmt.Println("--------------------------------")
	fmt.Printf("ID : %d\n", p.ID)
	fmt.Printf("Nom : %s\n", p.Nom)
	fmt.Printf("Marque : %s\n", p.Marque)
	fmt.Printf("Prix : %.2f €\n", p.Prix)
	fmt.Printf("Stock : %d\n", p.Stock)
	fmt.Printf("Catégorie : %s\n", p.Categorie)
	fmt.Printf("État : %s\n", etat)
}

func afficherMenu() {
	fmt.Println("\n========== TECHSHOP CLI ==========")
	fmt.Println("[1] Ajouter un produit")
	fmt.Println("[2] Chercher un produit par ID")
	fmt.Println("[3] Soldes par catégorie")
	fmt.Println("[4] Vendre un produit")
	fmt.Println("[5] Rapport du catalogue")
	fmt.Println("[0] Quitter")
	fmt.Print("Votre choix : ")
}

func main() {
	catalogue := Catalogue{
		Produits: []Produit{
			{ID: 1, Nom: "iPhone_15", Marque: "Apple", Prix: 899.99, Stock: 8, Categorie: "Smartphone", Actif: true},
			{ID: 2, Nom: "Galaxy_S24", Marque: "Samsung", Prix: 799.99, Stock: 12, Categorie: "Smartphone", Actif: true},
			{ID: 3, Nom: "MacBook_Air_M2", Marque: "Apple", Prix: 1199.99, Stock: 5, Categorie: "Ordinateur", Actif: true},
			{ID: 4, Nom: "ThinkPad_X1", Marque: "Lenovo", Prix: 1399.99, Stock: 4, Categorie: "Ordinateur", Actif: true},
			{ID: 5, Nom: "MX_Master_3S", Marque: "Logitech", Prix: 99.99, Stock: 20, Categorie: "Accessoire", Actif: true},
		},
	}

	var choix int

	for {
		afficherMenu()

		_, err := fmt.Scan(&choix)
		if err != nil {
			fmt.Println("Erreur : vous devez entrer un nombre.")
			return
		}

		switch choix {
		case 1:
			var id int
			var nom string
			var marque string
			var prix float64
			var stock int
			var categorie string
			var actifInput int

			fmt.Println("\n=== Ajouter un produit ===")

			fmt.Print("ID : ")
			fmt.Scan(&id)

			fmt.Print("Nom du produit sans espace : ")
			fmt.Scan(&nom)

			fmt.Print("Marque sans espace : ")
			fmt.Scan(&marque)

			fmt.Print("Prix : ")
			fmt.Scan(&prix)

			fmt.Print("Stock : ")
			fmt.Scan(&stock)

			fmt.Print("Catégorie sans espace : ")
			fmt.Scan(&categorie)

			fmt.Print("Actif ? 1 = oui, 0 = non : ")
			fmt.Scan(&actifInput)

			produit := Produit{
				ID:        id,
				Nom:       nom,
				Marque:    marque,
				Prix:      prix,
				Stock:     stock,
				Categorie: categorie,
				Actif:     actifInput == 1,
			}

			err := catalogue.AjouterProduit(produit)
			if err != nil {
				fmt.Println("Erreur :", err)
			} else {
				fmt.Println("Produit ajouté avec succès.")
			}

		case 2:
			var id int

			fmt.Println("\n=== Chercher un produit par ID ===")
			fmt.Print("ID du produit : ")
			fmt.Scan(&id)

			produit, err := catalogue.TrouverParID(id)
			if err != nil {
				fmt.Println("Erreur :", err)
			} else {
				afficherProduit(produit)
			}

		case 3:
			var categorie string
			var reduction float64

			fmt.Println("\n=== Appliquer les soldes ===")
			fmt.Print("Catégorie : ")
			fmt.Scan(&categorie)

			fmt.Print("Pourcentage de réduction : ")
			fmt.Scan(&reduction)

			nb := catalogue.AppliquerReduction(categorie, reduction)

			if nb == 0 {
				fmt.Println("Aucun produit modifié. Vérifiez la catégorie ou le pourcentage.")
			} else {
				fmt.Printf("Réduction appliquée sur %d produit(s).\n", nb)
			}

		case 4:
			var id int
			var quantite int

			fmt.Println("\n=== Vendre un produit ===")
			fmt.Print("ID du produit : ")
			fmt.Scan(&id)

			fmt.Print("Quantité vendue : ")
			fmt.Scan(&quantite)

			err := catalogue.Vendre(id, quantite)
			if err != nil {
				fmt.Println("Erreur :", err)
			} else {
				fmt.Println("Vente effectuée avec succès.")
			}

		case 5:
			fmt.Println("\n=== Rapport du catalogue ===")
			fmt.Println(catalogue.Rapport())

		case 0:
			fmt.Println("Fermeture de TechShop CLI. À bientôt !")
			return

		default:
			fmt.Println("Choix invalide. Merci de choisir une option entre 0 et 5.")
		}
	}
}