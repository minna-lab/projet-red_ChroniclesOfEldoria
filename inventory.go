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
	case "Augmentation d'Inventaire":
		upgradeInventorySlot(c, index)
	default:
		fmt.Printf("❌ '%s' ne peut pas être utilisé directement.\n", item)
	}
}

// ============================================================
//  TÂCHE 5 : Potion de vie
// ============================================================

func takePot(c *Character, index int) {
	if c.CurrentHP >= c.MaxHP {
		fmt.Println("💚 Vos points de vie sont déjà au maximum !")
		return
	}
	c.Inventory = removeInventory(c.Inventory, index)
	c.CurrentHP += 50
	if c.CurrentHP > c.MaxHP {
		c.CurrentHP = c.MaxHP
	}
	fmt.Printf("🧪 Vous buvez une Potion de Vie ! HP : %d/%d\n", c.CurrentHP, c.MaxHP)
}

// ============================================================
//  TÂCHE 9 : Potion de poison
// ============================================================

func poisonPot(c *Character, index int) {
	c.Inventory = removeInventory(c.Inventory, index)
	fmt.Println("☠️  Vous avalez la Potion de Poison... Vous vous sentez mal !")
	for i := 1; i <= 3; i++ {
		time.Sleep(1 * time.Second)
		c.CurrentHP -= 10
		if c.CurrentHP < 0 {
			c.CurrentHP = 0
		}
		fmt.Printf("💀 Poison — Tour %d/3 : HP %d/%d\n", i, c.CurrentHP, c.MaxHP)
	}
	isDead(c)
}

// ============================================================
//  TÂCHE 8 : isDead
// ============================================================

func isDead(c *Character) {
	if c.CurrentHP <= 0 {
		fmt.Println("\n💀 Vous êtes mort... Mais les dieux d'Eldoria vous rappellent à la vie !")
		c.CurrentHP = c.MaxHP / 2
		fmt.Printf("❤️  Ressuscité avec %d HP.\n", c.CurrentHP)
	}
}

// ============================================================
//  TÂCHE 10 : Livre de sort
// ============================================================

func spellBook(c *Character, index int) {
	spell := "Boule de Feu"
	for _, s := range c.Skills {
		if s == spell {
			fmt.Printf("📖 Vous connaissez déjà '%s' !\n", spell)
			return
		}
	}
	c.Inventory = removeInventory(c.Inventory, index)
	c.Skills = append(c.Skills, spell)
	fmt.Printf("✨ Vous apprenez le sort '%s' !\n", spell)
}

// ============================================================
//  TÂCHE 12 : Limite d'inventaire
// ============================================================

func isInventoryFull(c *Character) bool {
	return len(c.Inventory) >= c.MaxInventory
}

// ============================================================
//  Helpers inventaire
// ============================================================

func addInventory(c *Character, item string) bool {
	if isInventoryFull(c) {
		fmt.Printf("🎒 Inventaire plein ! Impossible d'ajouter '%s'.\n", item)
		return false
	}
	c.Inventory = append(c.Inventory, item)
	return true
}

func removeInventory(inventory []string, index int) []string {
	return append(inventory[:index], inventory[index+1:]...)
}

func hasItem(c *Character, item string) (int, bool) {
	for i, v := range c.Inventory {
		if v == item {
			return i, true
		}
	}
	return -1, false
}

func countItem(c *Character, item string) int {
	count := 0
	for _, v := range c.Inventory {
		if v == item {
			count++
		}
	}
	return count
}

func removeItemOnce(c *Character, item string) bool {
	idx, found := hasItem(c, item)
	if !found {
		return false
	}
	c.Inventory = removeInventory(c.Inventory, idx)
	return true
}

// ============================================================
//  TÂCHE 18 : Augmenter la capacité de l'inventaire
// ============================================================

func upgradeInventorySlot(c *Character, index int) {
	if c.UpgradeCount >= 3 {
		fmt.Println("❌ Vous avez déjà agrandi votre inventaire au maximum (3 fois) !")
		return
	}
	c.Inventory = removeInventory(c.Inventory, index)
	c.MaxInventory += 10
	c.UpgradeCount++
	fmt.Printf("🎒 Inventaire agrandi ! Nouvelle capacité : %d (agrandissement %d/3)\n", c.MaxInventory, c.UpgradeCount)
}