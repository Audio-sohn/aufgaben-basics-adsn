package triangles

import (
	"fmt"
)

func DrawVerticalLineflex(width int, inner, outer string) {

	for i := 0; i < width; i++ { //durch rows iterieren

		if i == 0 || i == width-1 {

			fmt.Print(outer) //wenn am rand, outer char printen

		} else {

			fmt.Print(inner) //wenn nicht am rand, dann inner char printen

		}
	}
	fmt.Println("") //line abschließen

}

// Erwartet eine Seitenlänge `length`.
// Zeichnet ein gleichschenkliges, rechtwinkliges Dreieck mit diesen Seitenlängen auf der Konsole.
// Das Dreieck soll komplett mit `#`-Zeichen gefüllt sein.
// Der rechte Winkel soll links unten liegen und die Seiten sollen
// vertikal bzw. horizontal verlaufen.
// Der Rand des Dreiecks soll aus `#`-Zeichen bestehen, der Innenraum soll leer sein.
func DrawEmptyTriangle(length int) {

	outer := "#"
	inner := " "

	for i := 1; i <= length; i++ {

		if i == length {

			DrawVerticalLineflex(i, outer, outer)

		} else {
			DrawVerticalLineflex(i, inner, outer)
		}

	}
}

// REMARKS
// - Diese Aufgabe ist eine etwas komplexere Variante der Aufgabe "Gefülltes Dreieck" bzw. "Leeres Rechteck".
