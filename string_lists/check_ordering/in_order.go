package check_ordering

// Erwartet eine Liste von Strings und zwei einzelne Strings.
// Überprüft, ob die beiden Strings in der Liste in der gegebenen Reihenfolge vorkommen.
// Gibt `true` zurück, wenn das der Fall ist, ansonsten `false`.
func CheckOrdering(strings []string, first, second string) bool {

	//durch string iterieren, checken welcher char zuerst gefunden wurden
	firstFound := false // flags für stringentdeckung
	secondFound := false

	for _, char := range strings { // schleife

		if char == first && secondFound { //wenn zweites NACH dem ersten gefunden, dann false returnen

			return false

		}

		if char == first && !firstFound { // wenn string erstmals gefunden, dann flag setzen

			firstFound = true

		}

		if char == second && !secondFound { // wenn string erstmals gefunden, dann flag setzen

			secondFound = true

		}

	}

	return firstFound && secondFound //wenn alles i.O. und beide strings vorhanden, true retzrnen
}

// REMARKS
// - Diese Aufgabe ist eine komplexere Variante der Aufgabe "Prüfen, ob ein Element in einer Liste vorkommt".
// - Sie können die Lösung der einfachen Variante als Grundlage verwenden und diese entsprechend erweitern.
