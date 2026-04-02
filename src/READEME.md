# Chronicles of Eldoria ⚔️
> Mini RPG en ligne de commande (CLI) développé en Go — Projet RED

---

## 📖 Histoire

Il y a trois cents ans, les peuples d'Eldoria — Humains, Elfes et Nains — vivaient en paix sous la protection de l'**Arcane Primordial**, gemme source de toute magie.

Jusqu'au jour où **Zarveth, le Dragon Corrupteur**, brisa la gemme en trois éclats. Sans elle, la magie s'effrite, les récoltes meurent et les ténèbres avancent.

Une prophétie désigne un héros au sang ordinaire pour réunir les éclats et terrasser le dragon.

**Ce héros, c'est vous.**

---

## 👥 Équipe

| Membre | Rôle | Tâches |
|--------|------|--------|
| **Minna** | Combat, Équipements, Missions | Tâches 16 à 22 + Missions |
| **Mariam** | Personnage, Inventaire, Économie | Tâches 1 à 15 |

---

## 🎮 Fonctionnalités

### Partie 1 — Personnage & Inventaire *(Mariam, tâches 1–15)*
- Création de personnage interactive : choix du nom et de la race (Humain, Elfe, Nain)
- Affichage de la fiche du héros (HP, Mana, XP, Or, sorts, équipements)
- Système d'inventaire avec limite de places (extensible jusqu'à 3 fois)
- Potion de Vie : restaure 50 HP
- Potion de Poison : inflige 10 dégâts/seconde pendant 3 secondes
- Livre de Sort : apprend Boule de Feu (un sort ne peut être appris qu'une seule fois)
- Marchand d'Eldoria : achat d'items avec un système de pièces d'or
- Augmentation d'inventaire : +10 places via le marchand (max 3 fois, 30 🪙)

### Partie 2 — Équipements & Forgeron *(Minna, tâches 16–22)*
- Structure Equipement avec 3 slots : Tête, Torse, Pieds
- Forgeron d'Ironholt : fabrication d'équipements à partir de matériaux
- Équiper un item modifie les HP max (Chapeau +10, Tunique +25, Bottes +15)
- Si un slot est déjà occupé, l'ancien équipement revient dans la sacoche

### Partie 3 — Combat *(Minna, tâches 19–22)*
- Combat tour par tour contre un Gobelin d'Entraînement
- IA du gobelin : attaque normale chaque tour, attaque double tous les 3 tours
- Actions du joueur : Attaque basique, Lancer un sort, Utiliser l'inventaire, Fuir
- Système d'initiative : le combattant avec la plus haute initiative attaque en premier

### Missions bonus *(Minna)*
- **Mission 1** — Initiative : détermine qui commence le combat
- **Mission 2** — Expérience : gain d'XP après victoire, montée de niveau avec excès reporté
- **Mission 3** — Sorts en combat : Coup de Poing (8 dmg) et Boule de Feu (18 dmg)
- **Mission 4** — Mana : les sorts consomment du mana, usage impossible si insuffisant
- **Mission 5** — Bannière animée : titre ASCII, histoire dynamique bloc par bloc
- **Mission 6** — Easter egg : option "Qui sont-ils ?" dans le menu principal

---

## 🛒 Catalogue du Marchand

| Item | Prix | Effet |
|------|------|-------|
| Potion de Vie | 3 🪙 | +50 HP |
| Potion de Poison | 6 🪙 | −10 HP/s pendant 3s |
| Livre de Sort : Boule de Feu | 25 🪙 | Apprend Boule de Feu |
| Augmentation d'Inventaire | 30 🪙 | +10 emplacements |
| Fourrure de Loup | 4 🪙 | Matériau de crafting |
| Peau de Troll | 7 🪙 | Matériau de crafting |
| Cuir de Sanglier | 3 🪙 | Matériau de crafting |
| Plume de Corbeau | 1 🪙 | Matériau de crafting |

---

## 🚀 Lancement

**Prérequis :** [Go 1.22+](https://go.dev/dl/)

```bash
# Cloner le dépôt
git clone https://github.com/votre-compte/projet-red_ChroniclesOfEldoria-MINNA.git
cd projet-red_ChroniclesOfEldoria-MINNA/src

# Lancer directement
go run .

# Ou compiler puis exécuter
go build -o eldoria
./eldoria
```

---

## 🗂️ Structure du projet

```
projet-red_ChroniclesOfEldoria-MINNA/
├── src/
│   ├── main.go        # Point d'entrée
│   ├── structs.go     # Structures : Personnage, Monstre, Equipement
│   ├── couleurs.go    # Palette ANSI (6 couleurs thématiques)
│   ├── banner.go      # ASCII art animé + histoire dynamique
│   ├── character.go   # Création et initialisation du personnage
│   ├── inventory.go   # Inventaire, potions, sorts, équipements
│   ├── menus.go       # Menus : principal, marchand, forgeron
│   └── combats.go     # Combat tour par tour, IA, XP, sorts
├── go.mod
└── README.md
```

---

## 🎯 Thème

Univers **Fantasy Médiévale** — elfes, nains, dragons, sorts magiques et forgerons dans le royaume d'Eldoria.
