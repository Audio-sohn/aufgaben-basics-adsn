package contains

// Erwartet einen String und eine Liste von Strings.
// Gibt true zurück, wenn der String in der Liste enthalten ist, ansonsten false.
func Contains(strings []string, search string) bool {

	if len(strings) <= 0 { //leeres array catchen

		return false

	}

	for _, speci := range strings { //iterieren und prüfen

		if speci == search {

			return true

		}

	}

	return false
}
