package strings

// Erwartet einen String und prüft, ob er korrekte Klammerpaare enthält.
// D.h. der Eingabestring enthält Klammern '(' und ')'.
// Die Funktion soll nun prüfen, ob der String für jede öffnende Klammer auch eine
// schließende Klammer enthält.
// Außerdem darf es keine schließenden Klammern geben, für die es nicht vorher eine
// passende öffnende Klammer gegeben hat.
// Die Funktion soll true liefern, falls der String korrekt geklammert ist.
func CheckParentheses(s string) bool {
	// TODO
	return false
}

// HINWEIS
// Laufen Sie in einer for-Schleife über alle Buchstaben des Strings.
// Jedes Mal, wenn Sie eine öffnende Klammer finden, erhöhen Sie einen counter um 1.
// Jedes Mal, wenn Sie eine schließende Klammer finden, verringern Sie den counter um 1.
// Wenn der counter negativ wird, ist der String nicht korrekt geklammert.
// Wenn der counter am Ende der Schleife nicht 0 ist, ist der String nicht korrekt geklammert.
