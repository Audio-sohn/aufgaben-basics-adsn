package rectangles

import "fmt"

// Erwartet zwei Seitenlängen `height` und `width`.
// Zeichnet ein Rechteck mit diesen Seitenlängen auf der Konsole.
// Die Zeichen für Rand und Füllung des Rechtecks werden als Parameter erwartet.

// func drawHorizontalLineflex(length int, character string) {

// 	for i := 0; i < length-1; i++ { //iterieren
// 		fmt.Print(character)
// 	}

// 	fmt.Println((character)) //nach letztem character umbruch reinhauen
// }

func DrawVerticalLinesflex(height int, width int, inner, outer string) {

	for i := 0; i < height-2; i++ { //durch rows iterieren

		fmt.Print(outer)

		for j := 0; j < width-2; j++ {

			fmt.Print(inner)

		}
		fmt.Println(outer)

	}

}

func DrawRectangle(height, width int, inner, outer string) {

	DrawHorizontalLine(width, outer)
	DrawVerticalLinesflex(height, width, inner, outer)
	DrawHorizontalLine(width, outer)

	// TODO
}

// REMARKS
// - Wenn Sie diese Aufgabe gelöst haben, können Sie die Aufgaben
//   für das Zeichnen von leeren und gefüllten Rechtecken sehr viel einfacher lösen.
