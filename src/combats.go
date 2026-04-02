package main

import (
	"bufio"
	"fmt"
	"os"
)

func initialiserGobelin() Monstre {
	return Monstre{
		Nom:        "Gobelin d'Entraînement",
		HPMax:      40,
		HPActuel:   40,
		Attaque:    5,
		Initiative: 8,
	}
}

func patternGobelin(gobelin *Monstre, joueur *Personnage, tour int) {
	degats := gobelin.Attaque
	if tour%3 == 0 {
		degats = gobelin.Attaque * 2
		fmt.Printf(ROUGE+GRAS+"💥 %s rugit et prépare une attaque dévastatrice !\n"+RESET, gobelin.Nom)
	}
	joueur.HPActuel -= degats
	if joueur.HPActuel < 0 {
		joueur.HPActuel = 0
	}
	fmt.Printf(ROUGE+"⚔️  %s inflige %d dégâts à %s\n"+RESET, gobelin.Nom, degats, joueur.Nom)
	fmt.Printf(ROUGE_VIF+"❤️  %s HP : %d/%d\n"+RESET, joueur.Nom, joueur.HPActuel, joueur.HPMax)
}

func tourJoueur(joueur *Personnage, gobelin *Monstre) bool {
	lecteur := bufio.NewReader(os.Stdin)
	for {
		fmt.Println(VERT_VIF + "\n┌─────────────────────────────────────┐")
		fmt.Println("│  ⚔️  VOTRE TOUR, BRAVE HÉROS !       │")
		fmt.Println("├─────────────────────────────────────┤" + RESET)
		fmt.Println(CYAN_VIF + "│  1. 🗡️  Attaque basique (5 dégâts)  │" + RESET)
		fmt.Println(EMERAUDE + "│  2. ✨ Lancer un sort               │" + RESET)
		fmt.Println(OR + "│  3. 🎒 Utiliser l'inventaire        │" + RESET)
		fmt.Println(ARGENT + "│  0. 🏃 Fuir (retour au menu)        │" + RESET)
		fmt.Println(VERT_VIF + "└─────────────────────────────────────┘" + RESET)

		var choix int
		fmt.Print(EMERAUDE + "➤ Votre choix : " + RESET)
		fmt.Scan(&choix)
		lecteur.ReadString('\n')

		switch choix {
		case 1:
			degats := 5
			gobelin.HPActuel -= degats
			if gobelin.HPActuel < 0 {
				gobelin.HPActuel = 0
			}
			fmt.Printf(VERT_VIF+"🗡️  %s frappe avec force : %d dégâts à %s !\n"+RESET, joueur.Nom, degats, gobelin.Nom)
			fmt.Printf(JAUNE_VIVE+"👹 %s HP : %d/%d\n"+RESET, gobelin.Nom, gobelin.HPActuel, gobelin.HPMax)
			return false
		case 2:
			if lancerSort(joueur, gobelin) {
				return false
			}
		case 3:
			if len(joueur.Inventaire) == 0 {
				fmt.Println(ROUGE + "🎒 Votre sac est vide, héros !" + RESET)
				continue
			}
			fmt.Println(OR + "\n🎒 Votre inventaire :" + RESET)
			for i, item := range joueur.Inventaire {
				fmt.Printf(CYAN_VIF+"  %d. %s\n"+RESET, i+1, item)
			}
			fmt.Println(ARGENT + "  0. Annuler" + RESET)
			var choixItem int
			fmt.Print(EMERAUDE + "➤ Votre choix : " + RESET)
			fmt.Scan(&choixItem)
			lecteur.ReadString('\n')
			if choixItem == 0 {
				continue
			}
			if choixItem < 1 || choixItem > len(joueur.Inventaire) {
				fmt.Println(ROUGE + "❌ Choix invalide." + RESET)
				continue
			}
			item := joueur.Inventaire[choixItem-1]
			utiliserItemCombat(joueur, item, choixItem-1)
			return false
		case 0:
			fmt.Println(JAUNE_VIVE + "🏃 Vous battez en retraite... La prudence est mère de sûreté !" + RESET)
			return true
		default:
			fmt.Println(ROUGE + "❌ Choix invalide." + RESET)
		}
	}
}

func utiliserItemCombat(p *Personnage, item string, index int) {
	switch item {
	case "Potion de Vie":
		fmt.Printf(VERT_VIF+"🧪 %s boit une Potion de Vie en plein combat !\n"+RESET, p.Nom)
		utiliserPotionVie(p, index)
	default:
		fmt.Printf(ROUGE+"❌ '%s' ne peut pas être utilisé en combat.\n"+RESET, item)
	}
}

var degatsSort = map[string]int{
	"Coup de Poing": 8,
	"Boule de Feu":  18,
}

