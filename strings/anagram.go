package strings

import (
	"slices"
	"strings"
)

// Erwartet zwei Strings s1 und s2 und prüft, ob die beiden Anagramme voneinander sind.
// Ein Anagramm von s1 ist ein String, der exakt die gleichen Buchstaben wie s1
// enthält, aber nicht unbedingt in der gleichen Reihenfolge.
func IsAnagram(s1, s2 string) bool {

	if len(s1) != len(s2) { // error catchen, wenn unterschiedlich lang, dann kein anagramm!
		return false
	}

	s1_slice := []byte(s1)
	s2_slice := []byte(s2)

	slices.Sort(s1_slice)
	slices.Sort(s2_slice)

	s1_string := string(s1_slice)
	s2_string := string(s2_slice)

	return s1_string == s2_string
}

// Erwartet zwei Strings s1 und s2 und prüft, ob die beiden Anagramme voneinander sind.
// Ein Anagramm von s1 ist ein String, der exakt die gleichen Buchstaben wie s1
// enthält, aber nicht unbedingt in der gleichen Reihenfolge.
// Diese Funktion soll keinen Unterschied zwischen Groß- und Kleinschreibung machen.
func IsAnagramIgnoreCase(s1, s2 string) bool {

	if len(s1) != len(s2) { // error catchen, wenn unterschiedlich lang, dann kein anagramm!
		return false
	}

	s1_lower := strings.ToLower(s1)
	s2_lower := strings.ToLower(s2)

	s1_slice := []byte(s1_lower)
	s2_slice := []byte(s2_lower)

	slices.Sort(s1_slice)
	slices.Sort(s2_slice)

	return string(s1_slice) == string(s2_slice)

}
