package main

// Equipement représente les slots d'équipement du personnage
type Equipement struct {
	Tete  string
	Torse string
	Pieds string
}

// Personnage représente le héros du joueur
type Personnage struct {
	Nom               string
	Classe            string
	Niveau            int
	HPMax             int
	HPActuel          int
	Or                int
	Inventaire        []string
	InventaireMax     int
	Sorts             []string
	Equip             Equipement
	NbAgrandissements int
	Initiative        int
	XPActuel          int
	XPMax             int
	ManaActuel        int
	ManaMax           int
}

// Monstre représente un ennemi
type Monstre struct {
	Nom        string
	HPMax      int
	HPActuel   int
	Attaque    int
	Initiative int
}
