package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/lukamilost-arch/vremenske-postaje/vreme"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "vreme",
		Usage: "Program za vremenske postaje",

		Flags: []cli.Flag{
			&cli.Float64Flag{
				Name:  "minTemp",
				Value: -50,
				Usage: "Najnižja dovoljena temperatura",
			},
			&cli.Float64Flag{
				Name:  "maxTemp",
				Value: 50,
				Usage: "Najvišja dovoljena temperatura",
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Value: 3,
				Usage: "Najmanjše število meritev za statistiko",
			},
		},

		Action: func(ctx context.Context, cmd *cli.Command) error {

			minTemp := cmd.Float64("minTemp")
			maxTemp := cmd.Float64("maxTemp")
			stMeritev := cmd.Int("stMeritev")

			if minTemp >= maxTemp {
				return fmt.Errorf("Neveljaven temperaturni obseg")
			}

			if stMeritev < 1 {
				return fmt.Errorf("stMeritev mora biti vsaj 1")
			}

			var postaje map[string]vreme.Postaja
			postaje = make(map[string]vreme.Postaja)

			initPostaje(postaje, minTemp, maxTemp)

			// Testiranje funkcij
			vreme.DodajMeritev(postaje, "LJ02", 15.5, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "MB01", -220.0, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "PT01", 10.2, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "CE01", 25.3, minTemp, maxTemp)

			fmt.Println()
			vreme.IzpisiMeritev(postaje)

			fmt.Println()
			vreme.IzpisPorocila(postaje, stMeritev)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func initPostaje(postaje map[string]vreme.Postaja,
	minTemp, maxTemp float64) {

	vreme.DodajPostajo(postaje, "LJ01", "Kredarca", "Ljubljana")
	vreme.DodajPostajo(postaje, "MB01", "Tezno", "Maribor")
	vreme.DodajPostajo(postaje, "PT01", "Ptujska Gora", "Ptuj")
	vreme.DodajPostajo(postaje, "CE01", "Kozjansko", "Celje")
	vreme.DodajPostajo(postaje, "PO01", "Pstwjna", "Postojna")

	meritve := map[string][]float64{
		"LJ01": {12.5, 13.2, 14.1},
		"MB01": {10.1, 12.0},
		"PT01": {9.8, 10.5, 11.2},
		"CE01": {11.2, 12.1, 13.3},
		"PO01": {45.3, 44.2, 25.1},
	}

	for oznaka, temperature := range meritve {
		for _, temperatura := range temperature {
			vreme.DodajMeritev(postaje, oznaka,
				temperatura, minTemp, maxTemp)
		}
	}
}
