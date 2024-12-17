package triangles

// Erwartet eine Seitenlänge `length`.
// Zeichnet ein gleichschenkliges, rechtwinkliges Dreieck mit diesen Seitenlängen auf der Konsole.
// Der rechte Winkel soll links unten liegen und die Seiten sollen
// vertikal bzw. horizontal verlaufen.
// Die Zeichen für Rand und Füllung des Rechtecks werden als Parameter erwartet.
func DrawTriangle(length int, inner, outer string) {

	for i := 1; i <= length; i++ {

		if i == length {

			DrawVerticalLineflex(i, outer, outer)

		} else {
			DrawVerticalLineflex(i, inner, outer)
		}

	}
}
