package main

import "fmt"

const (
	RangCommun = iota
	RangRare
	RangEpique
	RangLegendaire
)

var nomsRangs = map[int]string{
	RangCommun:     "Commun",
	RangRare:       "Rare",
	RangEpique:     "Épique",
	RangLegendaire: "Légendaire",
}

func calculerRang(points int) int {
	switch {
	case points >= 250:
		return RangLegendaire
	case points >= 180:
		return RangEpique
	case points >= 100:
		return RangRare
	default:
		return RangCommun
	}
}

func afficherBonus(rang int) {
	fmt.Println("\nBonus débloqués :")

	switch rang {
	case RangLegendaire:
		fmt.Println("- Couronne du maître explorateur")
		fallthrough
	case RangEpique:
		fmt.Println("- Coffre épique")
		fallthrough
	case RangRare:
		fmt.Println("- Badge rare")
		fallthrough
	case RangCommun:
		fmt.Println("- Badge de participation")
	default:
		fmt.Println("- Aucun bonus")
	}
}

func main() {
	var pseudo string

	fmt.Println("=== L'expédition des cristaux ===")
	fmt.Print("Entrez le nom de votre explorateur : ")
	fmt.Scan(&pseudo)

	zones := [5]string{
		"Forêt des Brumes",
		"Mine Oubliée",
		"Temple Solaire",
		"Lac Gelé",
		"Volcan Rouge",
	}

	zonesExplorees := zones[:]

	fmt.Printf("\nZones disponibles : %v\n", zonesExplorees)
	fmt.Printf("Nombre de zones avec len() : %d\n", len(zonesExplorees))
	fmt.Printf("Capacité de la slice avec cap() : %d\n", cap(zonesExplorees))

	objetsParZone := map[string]string{
		"Forêt des Brumes": "Cristal",
		"Mine Oubliée":    "Rune",
		"Temple Solaire":  "Relique",
		"Lac Gelé":        "Perle",
		"Volcan Rouge":    "Couronne",
	}

	pointsParObjet := map[string]int{
		"Cristal":  25,
		"Rune":     35,
		"Relique":  65,
		"Perle":    45,
		"Couronne": 95,
	}

	sac := make([]string, 0, 3)

	demoCapacite := make([]int, 3, 5)

	fmt.Printf("\nDémo make([]int, 3, 5) : %v\n", demoCapacite)
	fmt.Printf("len = %d | cap = %d\n", len(demoCapacite), cap(demoCapacite))

	scores := make([]int, 0, len(zonesExplorees))

	fmt.Println("\n=== Début de l'exploration ===")

	for i := 0; i < len(zonesExplorees); i++ {
		zone := zonesExplorees[i]
		objet := objetsParZone[zone]
		points := pointsParObjet[objet]

		sac = append(sac, objet)
		scores = append(scores, points)

		fmt.Printf("\nZone %d : %s\n", i+1, zone)
		fmt.Printf("Objet trouvé : %s\n", objet)
		fmt.Printf("Points gagnés : %d\n", points)
		fmt.Printf("Sac actuel : %v\n", sac)
		fmt.Printf("len(sac) = %d | cap(sac) = %d\n", len(sac), cap(sac))
	}

	sacSauvegarde := make([]string, len(sac))
	nombreCopies := copy(sacSauvegarde, sac)

	total := 0

	for _, score := range scores {
		total += score
	}

	rangFinal := calculerRang(total)

	fmt.Println("\n=== Résultat final ===")
	fmt.Printf("Explorateur : %s\n", pseudo)
	fmt.Printf("Objets dans le sac : %v\n", sac)
	fmt.Printf("Copie du sac : %v\n", sacSauvegarde)
	fmt.Printf("Nombre d'objets copiés : %d\n", nombreCopies)
	fmt.Printf("Scores obtenus : %v\n", scores)
	fmt.Printf("Score total : %d\n", total)
	fmt.Printf("Rang final : %s\n", nomsRangs[rangFinal])

	afficherBonus(rangFinal)
}