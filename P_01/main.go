package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Spaceline          Days  Trip type Price")
	fmt.Println("========================================")
	for i := 0; i < 10; i++ {
		switch rand.Intn(3) {
		case 0:
			switch rand.Intn(2) {
			case 0:
				fmt.Println("Virgin Galactic    ", 62100000/(rand.Intn(16)+15)/3600/24, "  ", "Round-trip $  ", (rand.Intn(15)+36)*2)
			case 1:
				fmt.Println("Virgin Galactic    ", 62100000/(rand.Intn(16)+15)/3600/24, "  ", "One-way    $  ", rand.Intn(15)+36)
			}
		case 1:
			switch rand.Intn(2) {
			case 0:
				fmt.Println("SpaceX             ", 62100000/(rand.Intn(16)+15)/3600/24, "  ", "Round-trip $  ", (rand.Intn(15)+36)*2)
			case 1:
				fmt.Println("SpaceX             ", 62100000/(rand.Intn(16)+15)/3600/24, "  ", "One-way    $  ", rand.Intn(15)+36)
			}
		case 2:
			switch rand.Intn(2) {
			case 0:
				fmt.Println("Space Adventures   ", 62100000/(rand.Intn(16)+15)/3600/24, "  ", "Round-trip $  ", (rand.Intn(15)+36)*2)
			case 1:
				fmt.Println("Space Adventures   ", 62100000/(rand.Intn(16)+15)/3600/24, "  ", "One-way    $  ", rand.Intn(15)+36)
			}
		}
	}
}
