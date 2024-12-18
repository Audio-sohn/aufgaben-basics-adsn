package strings

// Erwartet zwei Strings s und sep sowie eine Zahl n.
// Liefert einen String, der aus n Kopien von s besteht, die duch sep getrennt werden.
func ConcatN(s, sep string, n int) string {

	stringcat := "" //string, der returned wird

	for i := 0; i < n; i++ { //an den string pro iteration den "s" string appenden

		stringcat += s

		if i < n-1 { //wenn wir NICHT in der letzten iteration sind, dann auch "sep" appenden

			stringcat += sep

		}
	}

	return stringcat
}
