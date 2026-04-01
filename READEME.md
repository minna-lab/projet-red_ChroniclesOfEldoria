# Chronicles of Eldoria ⚔️

> Un RPG en ligne de commande (CLI) développé en Go — Projet RED

## 🏰 Présentation

**Chronicles of Eldoria** est un mini jeu de rôle en ligne de commande. Vous incarnez un héros (Humain, Elfe ou Nain) qui explore le royaume d'Eldoria, gère son inventaire, commerce avec des marchands, forge son équipement et affronte des monstres en combat tour par tour.

## 👥 Équipe

| Membre | Rôle |
|--------|------|
| **Minna** (chef d'équipe) | Personnage, Inventaire, Économie (Tâches 1–15) |
| **Mariam** | Équipements, Combat, Intégration (Tâches 16–22) |

## 🎮 Fonctionnalités

- ✅ Création de personnage (Humain / Elfe / Nain)
- ✅ Affichage des stats et équipements
- ✅ Système d'inventaire avec limite (extensible)
- ✅ Potions de vie et de poison
- ✅ Système de sorts (Coup de Poing, Boule de Feu)
- ✅ Marchand avec économie en pièces d'or
- ✅ Forgeron avec système de crafting
- ✅ Équipements avec bonus de HP
- ✅ Combat tour par tour avec IA gobelin
- ✅ Système d'expérience et montée de niveau
- ✅ Système de mana
- ✅ Initiative (qui attaque en premier)

## 🚀 Installation & Lancement

### Prérequis
- [Go](https://go.dev/dl/) 1.22 ou supérieur

### Lancer le jeu

```bash
# Cloner le repo
git clone https://github.com/votre-compte/projet-red_ChroniclesOfEldoria-MINNA.git
cd projet-red_ChroniclesOfEldoria-MINNA

# Lancer le jeu
cd src
go run .
```

### Compiler et exécuter

```bash
cd src
go build -o eldoria
./eldoria
```

## 🗂️ Structure du projet

```
projet-red_ChroniclesOfEldoria-MINNA/
├── src/
│   ├── main.go         # Point d'entrée
│   ├── structs.go      # Structures Character, Monster, Equipment
│   ├── character.go    # Création et initialisation du personnage
│   ├── inventory.go    # Inventaire, potions, sorts, équipements
│   ├── menus.go        # Menus principal, marchand, forgeron
│   ├── combat.go       # Combat tour par tour, IA, XP, sorts
│   ├── banner.go       # ASCII art du titre
│   └── go.mod
└── README.md
```

## 🎯 Thème

Univers **Fantasy Médiévale** — elfes, nains, gobelins, sorts magiques et forgerons dans le royaume d'Eldoria .