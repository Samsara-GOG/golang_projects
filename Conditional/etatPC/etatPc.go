package main

import (
	"fmt"
)

func main() {

	//Create conditionals for a flow chart, e.g. what to do/check if your computer won’t power on.

	var pece string = "allumé"
	var power string = "ON"
	var prise string = "branché"
	var etatPrise string = "OK"
	var courant string = "actif"
	var batterie string = "marche"

	if pece == "allumé" {
		if power == "ON" {
			if prise == "branché" {
				if etatPrise == "OK" {
					if courant == "actif" {
						if batterie == "marche" {
							fmt.Println("Votre PC fonctionne correctement")
						} else {
							fmt.Println("Remplacez l'alimentation de votre PC")
						}
					} else {
						fmt.Println("Vérifiez l'état d'installation de votre maison")
					}
				} else {
					fmt.Println("Remplacez votre prise d'alimentation de secteur.")
				}
			} else {
				fmt.Println("Branchez la prise d'alimentation de votre PC")
			}
		} else {
			fmt.Println("Allumez le PC")
		}
		println("Tout va bien.")
	}

}
