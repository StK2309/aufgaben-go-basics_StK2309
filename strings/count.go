package strings

// Erwartet einen string s und zählt, wie oft der Buchstabe 'A' in s vorkommt.
func CountA(s string) int {
	// TODO
	return 0
}

// HINWEIS
// Laufen Sie in einer for-Schleife über alle Buchstaben des Strings.
// Prüfen Sie in jedem Schleifendurchlauf, ob der aktuelle Buchstabe ein 'A' ist.
// Wenn ja, erhöhen Sie result um 1.
// Alternativ können Sie auch die Funktion CountChar() verwenden,
// die Sie weiter unten implementieren sollen.

// Erwartet einen string s und einen Buchstaben c.
// Zählt, wie oft c in s vorkommt.
func CountChar(s string, c rune) int {

	result := 0
	for _, char := range s {
		if char == c {
			result++
		}
	}
	return result

}

// HINWEIS
// Laufen Sie in einer for-Schleife über alle Buchstaben des Strings.
// Prüfen Sie in jedem Schleifendurchlauf, ob der aktuelle Buchstabe gleich c ist.
// Wenn ja, erhöhen Sie result um 1.
