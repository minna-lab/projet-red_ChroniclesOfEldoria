package main

import (
	"bufio"
	"fmt"
	"os"
)

// initialiserGobelin crée un gobelin d'entraînement
func initialiserGobelin() Monstre {
	return Monstre{
		Nom:        "Gobelin d'Entraînement",
		HPMax:      40,
		HPActuel:   40,
		Attaque:    5,
		Initiative: 8,
	}
}

// patternGobelin gère l'attaque du gobelin selon le tour
func patternGobelin(gobelin *Monstre, joueur *Personnage, tour int) {
	degats := gobelin.Attaque
	if tour%3 == 0 {
		degats = gobelin.Attaque * 2
		fmt.Printf("💥 %s prépare une attaque puissante !\n", gobelin.Nom)
	}
	joueur.HPActuel -= degats
	if joueur.HPActuel < 0 {
		joueur.HPActuel = 0
	}
	fmt.Printf("⚔️  %s inflige %d dégâts à %s\n", gobelin.Nom, degats, joueur.Nom)
	fmt.Printf("❤️  %s HP : %d/%d\n", joueur.Nom, joueur.HPActuel, joueur.HPMax)
}

// tourJoueur gère le tour du joueur en combat
func tourJoueur(joueur *Personnage, gobelin *Monstre) {
	lecteur := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n┌─────────────────────────────────────┐")
		fmt.Println("│  ⚔️  VOTRE TOUR                      │")
		fmt.Println("├─────────────────────────────────────┤")
		fmt.Println("│  1. Attaquer (Attaque basique)      │")
		fmt.Println("│  2. Lancer un sort                  │")
		fmt.Println("│  3. Utiliser l'inventaire           │")
		fmt.Println("└─────────────────────────────────────┘")

		var choix int
		fmt.Print("➤ Votre choix : ")
		fmt.Scan(&choix)
		lecteur.ReadString('\n')

		switch choix {
		case 1:
			// Attaque basique : 5 dégâts
			degats := 5
			gobelin.HPActuel -= degats
			if gobelin.HPActuel < 0 {
				gobelin.HPActuel = 0
			}
			fmt.Printf("🗡️  Attaque Basique : %s inflige %d dégâts à %s\n", joueur.Nom, degats, gobelin.Nom)
			fmt.Printf("👹 %s HP : %d/%d\n", gobelin.Nom, gobelin.HPActuel, gobelin.HPMax)
			return

		case 2:
			if lancerSort(joueur, gobelin) {
				return
			}

		case 3:
			if len(joueur.Inventaire) == 0 {
				fmt.Println("🎒 Inventaire vide !")
				continue
			}
			fmt.Println("\n🎒 Inventaire :")
			for i, item := range joueur.Inventaire {
				fmt.Printf("  %d. %s\n", i+1, item)
			}
			fmt.Println("  0. Annuler")
			var choixItem int
			fmt.Print("➤ Votre choix : ")
			fmt.Scan(&choixItem)
			lecteur.ReadString('\n')
			if choixItem == 0 {
				continue
			}
			if choixItem < 1 || choixItem > len(joueur.Inventaire) {
				fmt.Println("❌ Choix invalide.")
				continue
			}
			item := joueur.Inventaire[choixItem-1]
			utiliserItemCombat(joueur, item, choixItem-1)
			return

		default:
			fmt.Println("❌ Choix invalide.")
		}
	}
}

// utiliserItemCombat utilise un item pendant le combat
func utiliserItemCombat(p *Personnage, item string, index int) {
	switch item {
	case "Potion de Vie":
		fmt.Printf("🧪 %s utilise une Potion de Vie !\n", p.Nom)
		utiliserPotionVie(p, index)
	default:
		fmt.Printf("❌ '%s' ne peut pas être utilisé en combat.\n", item)
	}
}

// degatsSort contient les dégâts de chaque sort
var degatsSort = map[string]int{
	"Coup de Poing": 8,
	"Boule de Feu":  18,
}

// coutManaSort contient le coût en mana de chaque sort
var coutManaSort = map[string]int{
	"Coup de Poing": 5,
	"Boule de Feu":  15,
}

