package modules

// Contains returns true if an element is inside an array
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Abs returns the absolute value of x
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// CalculateAgeGuessResponse returns an appropriate response for a guessed age
func CalculateAgeGuessResponse(realAge int, guessedAge int) (bool, string) {
	diff := Abs(realAge - guessedAge)
	if diff == 0 {
		return true, "Respuesta correcta"
	} else if diff < 2 {
		return false, "Demasiado cerca"
	} else if diff < 5 {
		return false, "Muy cerca"
	} else if diff < 10 {
		return false, "Sigue intentando"
	}
	return false, "Dedícate a otra cosa"
}
