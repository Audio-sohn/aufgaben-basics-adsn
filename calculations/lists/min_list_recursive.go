package lists

// Erwartet eine Liste von Zahlen und gibt das Minimum zurück.
// Wenn die Liste leer ist, wird 0 zurückgegeben.
// ZUSATZBEDINGUNG: Diese Funktion darf keine Schleife verwenden.
func MinListRecursive(nums []int) int {

	// error catchen bei leerer liste
	if len(nums) == 0 {
		return 0

		// wenn liste nur 1 lang offensichtliches ergebnis ausgeben
	} else if len(nums) == 1 {

		return nums[0]

	}

	//letztes element und vorletztes vergleichen, wenn vorletztes kleiner,
	//dann kleineres nach vorne kopieren
	if nums[len(nums)-1] < nums[len(nums)-2] {

		nums[len(nums)-2] = nums[len(nums)-1]

	}

	//und letzten eintrag verwerfen, dann die rekursion weiter führen
	return MinListRecursive(nums[:len(nums)-1])

}

// REMARKS
// - Diese Funktion nennt man "rekursiv".
// - Rekursion ist ein größeres Thema, das in der Vorlesung separat behandelt wird.
//   Um die Denkweise frühzeitig zu üben, gibt es aber gelegentlich auch vorab Aufgaben dazu.
