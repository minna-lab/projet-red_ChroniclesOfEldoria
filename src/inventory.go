package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// afficherInfos affiche toutes les stats du héros
func afficherInfos(p *Personnage) {
	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Printf("║  🏰 FICHE DE %s\n", strings.ToUpper(p.Nom))
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Printf("║  Classe    : %s\n", p.Classe)
	fmt.Printf("║  Niveau    : %d  (XP: %d/%d)\n", p.Niveau, p.XPActuel, p.XPMax)
	fmt.Printf("║  HP        : %d/%d\n", p.HPActuel, p.HPMax)
	fmt.Printf("║  Mana      : %d/%d\n", p.ManaActuel, p.ManaMax)
	fmt.Printf("║  Or        : %d 🪙\n", p.Or)
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Println("║  ÉQUIPEMENTS :")
	tete := p.Equip.Tete
	if tete == "" {
		tete = "(aucun)"
	}
	torse := p.Equip.Torse
	if torse == "" {
		torse = "(aucun)"
	}
	pieds := p.Equip.Pieds
	if pieds == "" {
		pieds = "(aucun)"
	}
	fmt.Printf("║  🪖 Tête   : %s\n", tete)
	fmt.Printf("║  🧥 Torse  : %s\n", torse)
	fmt.Printf("║  👢 Pieds  : %s\n", pieds)
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Println("║  SORTS CONNUS :")
	for _, s := range p.Sorts {
		fmt.Printf("║  ✨ %s\n", s)
	}
	fmt.Println("╚══════════════════════════════════════════╝")
}

// accederInventaire affiche l'inventaire et permet d'utiliser un item
func accederInventaire(p *Personnage) {
	lecteur := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n╔══════════════════════════════════════╗")
		fmt.Println("║         🎒 INVENTAIRE                ║")
		fmt.Println("╠══════════════════════════════════════╣")
		if len(p.Inventaire) == 0 {
			fmt.Println("║  (inventaire vide)                   ║")
		} else {
			for i, item := range p.Inventaire {
				fmt.Printf("║  %d. %s\n", i+1, item)
			}
		}
		fmt.Printf("║  Place : %d/%d\n", len(p.Inventaire), p.InventaireMax)
		fmt.Println("╠══════════════════════════════════════╣")
		fmt.Println("║  0. Retour                           ║")
		fmt.Println("╚══════════════════════════════════════╝")

		var choix int
		fmt.Print("➤ Votre choix : ")
		fmt.Scan(&choix)
		lecteur.ReadString('\n')

		if choix == 0 {
			return
		}
		if choix < 1 || choix > len(p.Inventaire) {
			fmt.Println("❌ Choix invalide.")
			continue
		}
		item := p.Inventaire[choix-1]
		utiliserItem(p, item, choix-1)
	}
}

// utiliserItem redirige vers la bonne fonction selon l'item choisi
func utiliserItem(p *Personnage, item string, index int) {
	switch item {
	case "Potion de Vie":
		utiliserPotionVie(p, index)
	case "Potion de Poison":
		utiliserPotionPoison(p, index)
	case "Livre de Sort : Boule de Feu":
		apprendreSort(p, index)
	case "Chapeau de l'Aventurier", "Tunique de l'Aventurier", "Bottes de l'Aventurier":
		equiper(p, item, index)
	case "Augmentation d'Inventaire":
		agrandirInventaire(p, index)
	default:
		fmt.Printf("❌ '%s' ne peut pas être utilisé directement.\n", item)
	}
}

// utiliserPotionVie soigne le joueur de 50 HP
func utiliserPotionVie(p *Personnage, index int) {
	if p.HPActuel >= p.HPMax {
		fmt.Println("💚 Vos points de vie sont déjà au maximum !")
		return
	}
	p.Inventaire = retirerInventaire(p.Inventaire, index)
	p.HPActuel += 50
	if p.HPActuel > p.HPMax {
		p.HPActuel = p.HPMax
	}
	fmt.Printf("🧪 Vous buvez une Potion de Vie ! HP : %d/%d\n", p.HPActuel, p.HPMax)
}

// utiliserPotionPoison inflige 10 dégâts par seconde pendant 3 secondes
func utiliserPotionPoison(p *Personnage, index int) {
	p.Inventaire = retirerInventaire(p.Inventaire, index)
	fmt.Println("☠️  Vous avalez la Potion de Poison... Vous vous sentez mal !")
	for i := 1; i <= 3; i++ {
		time.Sleep(1 * time.Second)
		p.HPActuel -= 10
		if p.HPActuel < 0 {
			p.HPActuel = 0
		}
		fmt.Printf("💀 Poison — Tour %d/3 : HP %d/%d\n", i, p.HPActuel, p.HPMax)
	}
	estMort(p)
}

