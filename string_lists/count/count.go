package count

// Erwartet eine Liste von Strings sowie einen String, der gezählt werden soll.
// Liefer die Anzahl der Vorkommen des gesuchten Strings in der Liste.
func Count(strings []string, search string) int {

	count := 0

	for _, speci := range strings {

		if speci == search {

			count++

		}

	}

	return count
}
