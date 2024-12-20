package lists

// Erwartet eine Liste von Zahlen und gibt das Minimum zurück.
// Wenn die Liste leer ist, wird 0 zurückgegeben.

func MinList(nums []int) int {

	//error catching: wenn liste leer, dann 0 returnen
	if len(nums) == 0 {

		return 0

	}

	//erste stelle der liste als minimum annehmen
	minimum := nums[0]

	//durch liste iterieren, wenn kleinerer wert als minimum gefunden dann minimum überschreiben
	for _, speci := range nums {

		if speci < minimum {

			minimum = speci

		}
	}

	//minimum zurückgeben
	return minimum
}
