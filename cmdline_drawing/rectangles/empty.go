package rectangles

import "fmt"

// Erwartet zwei Seitenlängen `height` und `width`.
// Zeichnet ein Rechteck mit diesen Seitenlängen auf der Konsole.
// Der Rand des Rechtecks soll aus `#`-Zeichen bestehen, der Innenraum soll leer sein.
func DrawEmptyRectangle(height, width int) {
	// TODO
	if height < 2 || width < 2 {
		return
	}
	for i := 0; i < width; i++ {
		fmt.Print("#")
	}

	fmt.Println()

	for i := 0; i < height-2; i++ {
		fmt.Print("#")
		for j := 0; j < width-2; j++ {
			fmt.Print(" ")
		}
		fmt.Println("#")
	}

	for i := 0; i < width; i++ {
		fmt.Print("#")
	}

	fmt.Println()

}

// REMARKS
// - Diese Aufgabe ist eine etwas komplexere Variante der Aufgabe "Gefülltes Rechteck zeichnen".
//   Man geht dabei ähnlich vor, muss allerdings zusätzlich prüfen, ob man sich am Rand des Rechtecks befindet.
