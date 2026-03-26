package main

import "fmt"

func main() {
    player := CreateCharacter()

    fmt.Println("\n=== PERSONNAGE ===")
    fmt.Println("Nom :", player.Name)
    fmt.Println("HP :", player.HP, "/", player.MaxHP)
    fmt.Println("Or :", player.Gold)
    fmt.Println("Potions :", player.Potions)
    fmt.Println("Sorts :", player.Spells)
}