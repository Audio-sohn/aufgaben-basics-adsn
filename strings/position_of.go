package strings

// Erwartet einen String s und einen Buchstaben c.
// Liefert die Position, an der c in s vorkommt.
// Liefert die Länge von s, falls c nicht in s vorkommt.
// Kommt c mehrfach vor, soll die erste Position geliefert werden.
func PositionOf(s string, c byte) int {

	//s_slice := []byte(s)

	for i, char := range s {

		if char == rune(c) {

			return i

		}

	}

	return len(s)
}
