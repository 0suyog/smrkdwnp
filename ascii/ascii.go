package ascii

func IsPunct(r rune) bool {
	if (r >= 32 && r <= 47) || (r >= 58 && r <= 64) || (r >= 91 && r <= 96) || (r >= 123 && r <= 126) {
		return true
	}
	return false
}
