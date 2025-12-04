package strings

// Erwartet einen String s und liefert s rückwärts zurück.
func Reverse(s string) string {
	// TODO
	return ""
}

// HINWEIS
// Laufen Sie in einer for-Schleife rückwärts über alle Buchstaben des Strings.
// Fügen Sie in jedem Schleifendurchlauf den aktuellen Buchstaben an den result-String an.

// Erwartet zwei Strings s1 und s2 und prüft, ob der eine der andere umgedreht ist.
func IsReverse(s1, s2 string) bool {

	return s1 == Reverse(s2)

}

// HINWEIS
// Verwenden Sie die Funktion Reverse(), um s2 umzudrehen
// und prüfen Sie dann, ob s1 und das umgedrehte s2 gleich sind.

// Erwartet einen String s und prüft, ob s ein Palindrom ist.
// Ein Palindrom ist ein String, der vorwärts und rückwärts gelesen gleich ist.
func IsPalindrome(s string) bool {

	return IsReverse(s, s)

}

// HINWEIS
// Verwenden Sie die Funktion IsReverse(), um zu prüfen, ob s gleich seinem
// umgedrehten String ist.
