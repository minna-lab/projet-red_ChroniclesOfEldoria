package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// initialiserPersonnage crée un personnage avec les valeurs données
func initialiserPersonnage(nom, classe string, niveau, hpMax, hpActuel int, inventaire []string) Personnage {
	return Personnage{
		Nom:               nom,
		Classe:            classe,
		Niveau:            niveau,
		HPMax:             hpMax,
		HPActuel:          hpActuel,
		Or:                100,
		Inventaire:        inventaire,
		InventaireMax:     10,
		Sorts:             []string{"Coup de Poing"},
		Equip:             Equipement{},
		NbAgrandissements: 0,
		Initiative:        10,
		XPActuel:          0,
		XPMax:             100,
		ManaActuel:        50,
		ManaMax:           50,
	}
}

// formaterNom met la première lettre en majuscule, le reste en minuscule
func formaterNom(nom string) string {
	if len(nom) == 0 {
		return nom
	}
	runes := []rune(strings.ToLower(nom))
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// estAlpha vérifie que la chaîne ne contient que des lettres
func estAlpha(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// creationPersonnage permet au joueur de créer son héros
func creationPersonnage() Personnage {
	lecteur := bufio.NewReader(os.Stdin)

	fmt.Println("\n╔══════════════════════════════════════╗")
	fmt.Println("║   ✨ CRÉATION DE VOTRE HÉROS ✨       ║")
	fmt.Println("╚══════════════════════════════════════╝")

	// Choix du nom
	var nom string
	for {
		fmt.Print("\n🖊  Entrez votre nom (lettres uniquement) : ")
		saisie, _ := lecteur.ReadString('\n')
		saisie = strings.TrimSpace(saisie)
		if len(saisie) == 0 {
			fmt.Println("❌ Le nom ne peut pas être vide.")
			continue
		}
		if !estAlpha(saisie) {
			fmt.Println("❌ Le nom ne doit contenir que des lettres.")
			continue
		}
		nom = formaterNom(saisie)
		break
	}

	// Choix de la classe
	fmt.Println("\n⚔️  Choisissez votre classe :")
	fmt.Println("  1. Humain  (100 HP)")
	fmt.Println("  2. Elfe    (80 HP)")
	fmt.Println("  3. Nain    (120 HP)")

	var choixClasse int
	var classe string
	var hpMax int
	for {
		fmt.Print("\n➤ Votre choix : ")
		fmt.Scan(&choixClasse)
		lecteur.ReadString('\n')
		switch choixClasse {
		case 1:
			classe = "Humain"
			hpMax = 100
		case 2:
			classe = "Elfe"
			hpMax = 80
		case 3:
			classe = "Nain"
			hpMax = 120
		default:
			fmt.Println("❌ Choix invalide, entrez 1, 2 ou 3.")
			continue
		}
		break
	}

	hpDepart := hpMax / 2
	joueur := initialiserPersonnage(nom, classe, 1, hpMax, hpDepart, []string{"Potion de Vie", "Potion de Vie", "Potion de Vie"})

	fmt.Printf("\n✅ Bienvenue dans les Chroniques d'Eldoria, %s le %s !\n", joueur.Nom, joueur.Classe)
	return joueur
}
