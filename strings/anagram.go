package strings

// Erwartet zwei Strings s1 und s2 und prüft, ob die beiden Anagramme voneinander sind.
// Ein Anagramm von s1 ist ein String, der exakt die gleichen Buchstaben wie s1
// enthält, aber nicht unbedingt in der gleichen Reihenfolge.
func IsAnagram(s1, s2 string) bool {
	// TODO
	return false
}

// HINWEIS
// Lauft in einer for-Schleife über alle Buchstaben von s1.
// Prüfen Sie in jedem Schleifendurchlauf mittels der Funktion CountChar(),
// ob der aktuelle Buchstabe in s1 gleich oft vorkommt wie in s2.

// Erwartet zwei Strings s1 und s2 und prüft, ob die beiden Anagramme voneinander sind.
// Ein Anagramm von s1 ist ein String, der exakt die gleichen Buchstaben wie s1
// enthält, aber nicht unbedingt in der gleichen Reihenfolge.
// Diese Funktion soll keinen Unterschied zwischen Groß- und Kleinschreibung machen.
func IsAnagramIgnoreCase(s1, s2 string) bool {
	// TODO
	return false
}

// HINWEIS
// Verwenden Sie die Funktion strings.ToLower(), um s1 und s2 in Kleinbuchstaben
// umzuwandeln, bevor Sie die Funktion IsAnagram() aufrufen.
