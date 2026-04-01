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
		}
		if choice < 1 || choice > len(shopItems) {
			fmt.Println("❌ Choix invalide.")
			continue
		}

		item := shopItems[choice-1]
		if c.Gold < item.Price {
			fmt.Printf("💸 Pas assez d'or ! Il vous faut %d 🪙 (vous avez %d 🪙).\n", item.Price, c.Gold)
			continue
		}
		if isInventoryFull(c) {
			fmt.Println("🎒 Inventaire plein ! Faites de la place avant d'acheter.")
			continue
		}
		c.Gold -= item.Price
		c.Inventory = append(c.Inventory, item.Name)
		fmt.Printf("✅ Vous achetez '%s' pour %d 🪙. Il vous reste %d 🪙.\n", item.Name, item.Price, c.Gold)
	}
}

// ============================================================
//  TÂCHE 15-17 : Forgeron
// ============================================================

type Recipe struct {
	Name         string
	GoldCost     int
	Materials    map[string]int
	Emoji        string
}

var recipes = []Recipe{
	{
		Name:     "Chapeau de l'Aventurier",
		GoldCost: 5,
		Materials: map[string]int{
			"Plume de Corbeau": 1,
			"Cuir de Sanglier": 1,
		},
		Emoji: "🪖",
	},
	{
		Name:     "Tunique de l'Aventurier",
		GoldCost: 5,
		Materials: map[string]int{
			"Fourrure de Loup": 2,
			"Peau de Troll":    1,
		},
		Emoji: "🧥",
	},
	{
		Name:     "Bottes de l'Aventurier",
		GoldCost: 5,
		Materials: map[string]int{
			"Fourrure de Loup": 1,
			"Cuir de Sanglier": 1,
		},
		Emoji: "👢",
	},
}

func blacksmithMenu(c *Character) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n╔══════════════════════════════════════════╗")
		fmt.Println("║   🔨 FORGERON D'ELDORIA                  ║")
		fmt.Println("╠══════════════════════════════════════════╣")
		fmt.Printf("║  Votre bourse : %d 🪙\n", c.Gold)
		fmt.Println("╠══════════════════════════════════════════╣")
		for i, r := range recipes {
			fmt.Printf("║  %d. %s %s (coût : %d 🪙)\n", i+1, r.Emoji, r.Name, r.GoldCost)
			for mat, qty := range r.Materials {
				have := countItem(c, mat)
				fmt.Printf("║     - %dx %s (vous avez : %d)\n", qty, mat, have)
			}
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
		}
		if choice < 1 || choice > len(recipes) {
			fmt.Println("❌ Choix invalide.")
			continue
		}

		r := recipes[choice-1]
		craftItem(c, r)
	}
}

func craftItem(c *Character, r Recipe) {
	// Vérifier l'or
	if c.Gold < r.GoldCost {
		fmt.Printf("💸 Pas assez d'or ! Il vous faut %d 🪙.\n", r.GoldCost)
		return
	}
	// Vérifier les matériaux
	for mat, qty := range r.Materials {
		if countItem(c, mat) < qty {
			fmt.Printf("❌ Matériaux insuffisants : il vous faut %dx %s.\n", qty, mat)
			return
		}
	}
	// Vérifier la place
	if isInventoryFull(c) {
		fmt.Println("🎒 Inventaire plein ! Impossible de fabriquer cet objet.")
		return
	}
	// Consommer les matériaux
	for mat, qty := range r.Materials {
		for i := 0; i < qty; i++ {
			removeItemOnce(c, mat)
		}
	}
	c.Gold -= r.GoldCost
	c.Inventory = append(c.Inventory, r.Name)
	fmt.Printf("⚒️  Vous fabriquez '%s' ! (-5 🪙)\n", r.Name)
}

// ============================================================
//  TÂCHE 16-17 : Équiper un objet
// ============================================================

var equipBonuses = map[string]int{
	"Chapeau de l'Aventurier": 10,
	"Tunique de l'Aventurier": 25,
	"Bottes de l'Aventurier":  15,
}

func equipItem(c *Character, item string, index int) {
	bonus := equipBonuses[item]
	var oldItem string

	switch item {
	case "Chapeau de l'Aventurier":
		oldItem = c.Equip.Head
		if oldItem != "" {
			c.MaxHP -= equipBonuses[oldItem]
			addInventory(c, oldItem)
			fmt.Printf("🔄 Vous déséquipez '%s'.\n", oldItem)
		}
		c.Equip.Head = item
	case "Tunique de l'Aventurier":
		oldItem = c.Equip.Chest
		if oldItem != "" {
			c.MaxHP -= equipBonuses[oldItem]
			addInventory(c, oldItem)
			fmt.Printf("🔄 Vous déséquipez '%s'.\n", oldItem)
		}
		c.Equip.Chest = item
	case "Bottes de l'Aventurier":
		oldItem = c.Equip.Feet
		if oldItem != "" {
			c.MaxHP -= equipBonuses[oldItem]
			addInventory(c, oldItem)
			fmt.Printf("🔄 Vous déséquipez '%s'.\n", oldItem)
		}
		c.Equip.Feet = item
	}

	c.Inventory = removeInventory(c.Inventory, index)
	c.MaxHP += bonus
	fmt.Printf("✅ Vous équipez '%s' ! +%d HP max. (HP max : %d)\n", item, bonus, c.MaxHP)
}

// ============================================================
//  MISSION 6 : Qui sont-ils ?
// ============================================================

func whoAreThey() {
	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║   🎵 QUI SONT-ILS ?                      ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Println("║  Partie 2 : ABBA                         ║")
	fmt.Println("║  (Two for the Price of One,              ║")
	fmt.Println("║   Money Money Money,                     ║")
	fmt.Println("║   Gimme! Gimme! Gimme!,                  ║")
	fmt.Println("║   On and On and On)                      ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Println("║  Partie 3 : Daft Punk / films cultes     ║")
	fmt.Println("║  (Fighter Squad, Ready Player One,       ║")
	fmt.Println("║   A.I., Duel)                            ║")
	fmt.Println("╚══════════════════════════════════════════╝")
}