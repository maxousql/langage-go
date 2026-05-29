package main

import "fmt"

type Personne struct {
	Prenom string
	Nom    string
	Age    int
	Email  string
}

func (p Personne) NomComplet() string {
	return fmt.Sprintf("%s %s", p.Prenom, p.Nom)
}

func (p Personne) Presentation() string {
	return fmt.Sprintf(
		"Nom complet : %s\nÂge : %d ans\nEmail : %s",
		p.NomComplet(),
		p.Age,
		p.Email,
	)
}

type Adresse struct {
	Rue        string
	Ville      string
	CodePostal string
}

func (a Adresse) Format() string {
	return fmt.Sprintf("%s, %s %s", a.Rue, a.CodePostal, a.Ville)
}

type Employe struct {
	Personne
	Adresse
	Poste   string
	Salaire float64
}

func (e Employe) FicheEmploye() string {
	return fmt.Sprintf(
		"=== Fiche Employé ===\n%s\nAdresse : %s\nPoste : %s\nSalaire : %.2f €",
		e.Presentation(),
		e.Adresse.Format(),
		e.Poste,
		e.Salaire,
	)
}

func (e *Employe) AugmenterSalaire(pct float64) {
	e.Salaire += e.Salaire * pct / 100
}

type Etudiant struct {
	Personne
	Promo   string
	Moyenne float64
}

func (e Etudiant) MentionObtenue() string {
	switch {
	case e.Moyenne >= 16:
		return "Très Bien"
	case e.Moyenne >= 14:
		return "Bien"
	case e.Moyenne >= 12:
		return "Assez Bien"
	case e.Moyenne >= 10:
		return "Passable"
	default:
		return "Non admis"
	}
}

func (e Etudiant) FicheEtudiant() string {
	return fmt.Sprintf(
		"=== Fiche Étudiant ===\n%s\nPromo : %s\nMoyenne : %.2f\nMention : %s",
		e.Presentation(),
		e.Promo,
		e.Moyenne,
		e.MentionObtenue(),
	)
}

func main() {
	employe1 := Employe{
		Personne: Personne{
			Prenom: "Maxime",
			Nom:    "Lai Yio",
			Age:    23,
			Email:  "maxime@example.com",
		},
		Adresse: Adresse{
			Rue:        "12 rue des Lilas",
			Ville:      "Strasbourg",
			CodePostal: "67000",
		},
		Poste:   "Développeur Go",
		Salaire: 2800.00,
	}

	employe2 := Employe{
		Personne: Personne{
			Prenom: "Sarah",
			Nom:    "Martin",
			Age:    31,
			Email:  "sarah.martin@example.com",
		},
		Adresse: Adresse{
			Rue:        "8 avenue Victor Hugo",
			Ville:      "Lyon",
			CodePostal: "69000",
		},
		Poste:   "Cheffe de projet",
		Salaire: 3500.00,
	}

	etudiant1 := Etudiant{
		Personne: Personne{
			Prenom: "Lucas",
			Nom:    "Bernard",
			Age:    20,
			Email:  "lucas.bernard@example.com",
		},
		Promo:   "Bachelor Informatique",
		Moyenne: 15.5,
	}

	etudiant2 := Etudiant{
		Personne: Personne{
			Prenom: "Emma",
			Nom:    "Durand",
			Age:    22,
			Email:  "emma.durand@example.com",
		},
		Promo:   "Master SI",
		Moyenne: 17.2,
	}

	employe1.AugmenterSalaire(10)

	fmt.Println(employe1.FicheEmploye())
	fmt.Println()

	fmt.Println(employe2.FicheEmploye())
	fmt.Println()

	fmt.Println(etudiant1.FicheEtudiant())
	fmt.Println()

	fmt.Println(etudiant2.FicheEtudiant())
}