package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func effacer() { fmt.Print("\033[H\033[2J") }

func ecrireLent(s string, d time.Duration) {
	for _, c := range s {
		fmt.Print(string(c))
		time.Sleep(d)
	}
}

func afficherBanniere() {
	effacer()

	// Titre ligne par ligne
	for _, l := range []string{
		` ░█████╗░██╗░░██╗██████╗░░█████╗░███╗░░██╗██╗░█████╗░██╗░░░░░███████╗░██████╗`,
		` ██╔══██╗██║░░██║██╔══██╗██╔══██╗████╗░██║██║██╔══██╗██║░░░░░██╔════╝██╔════╝`,
		` ██║░░╚═╝███████║██████╔╝██║░░██║██╔██╗██║██║██║░░╚═╝██║░░░░░█████╗░░╚█████╗░`,
		` ╚█████╔╝██║░░██║██║░░██║╚█████╔╝██║░╚███║██║╚█████╔╝███████╗███████╗██████╔╝`,
	} {
		fmt.Println(POURPRE + l + RESET)
		time.Sleep(90 * time.Millisecond)
	}

	fmt.Println()
	ecrireLent(OR+"                        ─── of ───"+RESET, 25*time.Millisecond)
	fmt.Println("\n")
	for _, l := range []string{
		`          ███████╗██╗░░░░░██████╗░░█████╗░██████╗░██╗░█████╗░`,
		`          ██╔════╝██║░░░░░██╔══██╗██╔══██╗██╔══██╗██║██╔══██╗`,
		`          █████╗░░██║░░░░░██║░░██║██║░░██║██████╔╝██║███████║`,
		`          ██╔══╝░░██║░░░░░██║░░██║██║░░██║██╔══██╗██║██╔══██║`,
		`          ███████╗███████╗██████╔╝╚█████╔╝██║░░██║██║██║░░██║`,
		`          ╚══════╝╚══════╝╚═════╝░░╚════╝░╚═╝░░╚═╝╚═╝╚═╝░░╚═╝`,
	} {
		fmt.Println(OR + l + RESET)
		time.Sleep(90 * time.Millisecond)
	}
	fmt.Println()
	ecrireLent(POURPRE+"     ·  ≪  ⚔  Magie · Combats · Aventure  ⚔  ≫  ·"+RESET, 15*time.Millisecond)
	fmt.Println("\n")
	time.Sleep(300 * time.Millisecond)

	// Histoire en 3 blocs — Entrée pour avancer
	lecteur := bufio.NewReader(os.Stdin)
	blocs := []struct {
		couleur string
		titre   string
		texte   string
	}{
		{POURPRE, "I. L'ÂGE D'OR",
			"  Pendant trois cents ans, Humains, Elfes et Nains\n  vécurent en paix grâce à l'Arcane Primordial,\n  gemme sacrée source de toute magie à Eldoria."},
		{ROUGE, "II. LA CHUTE",
			"  Zarveth, le Dragon Corrupteur, brisa la gemme\n  en trois éclats. Sans elle, la magie s'effrite,\n  les récoltes meurent et les ténèbres avancent."},
		{OR, "III. VOTRE DESTIN",
			"  Une prophétie désigne un héros au sang ordinaire\n  pour réunir les éclats et terrasser le dragon.\n\n  Ce héros... c'est vous."},
	}

	for _, b := range blocs {
		effacer()
		ecrireLent(b.couleur+"  ✦  "+b.titre+"  ✦"+RESET, 20*time.Millisecond)
		fmt.Println("\n")
		ecrireLent(ARGENT+b.texte+RESET, 8*time.Millisecond)
		fmt.Println("\n")
		ecrireLent(ARGENT+"  [ Entrée pour continuer ]"+RESET, 10*time.Millisecond)
		fmt.Println()
		lecteur.ReadString('\n')
	}

	// Prompt final séparé
	effacer()
	ecrireLent(POURPRE+"  ┌──────────────────────────────────────────┐"+RESET, 8*time.Millisecond)
	fmt.Println()
	ecrireLent(OR+"  │   ✦  Prêt à sauver Eldoria, héros ?  ✦   │"+RESET, 12*time.Millisecond)
	fmt.Println()
	ecrireLent(OR+"  │      Appuyez sur Entrée pour commencer    │"+RESET, 12*time.Millisecond)
	fmt.Println()
	ecrireLent(POURPRE+"  └──────────────────────────────────────────┘"+RESET, 8*time.Millisecond)
	fmt.Println("\n")
	lecteur.ReadString('\n')
}
