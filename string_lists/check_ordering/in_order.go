package check_ordering

// Erwartet eine Liste von Strings und zwei einzelne Strings.
// Überprüft, ob die beiden Strings in der Liste in der gegebenen Reihenfolge vorkommen.
// Gibt `true` zurück, wenn das der Fall ist, ansonsten `false`.
func CheckOrdering(strings []string, first, second string) bool {

	//durch string iterieren, checken welcher char zuerst gefunden wurden
	firstFound := false // flags für stringentdeckung
	secondFound := false

	for _, char := range strings { // schleife

		//wenn zweites NACH dem ersten gefunden, dann false returnen
		if char == first && secondFound {

			return false

		}

		// wenn string erstmals gefunden, dann flag setzen
		if char == first && !firstFound {

			firstFound = true

		}

		// wenn string erstmals gefunden, dann flag setzen
		if char == second && !secondFound {

			secondFound = true

		}

	}

	//wenn alles i.O. und beide strings vorhanden, true returnen
	return firstFound && secondFound
}

// REMARKS
// - Diese Aufgabe ist eine komplexere Variante der Aufgabe "Prüfen, ob ein Element in einer Liste vorkommt".
// - Sie können die Lösung der einfachen Variante als Grundlage verwenden und diese entsprechend erweitern.
