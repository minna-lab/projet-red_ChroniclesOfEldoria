package main

import "fmt"

func afficherBanniere() {
	fmt.Println(POURPRE + GRAS + `
  ██████╗██╗  ██╗██████╗  ██████╗ ███╗   ██╗██╗ ██████╗██╗     ███████╗███████╗
 ██╔════╝██║  ██║██╔══██╗██╔═══██╗████╗  ██║██║██╔════╝██║     ██╔════╝██╔════╝
 ██║     ███████║██████╔╝██║   ██║██╔██╗ ██║█ ║██║     ██║     █████╗  ███████╗
 ██║     ██╔══██║██╔══██╗██║   ██║██║╚██╗██║██║██║     ██║     ██╔══╝  ╚════██║
 ╚██████╗██║  ██║██║  ██║╚██████╔╝██║ ╚████║██║╚██████╗███████╗███████╗███████║
  ╚═════╝╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═══╝╚═╝ ╚═════╝╚══════╝╚══════╝╚══════╝` + RESET)

	fmt.Println(OR + GRAS + `
           ██████╗ ███████╗    ███████╗██╗     ██████╗  ██████╗ ██████╗ ██╗ █████╗
           ██╔══██╗██╔════╝    ██╔════╝██║     ██╔══██╗██╔═══██╗██╔══██╗██║██╔══██╗
           ██║  ██║█████╗      █████╗  ██║     ██║  ██║██║   ██║██████╔╝██║███████║
           ██║  ██║██╔══╝      ██╔══╝  ██║     ██║  ██║██║   ██║██╔══██╗██║██╔══██║
           ██████╔╝███████╗    ███████╗███████╗██████╔╝╚██████╔╝██║  ██║██║██║  ██║
           ╚═════╝ ╚══════╝    ╚══════╝╚══════╝╚═════╝  ╚═════╝ ╚═╝  ╚═╝╚═╝╚═╝  ╚═╝` + RESET)

	fmt.Println()
	afficherPrologue()
}

func afficherPrologue() {
	fmt.Println(POURPRE + "════════════════════════════════════════════════════════" + RESET)
	fmt.Println(OR + GRAS + "  ✦  PROLOGUE  ✦" + RESET)
	fmt.Println(POURPRE + "════════════════════════════════════════════════════════" + RESET)
	fmt.Println()
	fmt.Println(OR + `  Il était une fois un royaume appelé Eldoria.
  Le peuple travaillait ensemble, les champs étaient verts, et
  la magie aidait chacun à vivre mieux.

  Mais un jour, la pierre d'or qui protégeait le royaume fut volée.
  Sans elle, les rivières se sont refroidies et la peur a commencé.

  Votre mission est simple : retrouver trois fragments de la pierre,
  rendre Eldoria lumineuse et ramener la paix.` + RESET)

	fmt.Println()
	fmt.Println(POURPRE + "════════════════════════════════════════════════════════" + RESET)
	fmt.Println()
}
