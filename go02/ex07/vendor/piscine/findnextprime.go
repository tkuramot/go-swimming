package piscine

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	}
	if nb == 2 {
		return 2
	}
	if nb%2 == 0 {
		return FindNextPrime(nb + 1)
	}
	for i := 3; i*i <= nb; i = i + 2 {
		if nb%i == 0 {
			return FindNextPrime(nb + 2)
		}
	}
	return nb
}
