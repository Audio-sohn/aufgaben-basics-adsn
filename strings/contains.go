package strings

// Erwartet einen String s und einen Buchstaben c.
// Prüft, ob c in s vorkommt.
func Contains(s string, c byte) bool {

	for i := 0; i < len(s); i++ {

		if s[i] == c { //mit c vergleichen
			return true
		}

	}

	return false
}
