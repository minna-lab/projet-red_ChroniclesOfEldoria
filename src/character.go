package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// initialiserPersonnage crée un personnage avec les valeurs fournies
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

// creationPersonnage guide le joueur pour créer son héros
func creationPersonnage() Personnage {
	lecteur := bufio.NewReader(os.Stdin)

	fmt.Println(ROUGE + "\nFORGER VOTRE DESTINÉE" + RESET)
	fmt.Println(ROUGE + "\nLa Prophétie vous appelle. Qui êtes-vous ?" + RESET)

	// Saisie du nom
	var nom string
	for {
		fmt.Print(VERT + "\n  Votre nom (lettres uniquement) : " + RESET)
		saisie, _ := lecteur.ReadString('\n')
		saisie = strings.TrimSpace(saisie)
		if len(saisie) == 0 {
			fmt.Println(ROUGE + "  ✗ Le nom ne peut pas être vide." + RESET)
			continue
		}
		if !estAlpha(saisie) {
			fmt.Println(ROUGE + "  ✗ Lettres uniquement, sans espaces ni chiffres." + RESET)
			continue
		}
		nom = formaterNom(saisie)
		break
	}

	// Choix de la race
	fmt.Println(JAUNE + "\n  Choisissez votre lignée :" + RESET)
	fmt.Println(BLEU + "  1. Humain  (100 HP) — Né de la terre, forgé par la volonté" + RESET)
	fmt.Println(VERT + "  2. Elfe    ( 80 HP) — Fils des forêts, maître des arcanes" + RESET)
	fmt.Println(BLANC + "  3. Nain    (120 HP) — Enfant des montagnes, résistance de pierre" + RESET)

	var choixClasse int
	var classe string
	var hpMax int
	for {
		fmt.Print(VERT + "\n  Votre choix : " + RESET)
		fmt.Scan(&choixClasse)
		lecteur.ReadString('\n')
		switch choixClasse {
		case 1:
			classe, hpMax = "Humain", 100
		case 2:
			classe, hpMax = "Elfe", 80
		case 3:
			classe, hpMax = "Nain", 120
		default:
			fmt.Println(ROUGE + "  ✗ Entrez 1, 2 ou 3." + RESET)
			continue
		}
		break
	}

	joueur := initialiserPersonnage(nom, classe, 1, hpMax, hpMax/2, []string{"Potion de Vie", "Potion de Vie", "Potion de Vie"})

	fmt.Println()
	fmt.Printf(JAUNE+"  ✦ %s le %s, Eldoria a besoin de vous. ✦\n"+RESET, joueur.Nom, joueur.Classe)
	return joueur
}
