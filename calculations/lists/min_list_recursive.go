package lists

// Erwartet eine Liste von Zahlen und gibt das Minimum zurück.
// Wenn die Liste leer ist, wird 0 zurückgegeben.
// ZUSATZBEDINGUNG: Diese Funktion darf keine Schleife verwenden.
func MinListRecursive(nums []int) int {
	// TODO
	return 0
}

// HINWEIS
// - Vergleichen Sie jeweils das erste Element mit dem Minimum der restlichen Elemente.
//   Sie können dazu, die Funktion `Min` verwenden, falls Sie Sie bereits implementiert haben.
// - Um das Minimum der restlichen Elemente zu berechnen, rufen Sie die Funktion
//   erneut auf und übergeben Sie eine kürzere Liste, in der das erste Element entfernt wurde.
// - Der Fall, dass die Liste leer ist, ist ein Spezialfall, den Sie separat (und vorab) behandeln müssen.

// REMARKS
// - Diese Funktion nennt man "rekursiv".
// - Rekursion ist ein größeres Thema, das in der Vorlesung separat behandelt wird.
//   Um die Denkweise frühzeitig zu üben, gibt es aber gelegentlich auch vorab Aufgaben dazu.
