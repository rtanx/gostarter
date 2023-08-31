package utils

func FirstOrDefault[T any](defaultVal T, valList ...T) T {
	if len(valList) > 0 {
		return valList[0]
	}
	return defaultVal
}
