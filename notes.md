# Notes de cours — Langage Go

## Séance du jour — Rappels et nouvelles notions

Aujourd’hui, on a revu plusieurs bases du langage Go, mais aussi des notions un peu plus avancées comme les slices, les maps, les pointeurs, les closures, les structs ou encore la visibilité avec les majuscules/minuscules.

L’idée générale à retenir, c’est que Go est un langage assez strict, mais justement pensé pour rester simple, lisible et performant.

---

## 1. Les entiers

En Go, les entiers servent à stocker des nombres sans virgule.

On peut utiliser plusieurs types :

```go
int
int8
int16
int32
int64
uint
```

Le type le plus courant est simplement :

```go
int
```

C’est celui qu’on utilise dans la majorité des cas quand on veut manipuler des nombres entiers classiques.

---

## 2. Variables et constantes

Une variable est une valeur qui peut changer pendant l’exécution du programme.

Exemple :

```go
var age int = 20
age = 21
```

Une constante, au contraire, ne peut plus être modifiée après sa déclaration.

Exemple :

```go
const PI = 3.14
```

On peut aussi typer une constante si besoin.

---

## 3. Les différentes façons de déclarer une variable

En Go, on peut déclarer une variable de plusieurs manières.

### Avec `var`

```go
var nom string = "Maxime"
```

On indique le nom de la variable, son type, puis sa valeur.

### Avec `:=`

```go
nom := "Maxime"
```

Cette syntaxe permet de déclarer et initiLAI YIOser une variable directement.

Attention : `:=` ne s’utilise qu’à l’intérieur des fonctions.

### Déclarer plusieurs variables en même temps

```go
x, y := 10, 20
```

Go permet aussi d’échanger deux valeurs sans variable temporaire :

```go
x, y = y, x
```

C’est plus simple et plus lisible.

---

## 4. Différence entre `:=` et `=`

`:=` sert à créer une nouvelle variable.

```go
nom := "Maxime"
```

`=` sert à modifier une variable qui existe déjà.

```go
nom = "LAI YIO"
```

Il ne faut donc pas confondre les deux.

---

## 5. Les slices

Un slice est une sorte de tableau dynamique.

La différence avec un tableau classique, c’est que sa taille peut évoluer.

Exemple :

```go
notes := []int{12, 15, 18}
notes = append(notes, 20)
```

Ici, on ajoute une nouvelle valeur au slice avec `append`.

À retenir :

```go
Un slice = tableau dont la taille est modifiable
```

---

## 6. `len` et `cap`

Sur un slice, on peut utiliser deux fonctions importantes.

### `len`

`len` donne le nombre d’éléments actuellement présents.

```go
len(notes)
```

### `cap`

`cap` donne la capacité du slice, c’est-à-dire l’espace prévu en mémoire.

```go
cap(notes)
```

Exemple avec `make` :

```go
slice := make([]int, 3, 5)
```

Ici :

* longueur : 3
* capacité : 5

Donc le slice contient 3 éléments, mais Go a prévu de la place pour 5.

---

## 7. Copier un slice

En Go, il faut faire attention avec les slices.

Si on fait simplement :

```go
copie := original
```

on ne crée pas forcément une vraie copie indépendante.

Pour faire une vraie copie, on utilise `copy`.

Exemple :

```go
source := []int{1, 2, 3}
destination := make([]int, len(source))

copy(destination, source)
```

Là, `destination` contient les mêmes valeurs que `source`, mais c’est bien une copie séparée.

---

## 8. Les constantes avec `iota`

`iota` permet de créer des constantes automatiquement.

C’est pratique pour faire des énumérations.

Exemple :

```go
const (
	Faible = iota
	Moyen
	Eleve
)
```

Résultat :

```go
Faible = 0
Moyen = 1
Eleve = 2
```

À retenir :

```go
iota = générateur automatique de valeurs pour des constantes
```

---

## 9. Les maps

Une map est une structure clé / valeur.

C’est l’équivalent d’un dictionnaire dans d’autres langages.

Exemple :

```go
ages := map[string]int{
	"Maxime": 24,
	"LAI YIO":     22,
}
```

On accède ensuite à une valeur grâce à sa clé :

```go
fmt.Println(ages["Maxime"])
```

Ici, la clé est `"Maxime"` et la valeur est `24`.

---

## 10. La boucle `for`

En Go, il n’y a qu’une seule vraie boucle : `for`.

Exemple classique :

```go
for i := 0; i < 5; i++ {
	fmt.Println(i)
}
```

Elle peut aussi remplacer une boucle `while`.

Exemple :

```go
for condition {
	// code
}
```

Donc en Go, on utilise toujours `for`, même quand on veut faire une logique de type `while`.

---

## 11. Le `switch case`

Go possède aussi une structure `switch case`.

Exemple :

```go
switch niveau {
case 1:
	fmt.Println("Niveau 1")
case 2:
	fmt.Println("Niveau 2")
default:
	fmt.Println("Niveau inconnu")
}
```

