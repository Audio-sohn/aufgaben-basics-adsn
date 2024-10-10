package rectangles

import "fmt"

// Erwartet eine Seitenlänge `length`.
// Zeichnet ein Quadrat mit dieser Seitenlänge auf der Konsole.
// Das Quadrat soll komplett mit `#`-Zeichen gefüllt sein.
func DrawSolidSquare(length int) {

	for row := 0; row < length; row++ {
		for col := 0; col < length; col++ {
			fmt.Print("#")
		}
		fmt.Println()
	}

}

// HINWEISS
// - Sie können diese Aufgabe als Spezialfall der Rechtecksaufgabe in `solid.go` betrachten.
// - Wenn Sie die Quadrat-Aufgabe hier vorher bearbeiten, dann ist sie eine Vorübung
//   für die Rechteck-Aufgabe.
// - Wenn Sie die Quadrat-Aufgabe hier dach der Rechtecksaufgabe bearbeiten,
//   dann können Sie sie in einer Zeile lösen, indem Sie die andere Funktion aufrufen.