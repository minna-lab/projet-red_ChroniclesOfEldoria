package main

import "fmt"

func CreateCharacter() Character {
    var name string
    var choice int

    fmt.Print("Entrez votre nom : ")
    fmt.Scanln(&name)

    fmt.Println("Choisissez votre classe :")
    fmt.Println("1. Humain (100 HP)")
    fmt.Println("2. Elfe (80 HP)")
    fmt.Println("3. Nain (120 HP)")
    fmt.Print("Votre choix : ")
    fmt.Scanln(&choice)

    var maxHP int

    switch choice {
    case 1:
        maxHP = 100
    case 2:
        maxHP = 80
    case 3:
        maxHP = 120
    default:
        fmt.Println("Choix invalide → Humain par défaut")
        maxHP = 100
    }

    character := Character{
        Name:    name,
        MaxHP:   maxHP,
        HP:      maxHP / 2,
        Gold:    100,
        Potions: 3,
        Spells:  []string{"Coup de Poing"},
    }

    return character
}