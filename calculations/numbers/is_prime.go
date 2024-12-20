package numbers

// Erwartet eine Zahl n und prüft, ob n eine Primzahl ist.
func IsPrime(n int) bool {

	//error catchen, 1 ist keine primzahl
	if n == 1 {

		return false

	}

	divisorFound := false

	//durch alle zahlen zwischen 2 und n iterieren
	for i := 2; i < n; i++ {

		//checken ob ein teiler von n dabei ist
		if n%i == 0 {

			divisorFound = true
		}

	}

	//negiertes ergebnis für aussage über primzahl returnen
	return !divisorFound
}
