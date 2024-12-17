package triangles

// Erwartet eine Seitenlänge `length`.
// Zeichnet ein gleichschenkliges, rechtwinkliges Dreieck mit diesen Seitenlängen auf der Konsole.
// Das Dreieck soll komplett mit `#`-Zeichen gefüllt sein.
// Der rechte Winkel soll links unten liegen und die Seiten sollen
// vertikal bzw. horizontal verlaufen.
func DrawSolidTriangle(length int) {

	inner, outer := "#", "#"

	for i := 1; i <= length; i++ {

		if i == length {

			DrawVerticalLineflex(i, outer, outer)

		} else {
			DrawVerticalLineflex(i, inner, outer)
		}

	}
}
