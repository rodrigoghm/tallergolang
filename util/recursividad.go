package util

func Potencia(b, e uint64) uint64 {
	if e == 0 {
		return 1
	}

	return b * Potencia(b, e-1)
}
