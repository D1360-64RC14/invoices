package functions

func Reverse(text string) string {
	result := []rune(text)

	for s, e := 0, len(result)-1; s < len(result)/2; s, e = s+1, e-1 {
		result[s], result[e] = result[e], result[s]
	}

	return string(result)
}
