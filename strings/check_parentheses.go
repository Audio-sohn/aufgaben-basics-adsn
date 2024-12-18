package strings

// Erwartet einen String und prüft, ob er korrekte Klammerpaare enthält.
// D.h. der Eingabestring enthält Klammern '(' und ')'.
// Die Funktion soll nun prüfen, ob der String für jede öffnende Klammer auch eine
// schließende Klammer enthält.
// Außerdem darf es keine schließenden Klammern geben, für die es nicht vorher eine
// passende öffnende Klammer gegeben hat.
// Die Funktion soll true liefern, falls der String korrekt geklammert ist.
func CheckParentheses(s string) bool {

	open_cnt := 0

	s_slice := []byte(s) //

	for i := 0; i < len(s_slice); i++ {

		if s_slice[i] == '(' {

			open_cnt++

		} else if s_slice[i] == ')' {

			open_cnt--

		}

	}

	return open_cnt == 0
}