C’est utile quand on veut tester plusieurs cas possibles.

---

## 12. `fallthrough`

Le mot-clé `fallthrough` force l’exécution du `case` suivant dans un `switch`, même si la condition du cas suivant ne correspond pas.

Exemple :

```go
switch niveau {
case 1:
	fmt.Println("Niveau 1")
	fallthrough
case 2:
	fmt.Println("Niveau 2")
}
```

Si `niveau = 1`, Go va afficher :

```go
Niveau 1
Niveau 2
```

À retenir : `fallthrough` est à utiliser avec prudence, car il force le passage au cas suivant.

---

## 13. Les fonctions variadiques

Une fonction variadique est une fonction qui peut recevoir un nombre variable d’arguments.

Exemple :

```go
func afficherNoms(noms ...string) {
	for _, nom := range noms {
		fmt.Println(nom)
	}
}
```

On peut ensuite l’appeler avec plusieurs valeurs :

```go
afficherNoms("Maxime", "LAI YIO", "Sarah")
```

À retenir :

```go
Fonction variadique = fonction avec un nombre d’arguments non défini à l’avance
```

---

## 14. Le retour multiple

En Go, les fonctions peuvent retourner plusieurs valeurs.

C’est très utilisé, surtout pour retourner un résultat et une erreur.

Exemple :

```go
resultat, err := calculer()

if err != nil {
	fmt.Println("Erreur :", err)
}
```

C’est une façon très courante de gérer les erreurs en Go.

---

## 15. Gestion des erreurs : Go vs Java

En Java, on utilise souvent `try / catch`.

```java
try {
	// code
} catch (Exception e) {
	// erreur
}
```

En Go, on préfère retourner une erreur explicitement.

```go
resultat, err := calculer()

if err != nil {
	fmt.Println("Erreur :", err)
}
```

Go oblige donc à gérer les erreurs de manière visible et directe.

---

## 16. Les closures

Une closure est une fonction qui garde en mémoire une variable de son environnement.

Exemple :

```go
func compteur() func() int {
	n := 0

	return func() int {
		n++
		return n
	}
}
```

Ici, la fonction retournée garde en mémoire la variable `n`.

Les closures sont utilisées dans plusieurs cas, par exemple :

* callbacks ;
* tri ;
* fonctions HTTP ;
* logique réutilisable.

À retenir :

```go
Closure = fonction qui garde un état en mémoire
```

---

## 17. Les structures

Une structure permet de regrouper plusieurs informations dans un seul type.

Exemple :

```go
type Personne struct {
	Nom string
	Age int
}
```

On peut ensuite créer une personne :

```go
p := Personne{
	Nom: "Maxime",
	Age: 24,
}
```

Une struct permet donc de représenter un objet ou une entité.

---

## 18. Les méthodes

Une méthode est une fonction liée à une struct.

Exemple :

```go
func (p Personne) SePresenter() {
	fmt.Println("Bonjour, je suis", p.Nom)
}
```

Ici, `SePresenter()` appartient à la struct `Personne`.

On peut ensuite faire :

```go
p.SePresenter()
```

---

## 19. La notion d’objet en Go

Go n’a pas de classes comme Java.

À la place, on utilise :

* des structs ;
* des méthodes ;
* des interfaces ;
* de la composition.

Exemple :

```go
type Produit struct {
	Nom  string
	Prix float64
}
```

Même sans classes, on peut quand même organiser son code de manière orientée objet.

---

## 20. Classes, héritage et composition

Go ne fonctionne pas avec un héritage classique comme Java.

À la place, il utilise la composition, notamment avec l’embedding.

Exemple :

```go
type Personne struct {
	Nom string
}

type Employe struct {
	Personne
	Poste string
}
```

Ici, `Employe` contient une `Personne`.

À retenir :

```go
Go préfère la composition plutôt que l’héritage
```

---

## 21. L’embedding

L’embedding permet d’inclure une struct dans une autre struct.

Exemple :

```go
type Employe struct {
	Personne
	Poste string
}
```

Grâce à ça, `Employe` peut accéder directement aux champs de `Personne`.

C’est une manière simple de réutiliser du code sans héritage classique.

---

## 22. Les pointeurs

Un pointeur permet de travailler avec l’adresse mémoire d’une variable.

Exemple :

```go
func modifierAge(p *Personne) {
	p.Age = 25
}
```

L’intérêt est de modifier directement la valeur originale au lieu de travailler sur une copie.

À retenir :

```go
Pointeur = accès direct à la donnée originale en mémoire
```

Les pointeurs peuvent aussi être utilisés pour optimiser certains usages, notamment quand on manipule de grosses structures.

---

## 23. Les struct tags

Les struct tags permettent d’ajouter des métadonnées aux champs d’une struct.

Ils sont souvent utilisés pour le JSON.

Exemple :

```go
type Personne struct {
	Nom string `json:"nom,omitempty"`
}
```