var coutManaSort = map[string]int{
	"Coup de Poing": 5,
	"Boule de Feu":  15,
}

func lancerSort(joueur *Personnage, gobelin *Monstre) bool {
	lecteur := bufio.NewReader(os.Stdin)
	fmt.Printf(EMERAUDE+"\n✨ Sorts disponibles (Mana : %d/%d) :\n"+RESET, joueur.ManaActuel, joueur.ManaMax)
	for i, s := range joueur.Sorts {
		cout := coutManaSort[s]
		dgts := degatsSort[s]
		fmt.Printf(CYAN_VIF+"  %d. %s (dégâts : %d | mana : %d)\n"+RESET, i+1, s, dgts, cout)
	}
	fmt.Println(ARGENT + "  0. Annuler" + RESET)

	var choix int
	fmt.Print(EMERAUDE + "➤ Votre choix : " + RESET)
	fmt.Scan(&choix)
	lecteur.ReadString('\n')

	if choix == 0 {
		return false
	}
	if choix < 1 || choix > len(joueur.Sorts) {
		fmt.Println(ROUGE + "❌ Choix invalide." + RESET)
		return false
	}

	sort := joueur.Sorts[choix-1]
	cout := coutManaSort[sort]
	dgts := degatsSort[sort]

	if joueur.ManaActuel < cout {
		fmt.Printf(BLEU_VIF+"💧 Mana insuffisant ! Il vous faut %d mana (vous avez %d).\n"+RESET, cout, joueur.ManaActuel)
		return false
	}

	joueur.ManaActuel -= cout
	gobelin.HPActuel -= dgts
	if gobelin.HPActuel < 0 {
		gobelin.HPActuel = 0
	}
	fmt.Printf(EMERAUDE+"✨ %s invoque '%s' : %d dégâts magiques à %s !\n"+RESET, joueur.Nom, sort, dgts, gobelin.Nom)
	fmt.Printf(JAUNE_VIVE+"👹 %s HP : %d/%d\n"+RESET, gobelin.Nom, gobelin.HPActuel, gobelin.HPMax)
	return true
}

func combatEntrainement(joueur *Personnage) {
	gobelin := initialiserGobelin()
	tour := 1

	fmt.Println(ROUGE_VIF + "\nCOMBAT D'ENTRAÎNEMENT" + RESET)
	fmt.Printf(JAUNE_VIVE+"%s VS %s\n"+RESET, joueur.Nom, gobelin.Nom)

	joueurCommence := joueur.Initiative >= gobelin.Initiative
	if joueurCommence {
		fmt.Printf(VERT_VIF+"🎲 %s a plus d'initiative et attaque en premier !\n"+RESET, joueur.Nom)
	} else {
		fmt.Printf(ROUGE+"🎲 %s surgit des ombres et attaque en premier !\n"+RESET, gobelin.Nom)
	}

	for {
		fmt.Printf(POURPRE+"\n--- TOUR %d ---\n"+RESET, tour)
		fmt.Printf(VERT_VIF+"❤️  %s : %d/%d"+RESET+" | "+ROUGE+"👹 %s : %d/%d\n"+RESET,
			joueur.Nom, joueur.HPActuel, joueur.HPMax,
			gobelin.Nom, gobelin.HPActuel, gobelin.HPMax)

		if joueurCommence {
			aFui := tourJoueur(joueur, &gobelin)
			if aFui {
				return
			}
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
			aFui := tourJoueur(joueur, &gobelin)
			if aFui {
				return
			}
			if gobelin.HPActuel <= 0 {
				victoire(joueur, &gobelin)
				return
			}
		}

		estMort(joueur)
		tour++
	}
}

func victoire(joueur *Personnage, gobelin *Monstre) {
	fmt.Println(OR + GRAS + "\nVICTOIRE !" + RESET)
	fmt.Printf(OR+"🏆 Vous avez vaincu %s !\n"+RESET, gobelin.Nom)
	gagnerXP(joueur, 30)
}

func defaite(joueur *Personnage) {
	fmt.Println(ROUGE + GRAS + "\nDÉFAITE..." + RESET)
	fmt.Println(ROUGE + "💀 Vous retournez au camp..." + RESET)
	estMort(joueur)
}

func gagnerXP(p *Personnage, xp int) {
	fmt.Printf(OR+"⭐ Vous gagnez %d points d'expérience !\n"+RESET, xp)
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
		fmt.Printf(OR+GRAS+"🎉 NIVEAU %d ! HP max : %d | Mana max : %d\n"+RESET, p.Niveau, p.HPMax, p.ManaMax)
	}
	fmt.Printf(JAUNE_VIVE+"📊 XP : %d/%d\n"+RESET, p.XPActuel, p.XPMax)
}