// estMort vérifie si le joueur est mort et le ressuscite à 50% si c'est le cas
func estMort(p *Personnage) {
	if p.HPActuel <= 0 {
		fmt.Println("\n💀 Vous êtes mort... Mais les dieux d'Eldoria vous rappellent à la vie !")
		p.HPActuel = p.HPMax / 2
		fmt.Printf("❤️  Ressuscité avec %d HP.\n", p.HPActuel)
	}
}

// apprendreSort apprend "Boule de Feu" si le joueur ne le connaît pas déjà
func apprendreSort(p *Personnage, index int) {
	sort := "Boule de Feu"
	for _, s := range p.Sorts {
		if s == sort {
			fmt.Printf("📖 Vous connaissez déjà '%s' !\n", sort)
			return
		}
	}
	p.Inventaire = retirerInventaire(p.Inventaire, index)
	p.Sorts = append(p.Sorts, sort)
	fmt.Printf("✨ Vous apprenez le sort '%s' !\n", sort)
}

// inventairePlein vérifie si l'inventaire est plein
func inventairePlein(p *Personnage) bool {
	return len(p.Inventaire) >= p.InventaireMax
}

// ajouterInventaire ajoute un item si l'inventaire n'est pas plein
func ajouterInventaire(p *Personnage, item string) bool {
	if inventairePlein(p) {
		fmt.Printf("🎒 Inventaire plein ! Impossible d'ajouter '%s'.\n", item)
		return false
	}
	p.Inventaire = append(p.Inventaire, item)
	return true
}

// retirerInventaire retire l'item à l'index donné
func retirerInventaire(inventaire []string, index int) []string {
	return append(inventaire[:index], inventaire[index+1:]...)
}

// chercherItem cherche un item dans l'inventaire
func chercherItem(p *Personnage, item string) (int, bool) {
	for i, v := range p.Inventaire {
		if v == item {
			return i, true
		}
	}
	return -1, false
}

// compterItem compte combien de fois un item apparaît dans l'inventaire
func compterItem(p *Personnage, item string) int {
	compteur := 0
	for _, v := range p.Inventaire {
		if v == item {
			compteur++
		}
	}
	return compteur
}

// retirerItemUneFois retire une seule occurrence d'un item
func retirerItemUneFois(p *Personnage, item string) bool {
	idx, trouve := chercherItem(p, item)
	if !trouve {
		return false
	}
	p.Inventaire = retirerInventaire(p.Inventaire, idx)
	return true
}

// bonusEquipement contient les bonus HP de chaque équipement
var bonusEquipement = map[string]int{
	"Chapeau de l'Aventurier": 10,
	"Tunique de l'Aventurier": 25,
	"Bottes de l'Aventurier":  15,
}

// equiper équipe un objet depuis l'inventaire
func equiper(p *Personnage, item string, index int) {
	bonus := bonusEquipement[item]
	var ancienItem string

	switch item {
	case "Chapeau de l'Aventurier":
		ancienItem = p.Equip.Tete
		if ancienItem != "" {
			p.HPMax -= bonusEquipement[ancienItem]
			ajouterInventaire(p, ancienItem)
			fmt.Printf("🔄 Vous déséquipez '%s'.\n", ancienItem)
		}
		p.Equip.Tete = item
	case "Tunique de l'Aventurier":
		ancienItem = p.Equip.Torse
		if ancienItem != "" {
			p.HPMax -= bonusEquipement[ancienItem]
			ajouterInventaire(p, ancienItem)
			fmt.Printf("🔄 Vous déséquipez '%s'.\n", ancienItem)
		}
		p.Equip.Torse = item
	case "Bottes de l'Aventurier":
		ancienItem = p.Equip.Pieds
		if ancienItem != "" {
			p.HPMax -= bonusEquipement[ancienItem]
			ajouterInventaire(p, ancienItem)
			fmt.Printf("🔄 Vous déséquipez '%s'.\n", ancienItem)
		}
		p.Equip.Pieds = item
	}

	p.Inventaire = retirerInventaire(p.Inventaire, index)
	p.HPMax += bonus
	fmt.Printf("✅ Vous équipez '%s' ! +%d HP max. (HP max : %d)\n", item, bonus, p.HPMax)
}

// agrandirInventaire augmente la capacité de l'inventaire de +10
func agrandirInventaire(p *Personnage, index int) {
	if p.NbAgrandissements >= 3 {
		fmt.Println("❌ Vous avez déjà agrandi votre inventaire au maximum (3 fois) !")
		return
	}
	p.Inventaire = retirerInventaire(p.Inventaire, index)
	p.InventaireMax += 10
	p.NbAgrandissements++
	fmt.Printf("🎒 Inventaire agrandi ! Nouvelle capacité : %d (agrandissement %d/3)\n", p.InventaireMax, p.NbAgrandissements)
}
