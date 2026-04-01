package main

import (
	"bufio"
	"fmt"
	"os"
)

// menuPrincipal affiche le menu principal et gère la navigation
func menuPrincipal(p *Personnage) {
	lecteur := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n╔══════════════════════════════════════════╗")
		fmt.Println("║    🏰 CHRONICLES OF ELDORIA 🏰           ║")
		fmt.Println("╠══════════════════════════════════════════╣")
		fmt.Printf("║  Héros : %-10s | HP : %d/%d\n", p.Nom, p.HPActuel, p.HPMax)
		fmt.Printf("║  Or    : %d 🪙\n", p.Or)
		fmt.Println("╠══════════════════════════════════════════╣")
		fmt.Println("║  1. Afficher les informations            ║")
		fmt.Println("║  2. Accéder à l'inventaire               ║")
		fmt.Println("║  3. Marchand                             ║")
		fmt.Println("║  4. Forgeron                             ║")
		fmt.Println("║  5. Combat d'entraînement                ║")
		fmt.Println("║  6. Qui sont-ils ?                       ║")
		fmt.Println("║  0. Quitter                              ║")
		fmt.Println("╚══════════════════════════════════════════╝")

		var choix int
		fmt.Print("➤ Votre choix : ")
		fmt.Scan(&choix)
		lecteur.ReadString('\n')

		switch choix {
		case 1:
			afficherInfos(p)
		case 2:
			accederInventaire(p)
		case 3:
			menuMarchand(p)
		case 4:
			menuForgeron(p)
		case 5:
			combatEntrainement(p)
		case 6:
			quiSontIls()
		case 0:
			fmt.Println("\n⚔️  Les Chroniques d'Eldoria vous attendent. À bientôt !")
			return
		default:
			fmt.Println("❌ Choix invalide.")
		}
	}
}

// ObjetBoutique représente un objet vendu par le marchand
type ObjetBoutique struct {
	Nom   string
	Prix  int
	Emoji string
}

// listeObjets contient tous les objets disponibles chez le marchand
var listeObjets = []ObjetBoutique{
	{"Potion de Vie", 3, "🧪"},
	{"Potion de Poison", 6, "☠️"},
	{"Livre de Sort : Boule de Feu", 25, "📖"},
	{"Fourrure de Loup", 4, "🐺"},
	{"Peau de Troll", 7, "👹"},
	{"Cuir de Sanglier", 3, "🐗"},
	{"Plume de Corbeau", 1, "🪶"},
	{"Augmentation d'Inventaire", 30, "🎒"},
}

// menuMarchand affiche le menu du marchand et gère les achats
func menuMarchand(p *Personnage) {
	lecteur := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n╔══════════════════════════════════════════╗")
		fmt.Println("║   🛒 MARCHAND D'ELDORIA                  ║")
		fmt.Println("╠══════════════════════════════════════════╣")
		fmt.Printf("║  Votre bourse : %d 🪙\n", p.Or)
		fmt.Println("╠══════════════════════════════════════════╣")
		for i, obj := range listeObjets {
			fmt.Printf("║  %d. %s %-25s %d 🪙\n", i+1, obj.Emoji, obj.Nom, obj.Prix)
		}
		fmt.Println("╠══════════════════════════════════════════╣")
		fmt.Println("║  0. Retour                               ║")
		fmt.Println("╚══════════════════════════════════════════╝")

		var choix int
		fmt.Print("➤ Votre choix : ")
		fmt.Scan(&choix)
		lecteur.ReadString('\n')

		if choix == 0 {
			return
		}
		if choix < 1 || choix > len(listeObjets) {
			fmt.Println("❌ Choix invalide.")
			continue
		}

		obj := listeObjets[choix-1]
		if p.Or < obj.Prix {
			fmt.Printf("💸 Pas assez d'or ! Il vous faut %d 🪙 (vous avez %d 🪙).\n", obj.Prix, p.Or)
			continue
		}
		if inventairePlein(p) {
			fmt.Println("🎒 Inventaire plein ! Faites de la place avant d'acheter.")
			continue
		}
		p.Or -= obj.Prix
		p.Inventaire = append(p.Inventaire, obj.Nom)
		fmt.Printf("✅ Vous achetez '%s' pour %d 🪙. Il vous reste %d 🪙.\n", obj.Nom, obj.Prix, p.Or)
	}
}

