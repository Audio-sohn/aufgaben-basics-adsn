package strings

import "math"

// Erwartet zwei strings s1 und s2.
// Liefert einen String, der abwechselnd aus den Buchstaben des einen und des anderen
// Strings besteht.
func Zip(s1, s2 string) string {

	s1_slice := []rune(s1)
	s2_slice := []rune(s2)

	result := ""

	maxcount := int(math.Max(float64(len(s1)), float64(len(s2)))) // länge des längeren strings ermitteln

	for i := 0; i < maxcount; i++ {

		if i < len(s1) {

			result += string(s1_slice[i])

		}

		if i < len(s2) {

			result += string(s2_slice[i])

		}

	}

	//durch den längeren string iterieren, wenn einer von den strings "leer" ist, ignorieren

	return result
}
