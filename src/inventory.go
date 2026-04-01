package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// ============================================================
//  TÂCHE 3 : Affichage des informations du personnage
// ============================================================

func displayInfo(c *Character) {
	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Printf("║  🏰 FICHE DE %s\n", strings.ToUpper(c.Name))
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Printf("║  Classe    : %s\n", c.Class)
	fmt.Printf("║  Niveau    : %d  (XP: %d/%d)\n", c.Level, c.CurrentXP, c.MaxXP)
	fmt.Printf("║  HP        : %d/%d\n", c.CurrentHP, c.MaxHP)
	fmt.Printf("║  Mana      : %d/%d\n", c.CurrentMana, c.MaxMana)
	fmt.Printf("║  Or        : %d 🪙\n", c.Gold)
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Println("║  ÉQUIPEMENTS :")
	head := c.Equip.Head
	if head == "" {
		head = "(aucun)"
	}
	chest := c.Equip.Chest
	if chest == "" {
		chest = "(aucun)"
	}
	feet := c.Equip.Feet
	if feet == "" {
		feet = "(aucun)"
	}
	fmt.Printf("║  🪖 Tête   : %s\n", head)
	fmt.Printf("║  🧥 Torse  : %s\n", chest)
	fmt.Printf("║  👢 Pieds  : %s\n", feet)
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Println("║  SORTS CONNUS :")
	for _, s := range c.Skills {
		fmt.Printf("║  ✨ %s\n", s)
	}
	fmt.Println("╚══════════════════════════════════════════╝")
}

// ============================================================
//  TÂCHE 4 : Accès à l'inventaire
// ============================================================

func accessInventory(c *Character) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n╔══════════════════════════════════════╗")
		fmt.Println("║         🎒 INVENTAIRE                ║")
		fmt.Println("╠══════════════════════════════════════╣")
		if len(c.Inventory) == 0 {
			fmt.Println("║  (inventaire vide)                   ║")
		} else {
			for i, item := range c.Inventory {
				fmt.Printf("║  %d. %s\n", i+1, item)
			}
		}
		fmt.Printf("║  Place : %d/%d\n", len(c.Inventory), c.MaxInventory)
		fmt.Println("╠══════════════════════════════════════╣")
		fmt.Println("║  Choisissez un item à utiliser       ║")
		fmt.Println("║  0. Retour                           ║")
		fmt.Println("╚══════════════════════════════════════╝")

		var choice int
		fmt.Print("➤ Votre choix : ")
		fmt.Scan(&choice)
		reader.ReadString('\n')

		if choice == 0 {
			return
		}
		if choice < 1 || choice > len(c.Inventory) {
			fmt.Println("❌ Choix invalide.")
			continue
		}

		item := c.Inventory[choice-1]
		useItem(c, item, choice-1)
	}
}

// useItem utilise un item de l'inventaire selon son type
func useItem(c *Character, item string, index int) {
	switch item {
	case "Potion de Vie":
		takePot(c, index)
	case "Potion de Poison":
		poisonPot(c, index)
	case "Livre de Sort : Boule de Feu":
		spellBook(c, index)
	case "Chapeau de l'Aventurier", "Tunique de l'Aventurier", "Bottes de l'Aventurier":
		equipItem(c, item, index)
... (139lignes restantes)