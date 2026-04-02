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
		fmt.Println(ROUGE + l + RESET)
		time.Sleep(90 * time.Millisecond)
	}

	fmt.Println()
	ecrireLent(JAUNE+"                        ─── of ───"+RESET, 25*time.Millisecond)
	fmt.Println("\n")
	for _, l := range []string{
		`          ███████╗██╗░░░░░██████╗░░█████╗░██████╗░██╗░█████╗░`,
		`          ██╔════╝██║░░░░░██╔══██╗██╔══██╗██╔══██╗██║██╔══██╗`,
		`          █████╗░░██║░░░░░██║░░██║██║░░██║██████╔╝██║███████║`,
		`          ██╔══╝░░██║░░░░░██║░░██║██║░░██║██╔══██╗██║██╔══██║`,
		`          ███████╗███████╗██████╔╝╚█████╔╝██║░░██║██║██║░░██║`,
		`          ╚══════╝╚══════╝╚═════╝░░╚════╝░╚═╝░░╚═╝╚═╝╚═╝░░╚═╝`,
	} {
		fmt.Println(JAUNE + l + RESET)
		time.Sleep(90 * time.Millisecond)
	}
	fmt.Println()
	ecrireLent(ROUGE+"     ·  ≪  ⚔  Magie · Combats · Aventure  ⚔  ≫  ·"+RESET, 15*time.Millisecond)
	fmt.Println("\n")
	time.Sleep(300 * time.Millisecond)

	// Histoire en 3 blocs — Entrée pour avancer
	lecteur := bufio.NewReader(os.Stdin)
	blocs := []struct {
		couleur string
		titre   string
		texte   string
	}{
		{ROUGE, "I. L'ÂGE D'OR",
			"  Il était une fois un royaume appelé Eldoria.\n  Le peuple travaillait ensemble, les champs étaient verts,\n  et la magie aidait chacun à vivre mieux."},
		{ROUGE, "II. LA CHUTE",
			"  Mais un jour, la pierre d'or qui protégeait le royaume\n  fut volée. Sans elle, les rivières se sont refroidies\n  et la peur a commencé."},
		{JAUNE, "III. VOTRE DESTIN",
			"  Votre mission est simple : retrouver trois fragments\n  de la pierre, rendre Eldoria lumineuse et ramener la paix.\n\n  Ce héros... c'est vous."},
	}

	for _, b := range blocs {
		effacer()
		ecrireLent(b.couleur+"  ✦  "+b.titre+"  ✦"+RESET, 20*time.Millisecond)
		fmt.Println("\n")
		ecrireLent(BLANC+b.texte+RESET, 8*time.Millisecond)
		fmt.Println("\n")
		ecrireLent(BLANC+"  [ Entrée pour continuer ]"+RESET, 10*time.Millisecond)
		fmt.Println()
		lecteur.ReadString('\n')
	}

	// Prompt final séparé
	effacer()
	ecrireLent(JAUNE+"  Prêt à sauver Eldoria, héros ?"+RESET, 12*time.Millisecond)
	fmt.Println()
	ecrireLent(JAUNE+"  Appuyez sur Entrée pour commencer"+RESET, 12*time.Millisecond)
	fmt.Println("\n")
	lecteur.ReadString('\n')
}
