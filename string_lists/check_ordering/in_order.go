package check_ordering

// Erwartet eine Liste von Strings und zwei einzelne Strings.
// Überprüft, ob die beiden Strings in der Liste in der gegebenen Reihenfolge vorkommen.
// Gibt `true` zurück, wenn das der Fall ist, ansonsten `false`.
func CheckOrdering(strings []string, first, second string) bool {

	ifirst := 0
	isecond := 0

	//error sofort catchen und abbrechen
	if len(strings) == 0 {

		return false

	}

	// i ist der zähler, speci der respektive wert an der stelle i im array "strings"
	// so wird durchiteriert ohne dass du manuell nochmal conditions setzen musst
	for i, speci := range strings {
		if speci == first {
			ifirst = i
		}
		if speci == second {
			isecond = i
		}
	}

	return ifirst != 0 && isecond != 0 && ifirst < isecond
}

// REMARKS
// - Diese Aufgabe ist eine komplexere Variante der Aufgabe "Prüfen, ob ein Element in einer Liste vorkommt".
// - Sie können die Lösung der einfachen Variante als Grundlage verwenden und diese entsprechend erweitern.