// lancerSort gère l'utilisation d'un sort en combat
func lancerSort(joueur *Personnage, gobelin *Monstre) bool {
	lecteur := bufio.NewReader(os.Stdin)
	fmt.Printf("\n✨ Sorts disponibles (Mana : %d/%d) :\n", joueur.ManaActuel, joueur.ManaMax)
	for i, s := range joueur.Sorts {
		cout := coutManaSort[s]
		dgts := degatsSort[s]
		fmt.Printf("  %d. %s (dégâts : %d | mana : %d)\n", i+1, s, dgts, cout)
	}
	fmt.Println("  0. Annuler")

	var choix int
	fmt.Print("➤ Votre choix : ")
	fmt.Scan(&choix)
	lecteur.ReadString('\n')

	if choix == 0 {
		return false
	}
	if choix < 1 || choix > len(joueur.Sorts) {
		fmt.Println("❌ Choix invalide.")
		return false
	}

	sort := joueur.Sorts[choix-1]
	cout := coutManaSort[sort]
	dgts := degatsSort[sort]

	if joueur.ManaActuel < cout {
		fmt.Printf("💧 Mana insuffisant ! Il vous faut %d mana (vous avez %d).\n", cout, joueur.ManaActuel)
		return false
	}

	joueur.ManaActuel -= cout
	gobelin.HPActuel -= dgts
	if gobelin.HPActuel < 0 {
		gobelin.HPActuel = 0
	}
	fmt.Printf("✨ %s lance '%s' : %d dégâts à %s !\n", joueur.Nom, sort, dgts, gobelin.Nom)
	fmt.Printf("👹 %s HP : %d/%d\n", gobelin.Nom, gobelin.HPActuel, gobelin.HPMax)
	return true
}

// combatEntrainement lance un combat contre le gobelin
func combatEntrainement(joueur *Personnage) {
	gobelin := initialiserGobelin()
	tour := 1

	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║   ⚔️  COMBAT D'ENTRAÎNEMENT              ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Printf("║  %s VS %s\n", joueur.Nom, gobelin.Nom)
	fmt.Println("╚══════════════════════════════════════════╝")

	// Initiative : celui qui a le plus commence
	joueurCommence := joueur.Initiative >= gobelin.Initiative
	if joueurCommence {
		fmt.Printf("🎲 %s a plus d'initiative et commence !\n", joueur.Nom)
	} else {
		fmt.Printf("🎲 %s a plus d'initiative et commence !\n", gobelin.Nom)
	}

	for {
		fmt.Printf("\n══════════ TOUR %d ══════════\n", tour)
		fmt.Printf("❤️  %s : %d/%d | 👹 %s : %d/%d\n",
			joueur.Nom, joueur.HPActuel, joueur.HPMax,
			gobelin.Nom, gobelin.HPActuel, gobelin.HPMax)

		if joueurCommence {
			tourJoueur(joueur, &gobelin)
			if gobelin.HPActuel <= 0 {
				victoire(joueur, &gobelin)
				return
			}
			patternGobelin(&gobelin, joueur, tour)
			if joueur.HPActuel <= 0 {
				defaite(joueur)
				return
			}
		} else {
			patternGobelin(&gobelin, joueur, tour)
			if joueur.HPActuel <= 0 {
				defaite(joueur)
				return
			}
			tourJoueur(joueur, &gobelin)
			if gobelin.HPActuel <= 0 {
				victoire(joueur, &gobelin)
				return
			}
		}

		estMort(joueur)
		tour++
	}
}

// victoire gère la fin de combat gagnée
func victoire(joueur *Personnage, gobelin *Monstre) {
	fmt.Printf("\n🏆 Victoire ! Vous avez vaincu %s !\n", gobelin.Nom)
	gagnerXP(joueur, 30)
}

// defaite gère la fin de combat perdue
func defaite(joueur *Personnage) {
	fmt.Println("\n💀 Vous avez été vaincu... Retour au camp de base.")
	estMort(joueur)
}

// gagnerXP ajoute de l'expérience et gère la montée de niveau
func gagnerXP(p *Personnage, xp int) {
	fmt.Printf("⭐ Vous gagnez %d points d'expérience !\n", xp)
	p.XPActuel += xp
	for p.XPActuel >= p.XPMax {
		exces := p.XPActuel - p.XPMax
		p.Niveau++
		p.XPMax = int(float64(p.XPMax) * 1.5)
		p.XPActuel = exces
		p.HPMax += 10
		p.HPActuel = p.HPMax
		p.ManaMax += 10
		p.ManaActuel = p.ManaMax
		fmt.Printf("🎉 NIVEAU %d ! HP max : %d | Mana max : %d\n", p.Niveau, p.HPMax, p.ManaMax)
	}
	fmt.Printf("📊 XP : %d/%d\n", p.XPActuel, p.XPMax)
}
