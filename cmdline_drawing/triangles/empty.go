package triangles

import (
	"aufgaben-basics/cmdline_drawing/rectangles"
	"fmt"
)

func drawVerticalLinesTriang () {

	


}


// Erwartet eine Seitenlänge `length`.
// Zeichnet ein gleichschenkliges, rechtwinkliges Dreieck mit diesen Seitenlängen auf der Konsole.
// Das Dreieck soll komplett mit `#`-Zeichen gefüllt sein.
// Der rechte Winkel soll links unten liegen und die Seiten sollen
// vertikal bzw. horizontal verlaufen.
// Der Rand des Dreiecks soll aus `#`-Zeichen bestehen, der Innenraum soll leer sein.
func DrawEmptyTriangle(length int) {

	character := "#"

	for i := 1; i < length; i++ {
	
		

		switch(i) {
			
		default:
	


		}

		if i == length-1 {

			rectangles.DrawHorizontalLine(length, character)
		} else {

			rectangles.
		}





	}
		
	



	}

}

// REMARKS
// - Diese Aufgabe ist eine etwas komplexere Variante der Aufgabe "Gefülltes Dreieck" bzw. "Leeres Rechteck".
