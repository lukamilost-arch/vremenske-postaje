// Package vreme omogoča upravljanje vremenskih postaj,
// beleženje temperaturnih meritev in izračun statistike.
package vreme

import "fmt"

// Postaja predstavlja vremensko postajo z imenom, krajem in seznamom meritev.
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

// DodajPostajo doda novo vremensko postajo v zbirko postaj.
func DodajPostajo(postaje map[string]Postaja, oznaka, ime, kraj string) {
	postaje[oznaka] = Postaja{
		ime:     ime,
		kraj:    kraj,
		meritve: []float64{},
	}
}

// DodajMeritev doda novo temperaturno meritev k obstoječi vremenski postaji.
func DodajMeritev(postaje map[string]Postaja, oznaka string,
	temperatura, minTemp, maxTemp float64) {

	if maxTemp <= temperatura || temperatura <= minTemp {
		fmt.Printf("Napaka: Temperatura %.1f °C je izven obsega.\n", temperatura)
		return
	}

	_, ok := postaje[oznaka]
	if !ok {
		fmt.Printf("Napaka: Postaja z oznako %s ne obstaja.\n", oznaka)
		return
	} else {
		postaja := postaje[oznaka]
		postaja.meritve = append(postaja.meritve, temperatura)
		postaje[oznaka] = postaja
		fmt.Println("Meritev dodana.")
	}
}

// statistika izračuna najnižjo, najvišjo in povprečno
// temperaturo postaje z dano oznako.
// Vrne false, če postaja ne obstaja ali nima dovolj meritev.
func statistika(postaje map[string]Postaja, oznaka string,
	stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {

	_, ok = postaje[oznaka]

	if !ok {
		return
	}

	postaja := postaje[oznaka]

	if len(postaja.meritve) < stMeritev || len(postaja.meritve) == 0 {
		ok = false
		return
	}

	najnizja = postaja.meritve[0]
	najvisja = postaja.meritve[0]
	povprecna = 0

	for _, meritev := range postaja.meritve {
		if meritev < najnizja {
			najnizja = meritev
		}
		if meritev > najvisja {
			najvisja = meritev
		}
		povprecna += meritev
	}

	povprecna /= float64(len(postaja.meritve))

	return najnizja, najvisja, povprecna, ok
}

// IzpisiMeritev izpiše vse meritve za vse vremenske postaje.
func IzpisiMeritev(postaje map[string]Postaja) {
	fmt.Println("MERITVE:")

	for oznaka, postaja := range postaje {
		fmt.Printf("%s - %s (%s): %v\n", oznaka, postaja.ime, postaja.kraj, postaja.meritve)
	}
}

// IzpisPorocila izpiše najnižjo, najvišjo in povprečno
// temperaturo ter opis povprečne temperature za vsako postajo.
// Če postaja nima vsaj stMeritev meritev, izpiše opozorilo.
func IzpisPorocila(postaje map[string]Postaja, stMeritev int) {

	for oznaka, postaja := range postaje {
		najnizja, najvisja, povprecna, ok :=
			statistika(postaje, oznaka, stMeritev)

		status := ""

		switch {
		case povprecna >= 25:
			status = "Vroče"
		case povprecna >= 15 && povprecna < 25:
			status = "Toplo"
		case povprecna >= 5 && povprecna < 15:
			status = "Hladno"
		default:
			status = "Mrzlo"
		}

		if !ok {
			fmt.Printf("%s: premalo meritev.\n", postaja.ime)
			continue
		} else {
			fmt.Printf("%s: min: %.1f °C, max: %.1f °C, povprečje: %.1f °C -> %s\n",
				postaja.ime, najnizja, najvisja, povprecna, status)
		}
	}
}
