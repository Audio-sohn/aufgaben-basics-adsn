package strings

// Erwartet zwei Strings s und t und prüft, ob t in s als Teilstring vorkommt.
func ContainsSubstring(s, t string) bool {

	if len(t) <= 0 || len(t) > len(s) { //error catchen wenn string leer oder substring länger als überstring
		return false
	}

	//s_slice := []byte(s) //string in array an chars konvertieren
	var s_substring string

	for i := 0; i <= len(s)-len(t); i++ {

		//aus dem string s einen string raustrennen der der länge von t entspricht
		//danach um einen character "weiterhangeln"

		s_substring = s[i : i+len(t)]

		if s_substring == t { //mit t vergleichen
			return true
		}

	}

	return false
}