// Recette représente une recette de fabrication du forgeron
type Recette struct {
	Nom       string
	CoutOr    int
	Materiaux map[string]int
	Emoji     string
}

// listeRecettes contient toutes les recettes disponibles
var listeRecettes = []Recette{
	{
		Nom:    "Chapeau de l'Aventurier",
		CoutOr: 5,
		Materiaux: map[string]int{
			"Plume de Corbeau": 1,
			"Cuir de Sanglier": 1,
		},
		Emoji: "🪖",
	},
	{
		Nom:    "Tunique de l'Aventurier",
		CoutOr: 5,
		Materiaux: map[string]int{
			"Fourrure de Loup": 2,
			"Peau de Troll":    1,
		},
		Emoji: "🧥",
	},
	{
		Nom:    "Bottes de l'Aventurier",
		CoutOr: 5,
		Materiaux: map[string]int{
			"Fourrure de Loup": 1,
			"Cuir de Sanglier": 1,
		},
		Emoji: "👢",
	},
}

// menuForgeron affiche le menu du forgeron
func menuForgeron(p *Personnage) {
	lecteur := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n╔══════════════════════════════════════════╗")
		fmt.Println("║   🔨 FORGERON D'ELDORIA                  ║")
		fmt.Println("╠══════════════════════════════════════════╣")
		fmt.Printf("║  Votre bourse : %d 🪙\n", p.Or)
		fmt.Println("╠══════════════════════════════════════════╣")
		for i, r := range listeRecettes {
			fmt.Printf("║  %d. %s %s (coût : %d 🪙)\n", i+1, r.Emoji, r.Nom, r.CoutOr)
			for mat, qte := range r.Materiaux {
				possede := compterItem(p, mat)
				fmt.Printf("║     - %dx %s (vous avez : %d)\n", qte, mat, possede)
			}
		}
		fmt.Println("╠══════════════════════════════════════════╣")
		fmt.Println("║  0. Retour                               ║")
		fmt.Println("╚══════════════════════════════════════════╝")

		var choix int
		fmt.Print("➤ Votre choix : ")
		fmt.Scan(&choix)
		lecteur.ReadString('\n')

		if choix == 0 {
			return
		}
		if choix < 1 || choix > len(listeRecettes) {
			fmt.Println("❌ Choix invalide.")
			continue
		}
		fabriquer(p, listeRecettes[choix-1])
	}
}

// fabriquer tente de fabriquer un équipement
func fabriquer(p *Personnage, r Recette) {
	if p.Or < r.CoutOr {
		fmt.Printf("💸 Pas assez d'or ! Il vous faut %d 🪙.\n", r.CoutOr)
		return
	}
	for mat, qte := range r.Materiaux {
		if compterItem(p, mat) < qte {
			fmt.Printf("❌ Matériaux insuffisants : il vous faut %dx %s.\n", qte, mat)
			return
		}
	}
	if inventairePlein(p) {
		fmt.Println("🎒 Inventaire plein ! Impossible de fabriquer cet objet.")
		return
	}
	for mat, qte := range r.Materiaux {
		for i := 0; i < qte; i++ {
			retirerItemUneFois(p, mat)
		}
	}
	p.Or -= r.CoutOr
	p.Inventaire = append(p.Inventaire, r.Nom)
	fmt.Printf("⚒️  Vous fabriquez '%s' ! (-%d 🪙)\n", r.Nom, r.CoutOr)
}

// quiSontIls révèle les artistes cachés dans les noms des tâches
func quiSontIls() {
	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║   🎵 QUI SONT-ILS ?                      ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Println("║  Partie 2 : ABBA                         ║")
	fmt.Println("║  (Money Money Money,                     ║")
	fmt.Println("║   Gimme! Gimme! Gimme!,                  ║")
	fmt.Println("║   On and On and On)                      ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Println("║  Partie 3 : Références cinéma            ║")
	fmt.Println("║  (Ready Player One, A.I., Duel)          ║")
	fmt.Println("╚══════════════════════════════════════════╝")
}
