package strings

// Erwartet einen String s und liefert einen neuen String,
// in dem jeder Buchstabe aus s zweimal hintereinander vorkommt.
func DuplicateChars(s string) string {

	stringcat := ""

	for _, char := range s {

		stringcat += string(char)
		stringcat += string(char)

	}

	return stringcat
}
