package main

import (
	"bufio"
	"fmt"
	"os"
)

// ============================================================
//  TÂCHE 6 : Menu principal
// ============================================================

func mainMenu(c *Character) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n╔══════════════════════════════════════════╗")
		fmt.Println("║    🏰 CHRONICLES OF ELDORIA 🏰           ║")
		fmt.Println("╠══════════════════════════════════════════╣")
		fmt.Printf("║  Héros : %-10s | HP : %d/%d\n", c.Name, c.CurrentHP, c.MaxHP)
		fmt.Printf("║  Or    : %d 🪙\n", c.Gold)
		fmt.Println("╠══════════════════════════════════════════╣")
		fmt.Println("║  1. Afficher les informations            ║")
		fmt.Println("║  2. Accéder à l'inventaire               ║")
		fmt.Println("║  3. Marchand                             ║")
		fmt.Println("║  4. Forgeron                             ║")
		fmt.Println("║  5. Combat d'entraînement                ║")
		fmt.Println("║  6. Qui sont-ils ?                       ║")
		fmt.Println("║  0. Quitter                              ║")
		fmt.Println("╚══════════════════════════════════════════╝")

		var choice int
		fmt.Print("➤ Votre choix : ")
		fmt.Scan(&choice)
		reader.ReadString('\n')

		switch choice {
		case 1:
			displayInfo(c)
		case 2:
			accessInventory(c)
		case 3:
			merchantMenu(c)
		case 4:
			blacksmithMenu(c)
		case 5:
			trainingFight(c)
		case 6:
			whoAreThey()
		case 0:
			fmt.Println("\n⚔️  Les Chroniques d'Eldoria vous attendent, brave héros. À bientôt !")
			return
		default:
			fmt.Println("❌ Choix invalide.")
		}
	}
}

// ============================================================
//  TÂCHE 7 & 14 : Marchand
// ============================================================

type ShopItem struct {
	Name  string
	Price int
	Emoji string
}

var shopItems = []ShopItem{
	{"Potion de Vie", 3, "🧪"},
	{"Potion de Poison", 6, "☠️"},
	{"Livre de Sort : Boule de Feu", 25, "📖"},
	{"Fourrure de Loup", 4, "🐺"},
	{"Peau de Troll", 7, "👹"},
	{"Cuir de Sanglier", 3, "🐗"},
	{"Plume de Corbeau", 1, "🪶"},
	{"Augmentation d'Inventaire", 30, "🎒"},
}

func merchantMenu(c *Character) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n╔══════════════════════════════════════════╗")
		fmt.Println("║   🛒 MARCHAND D'ELDORIA                  ║")
		fmt.Println("╠══════════════════════════════════════════╣")
		fmt.Printf("║  Votre bourse : %d 🪙\n", c.Gold)
		fmt.Println("╠══════════════════════════════════════════╣")
		for i, item := range shopItems {
			fmt.Printf("║  %d. %s %-25s %d 🪙\n", i+1, item.Emoji, item.Name, item.Price)
		}
		fmt.Println("╠══════════════════════════════════════════╣")
		fmt.Println("║  0. Retour                               ║")
		fmt.Println("╚══════════════════════════════════════════╝")

		var choice int
		fmt.Print("➤ Votre choix : ")
		fmt.Scan(&choice)
		reader.ReadString('\n')

		if choice == 0 {
			return
... (194lignes restantes)