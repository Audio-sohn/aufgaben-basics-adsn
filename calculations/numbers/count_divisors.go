package numbers

// Erwartet eine Zahl n >= 1 und liefert die Anzahl der Teiler dieser Zahl zurück.
func CountDivisors(n int) int {

	//durch alle kleineren zahlen iterieren und mod rechnen, wenn ergebnis 0 dann counter increment
	//counter ausgeben

	//counter bei 1 starten, da die zahl selbst schon ein teiler ist und die schleife
	//vorher aufhört. man könnte die schleife auch eins weiter laufen lassen
	counter := 1

	for i := 1; i < n; i++ {

		if n%i == 0 {

			counter++
		}
	}

	return counter
}
