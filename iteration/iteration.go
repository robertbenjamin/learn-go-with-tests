package iteration

const REPEATS = 5

func Repeat(character string) string {
	var repeated string

	for i := 0; i < REPEATS; i++ {
		repeated += character
	}

	return repeated
}