Ici, le champ `Nom` sera transformé en `"nom"` dans le JSON.

Le mot-clé `omitempty` permet de ne pas afficher le champ s’il est vide.

Autre exemple :

```go
MotDePasse string `json:"-"`
```

Ici, le champ `MotDePasse` ne sera pas sériLAI YIOsé en JSON.

---

## 24. Visibilité en Go

En Go, la visibilité dépend de la casse.

Si un nom commence par une majuscule, il est exporté, donc accessible depuis un autre package.

Exemple :

```go
type Personne struct {}
```

Si un nom commence par une minuscule, il est privé au package.

Exemple :

```go
type personne struct {}
```

À retenir :

```go
Majuscule = public / exporté
Minuscule = privé au package
```

---

## 25. Camel Case et Pascal Case

En Go, la casse a une vraie importance.

### Camel case

Le camel case commence par une minuscule.

Exemples :

```go
nomProduit
calculerPrix
```

En Go, cela signifie que l’élément reste privé au package.

### Pascal case

Le Pascal case commence par une majuscule.

Exemples :

```go
NomProduit
CalculerPrix
```

En Go, cela signifie que l’élément est exporté et accessible depuis un autre package.

---

## 26. Pas de `private`, `public`, `protected`

Go n’utilise pas les mots-clés classiques comme :

```go
private
public
protected
```

La visibilité se fait uniquement avec la casse.

Résumé :

* majuscule : public / exporté ;
* minuscule : privé au package ;
* pas de vrai `protected` comme en Java.

---

## 27. La surcharge de fonctions

Go ne permet pas la surcharge de fonctions.

En Java, on peut avoir plusieurs fonctions avec le même nom mais des paramètres différents.

En Go, ce n’est pas possible.

Chaque fonction doit avoir un nom unique dans le même package.

Avantage :

```text
Le code reste plus simple et plus lisible.
```

Inconvénient :

```text
Il faut parfois créer plusieurs fonctions avec des noms différents.
```

---

## 28. Le mot-clé `defer`

`defer` permet de garantir l’exécution d’une instruction à la fin d’une fonction.

Exemple :

```go
defer fmt.Println("Fin du programme")
```

Même si le code continue, cette instruction sera exécutée à la fin de la fonction.

On l’utilise souvent pour fermer un fichier ou une connexion.

Exemple :

```go
defer fichier.Close()
```

À retenir :

```go
defer = exécution garantie à la fin de la fonction
```

---

## 29. `fmt` et le formatage

Le package `fmt` est souvent utilisé pour afficher du texte ou formater des valeurs.

Exemple :

```go
fmt.Println("Bonjour")
```

Go est assez strict dans sa manière de formater le code, mais cela permet d’avoir un style uniforme.

L’idée est que tout le monde écrit du Go de manière assez similaire, ce qui rend le code plus facile à lire.

---

## 30. Exemple marquant : le bug d’Ariane 5

Le cours a aussi évoqué le bug d’Ariane 5.

À retenir :

```text
Coût estimé du bug : environ 500 millions
```

Cela montre qu’une erreur informatique peut avoir des conséquences énormes, surtout dans des systèmes critiques.

---

## 31. Points importants à retenir

Go est un langage :

* compilé ;
* strict ;
* simple à lire ;
* orienté performance ;
* basé sur la composition plutôt que l’héritage ;
* sans classes classiques ;
* sans surcharge de fonctions ;
* avec une gestion explicite des erreurs ;
* avec une seule vraie boucle : `for`.

Les notions importantes vues aujourd’hui :

* variables et constantes ;
* `:=` et `=` ;
* slices ;
* `len` et `cap` ;
* copie de slices avec `copy` ;
* maps ;
* boucle `for` ;
* switch case ;
* `fallthrough` ;
* fonctions variadiques ;
* closures ;
* structs ;
* méthodes ;
* pointeurs ;
* embedding ;
* struct tags ;
* visibilité avec majuscule/minuscule ;
* `defer`.

---

## 32. Résumé rapide

```text
Variable = valeur modifiable
Constante = valeur fixe
Slice = tableau dynamique
Map = clé / valeur
For = seule boucle en Go
Iota = génération automatique de constantes
Closure = fonction qui garde une variable en mémoire
Struct = regroupement de données
Méthode = fonction liée à une struct
Pointeur = accès à l’adresse mémoire
Embedding = composition entre structs
Defer = exécution à la fin d’une fonction
Majuscule = exporté
Minuscule = privé au package
```

---

## Conclusion

Cette séance a permis de consolider les bases du langage Go tout en abordant des notions plus avancées.

Le point central à retenir, c’est que Go cherche à éviter la complexité inutile. Il ne reprend pas toutes les mécaniques classiques de Java comme les classes, l’héritage ou la surcharge, mais propose une approche plus simple avec les structs, les méthodes, les interfaces et la composition.

Go est donc un langage rigide, mais cette rigidité aide à écrire du code clair, lisible et maintenable.
