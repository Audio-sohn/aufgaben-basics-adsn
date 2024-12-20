package make_unique

import "fmt"

// Erwartet eine Liste von Strings.
// Hängt Zahlen an alle mehrfach vorkommenden Strings an, um sie eindeutig zu machen.
// s1 := []string{"a", "b", "a", "c", "b", "a"}

func MakeUnique(strings []string) {

	result := strings //return array

	//durch das array iterieren und an die jeweilige stelle

	//geschachtelte for schleife, zuerst wird das erste element gespeichert und mit allen
	//anderen elementen verglichen, gibt es dopplungen wird in das array an dieser stelle
	//das element mit dem "unique" element überschrieben

	for _, speci := range strings {

		//counter reset bei neuer iteration
		count := 1

		for j, comp := range strings {

			if comp == speci {

				//wenn erstmals gefunden nur den string ohne counter printen und an die stelle im array schreiben
				//eigentlich ändert sich hier nichts, ist aber notwendig um sich die "result" list aufzubauen
				if count == 1 {

					result[j] = comp
					count++
					//continue

				} else {

					//wenn weitere gefunden, dann string zusammenstellen und an die stelle im array schreiben
					result[j] = fmt.Sprintf("%v_%v", comp, count)
					count++

				}
			}
		}
	}
	//final dann die "original" liste überschreiben
	strings = result

}

// REMARKS
// - Dies ist eine ähnliche Aufgabe wie das Zählen von Vorkommen in einer Liste.
//   Die innere Schleife macht i.W. das gleiche wie die Funktion `Count` aus der Aufgabe `count`.
//   Zusätzlich wird der Wert von `count` an den String angehängt, um ihn eindeutig zu machen.
