package main

import (
	"bufio"
	"fmt"
	"os"
)

func menuPrincipal(p *Personnage) {
	lecteur := bufio.NewReader(os.Stdin)
	for {
		fmt.Println(POURPRE + "\nCHRONICLES OF ELDORIA")
		fmt.Printf(OR+"Héros : %s | HP : %d/%d\n"+RESET, p.Nom, p.HPActuel, p.HPMax)
		fmt.Printf(OR+"Or : %d pièces\n"+RESET, p.Or)
		fmt.Println(POURPRE + "---" + RESET)
		fmt.Println(OR + "1. 📜 Afficher les informations" + RESET)
		fmt.Println(VERT + "2. 🎒 Accéder à l'inventaire" + RESET)
		fmt.Println(BLANC + "3. 🛒 Marchand d'Eldoria" + RESET)
		fmt.Println(OR + "4. ⚔️  Combat d'entraînement" + RESET)
		fmt.Println(VERT + "5. 🎵 Qui sont-ils ?" + RESET)
		fmt.Println(BLANC + "0. 🚪 Quitter" + RESET)

		var choix int
		fmt.Print(POURPRE + "➤ Votre choix : " + RESET)
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
			combatEntrainement(p)
		case 5:
			quiSontIls()
		case 0:
			fmt.Println(POURPRE + "\n⚔️  Les Chroniques d'Eldoria vous attendent, brave héros. À bientôt !" + RESET)
			return
		default:
			fmt.Println(ROUGE + "❌ Choix invalide, entrez un nombre entre 0 et 5." + RESET)
		}
	}
}

type ObjetBoutique struct {
	Nom   string
	Prix  int
	Emoji string
	Desc  string
}

var listeObjets = []ObjetBoutique{
	{"Potion de Vie", 3, "🧪", "Restaure 50 points de vie"},
	{"Potion de Poison", 6, "☠️", "Inflige 10 dégâts/sec pendant 3s"},
	{"Livre de Sort : Boule de Feu", 25, "📖", "Apprend le sort légendaire Boule de Feu"},
	{"Augmentation d'Inventaire", 30, "🎒", "+10 places dans l'inventaire (max 3 fois)"},
}

func menuMarchand(p *Personnage) {
	lecteur := bufio.NewReader(os.Stdin)
	for {
		fmt.Println(OR + "\nMARCHAND D'ELDORIA" + RESET)
		fmt.Printf(OR+"💰 Votre bourse : %d pièces d'or\n"+RESET, p.Or)
		fmt.Printf(CYAN_VIF+"🎒 Inventaire   : %d/%d\n"+RESET, len(p.Inventaire), p.InventaireMax)
		fmt.Println(JAUNE_VIVE + "Que souhaitez-vous acquérir, voyageur ?" + RESET)
		for i, obj := range listeObjets {
			fmt.Printf(CYAN_VIF+"║  %d. %s %-28s"+OR+"%d 🪙\n"+RESET, i+1, obj.Emoji, obj.Nom, obj.Prix)
			fmt.Printf(ARGENT+"║     → %s\n"+RESET, obj.Desc)
		}
		fmt.Println(ARGENT + "0. 🚪 Retour au menu principal" + RESET)

		var choix int
		fmt.Print(EMERAUDE + "➤ Votre choix : " + RESET)
		fmt.Scan(&choix)
		lecteur.ReadString('\n')

		if choix == 0 {
			fmt.Println(OR + "👋 Que les dieux d'Eldoria vous protègent, voyageur !" + RESET)
			return
		}
		if choix < 1 || choix > len(listeObjets) {
			fmt.Printf(ROUGE+"❌ Choix invalide, entrez un nombre entre 0 et %d.\n"+RESET, len(listeObjets))
			continue
		}

		obj := listeObjets[choix-1]

		if p.Or < obj.Prix {
			fmt.Printf(ROUGE+"💸 Pas assez d'or ! Il vous faut %d 🪙 mais vous n'avez que %d 🪙.\n"+RESET, obj.Prix, p.Or)
			continue
		}
		if inventairePlein(p) {
			fmt.Printf(ROUGE+"🎒 Votre sac est plein (%d/%d) ! Allégez-le avant d'acheter.\n"+RESET, len(p.Inventaire), p.InventaireMax)
			continue
		}
		if obj.Nom == "Livre de Sort : Boule de Feu" {
			dejaConnu := false
			for _, s := range p.Sorts {
				if s == "Boule de Feu" {
					dejaConnu = true
					break
				}
			}
			if dejaConnu {
				fmt.Println(JAUNE_VIVE + "📖 Vous maîtrisez déjà l'art de la Boule de Feu !" + RESET)
				continue
			}
		}

		p.Or -= obj.Prix
		p.Inventaire = append(p.Inventaire, obj.Nom)
		fmt.Printf(VERT_VIF+"✅ Vous acquérez '%s' pour %d 🪙.\n"+RESET, obj.Nom, obj.Prix)
		fmt.Printf(OR+"💰 Il vous reste %d pièces d'or.\n"+RESET, p.Or)
	}
}

func quiSontIls() {
	fmt.Println(MAGENTA + "\n🎵 QUI SONT-ILS ?" + RESET)
	fmt.Println(CYAN_VIF + "Partie 2 : ABBA" + RESET)
	fmt.Println(ARGENT + "- Money Money Money" + RESET)
	fmt.Println(ARGENT + "- Gimme! Gimme! Gimme!" + RESET)
	fmt.Println(ARGENT + "- On and On and On" + RESET)
	fmt.Println(CYAN_VIF + "Partie 3 : Références cinéma" + RESET)
	fmt.Println(ARGENT + "- Ready Player One, A.I., Duel" + RESET)
}
