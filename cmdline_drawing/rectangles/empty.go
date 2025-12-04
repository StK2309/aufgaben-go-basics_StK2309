package rectangles

// Erwartet zwei Seitenlängen `height` und `width`.
// Zeichnet ein Rechteck mit diesen Seitenlängen auf der Konsole.
// Der Rand des Rechtecks soll aus `#`-Zeichen bestehen, der Innenraum soll leer sein.
func DrawEmptyRectangle(height, width int) {
	// TODO
}

// HINWEIS
// - Nutzen Sie zwei ineinander geschachtelte Schleifen, um die Zeilen und Spalten des Rechtecks zu durchlaufen.
// - Die äußere Schleife sollte die Zeilen des Rechtecks durchlaufen, die innere Schleife die Spalten.
// - Prüfen Sie in jeder Iteration, ob Sie sich am Rand des Rechtecks befinden.
//   Dies ist dann der Fall, wenn die Zeile oder Spalte gleich 0 oder `height-1` bzw. `width-1` ist.

// REMARKS
// - Diese Aufgabe ist eine etwas komplexere Variante der Aufgabe "Gefülltes Rechteck zeichnen".
//   Man geht dabei ähnlich vor, muss allerdings zusätzlich prüfen, ob man sich am Rand des Rechtecks befindet.
