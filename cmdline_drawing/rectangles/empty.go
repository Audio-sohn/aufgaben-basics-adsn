package rectangles

import "fmt"

func DrawHorizontalLine(length int, character string) {

	for i := 0; i < length-1; i++ { //iterieren
		fmt.Print(character)
	}

	fmt.Println((character)) //nach letztem character umbruch reinhauen
}

func drawVerticalLines(height int, width int, character string) {

	for i := 0; i < height-2; i++ { //durch rows iterieren

		fmt.Print(character)

		for j := 0; j < width-2; j++ {

			fmt.Print(" ")

		}
		fmt.Println(character)

	}

}

// Erwartet zwei Seitenlängen `height` und `width`.
// Zeichnet ein Rechteck mit diesen Seitenlängen auf der Konsole.
// Der Rand des Rechtecks soll aus `#`-Zeichen bestehen, der Innenraum soll leer sein.
func DrawEmptyRectangle(height, width int) {

	DrawHorizontalLine(width, "#")

	drawVerticalLines(height, width, "#")

	DrawHorizontalLine(width, "#")

}

// REMARKS
// - Diese Aufgabe ist eine etwas komplexere Variante der Aufgabe "Gefülltes Rechteck zeichnen".
//   Man geht dabei ähnlich vor, muss allerdings zusätzlich prüfen, ob man sich am Rand des Rechtecks befindet.
