package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func afficherInfos(p *Personnage) {
	fmt.Println(POURPRE + "\nFICHE DE " + strings.ToUpper(p.Nom) + RESET)
	fmt.Printf(CYAN_VIF+"Classe    : %s\n"+RESET, p.Classe)
	fmt.Printf(CYAN_VIF+"Niveau    : %d  (XP: %d/%d)\n"+RESET, p.Niveau, p.XPActuel, p.XPMax)
	fmt.Printf(ROUGE_VIF+"HP        : %d/%d\n"+RESET, p.HPActuel, p.HPMax)
	fmt.Printf(BLEU_VIF+"Mana      : %d/%d\n"+RESET, p.ManaActuel, p.ManaMax)
	fmt.Printf(OR+"Or        : %d pièces d'or\n"+RESET, p.Or)
	fmt.Println(JAUNE_VIVE + "Sorts connus:" + RESET)
	for _, s := range p.Sorts {
		fmt.Printf(EMERAUDE+"║    → %s\n"+RESET, s)
	}
	fmt.Println(POURPRE + "╚══════════════════════════════════════════╝" + RESET)
}

func accederInventaire(p *Personnage) {
	lecteur := bufio.NewReader(os.Stdin)
	for {
		fmt.Println(CYAN_VIF + "\nINVENTAIRE" + RESET)
		if len(p.Inventaire) == 0 {
			fmt.Println(ARGENT + "(votre sac est vide, voyageur...)" + RESET)
		} else {
			for i, item := range p.Inventaire {
				fmt.Printf(CYAN_VIF+"%d. "+BLANC+"%s\n"+RESET, i+1, item)
			}
		}
		fmt.Printf(ARGENT+"Place : %d/%d\n"+RESET, len(p.Inventaire), p.InventaireMax)
		fmt.Println(CYAN_VIF + "Choisissez un item à utiliser" + RESET)
		fmt.Println(ARGENT + "0. 🚪 Retour au menu principal" + RESET)

		var choix int
		fmt.Print(EMERAUDE + "➤ Votre choix : " + RESET)
		fmt.Scan(&choix)
		lecteur.ReadString('\n')

		if choix == 0 {
			return
		}
		if choix < 1 || choix > len(p.Inventaire) {
			fmt.Println(ROUGE + "❌ Choix invalide." + RESET)
			continue
		}
		item := p.Inventaire[choix-1]
		utiliserItem(p, item, choix-1)
	}
}

func utiliserItem(p *Personnage, item string, index int) {
	switch item {
	case "Potion de Vie":
		utiliserPotionVie(p, index)
	case "Potion de Poison":
		utiliserPotionPoison(p, index)
	case "Livre de Sort : Boule de Feu":
		apprendreSort(p, index)
	case "Augmentation d'Inventaire":
		agrandirInventaire(p, index)
	default:
		fmt.Printf(ROUGE+"❌ '%s' ne peut pas être utilisé directement.\n"+RESET, item)
	}
}

func utiliserPotionVie(p *Personnage, index int) {
	if p.HPActuel >= p.HPMax {
		fmt.Println(VERT_VIF + "💚 Vos blessures sont déjà guéries, héros !" + RESET)
		return
	}
	avant := p.HPActuel
	p.Inventaire = retirerInventaire(p.Inventaire, index)
	p.HPActuel += 50
	if p.HPActuel > p.HPMax {
		p.HPActuel = p.HPMax
	}
	fmt.Printf(VERT_VIF+"🧪 Vous buvez une Potion de Vie ! HP : %d → %d/%d\n"+RESET, avant, p.HPActuel, p.HPMax)
}

func utiliserPotionPoison(p *Personnage, index int) {
	p.Inventaire = retirerInventaire(p.Inventaire, index)
	fmt.Println(ROUGE + "☠️  Vous avalez la Potion de Poison... Le venin se répand dans vos veines !" + RESET)
	for i := 1; i <= 3; i++ {
		time.Sleep(1 * time.Second)
		p.HPActuel -= 10
		if p.HPActuel < 0 {
			p.HPActuel = 0
		}
		fmt.Printf(ROUGE+"💀 Poison — Tour %d/3 : HP %d/%d\n"+RESET, i, p.HPActuel, p.HPMax)
	}
	estMort(p)
}

func estMort(p *Personnage) {
	if p.HPActuel <= 0 {
		fmt.Println(ROUGE + "\n💀 Vous tombez... Les dieux d'Eldoria refusent de vous laisser mourir !" + RESET)
		p.HPActuel = p.HPMax / 2
		fmt.Printf(VERT_VIF+"❤️  Ressuscité avec %d HP. Continuez à vous battre !\n"+RESET, p.HPActuel)
	}
}

func apprendreSort(p *Personnage, index int) {
	sort := "Boule de Feu"
	for _, s := range p.Sorts {
		if s == sort {
			fmt.Printf(JAUNE_VIVE+"📖 Vous maîtrisez déjà le sort '%s' !\n"+RESET, sort)
			return
		}
	}
	p.Inventaire = retirerInventaire(p.Inventaire, index)
	p.Sorts = append(p.Sorts, sort)
	fmt.Printf(EMERAUDE+"✨ Les runes du livre brillent... Vous apprenez '%s' !\n"+RESET, sort)
}

func inventairePlein(p *Personnage) bool {
	return len(p.Inventaire) >= p.InventaireMax
}

func ajouterInventaire(p *Personnage, item string) bool {
	if inventairePlein(p) {
		fmt.Printf(ROUGE+"🎒 Votre sac est plein ! Impossible d'ajouter '%s'.\n"+RESET, item)
		return false
	}
	p.Inventaire = append(p.Inventaire, item)
	return true
}

func retirerInventaire(inventaire []string, index int) []string {
	return append(inventaire[:index], inventaire[index+1:]...)
}

func chercherItem(p *Personnage, item string) (int, bool) {
	for i, v := range p.Inventaire {
		if v == item {
			return i, true
		}
	}
	return -1, false
}

func compterItem(p *Personnage, item string) int {
	compteur := 0
	for _, v := range p.Inventaire {
		if v == item {
			compteur++
		}
	}
	return compteur
}

func retirerItemUneFois(p *Personnage, item string) bool {
	idx, trouve := chercherItem(p, item)
	if !trouve {
		return false
	}
	p.Inventaire = retirerInventaire(p.Inventaire, idx)
	return true
}

func agrandirInventaire(p *Personnage, index int) {
	if p.NbAgrandissements >= 3 {
		fmt.Println(ROUGE + "❌ Les dieux refusent d'agrandir davantage votre sac (3 fois max) !" + RESET)
		return
	}
	p.Inventaire = retirerInventaire(p.Inventaire, index)
	p.InventaireMax += 10
	p.NbAgrandissements++
	fmt.Printf(VERT_VIF+"🎒 Votre sac magique s'agrandit ! Capacité : %d (agrandissement %d/3)\n"+RESET, p.InventaireMax, p.NbAgrandissements)
}
