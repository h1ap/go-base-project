package generic

type Number interface {
	int8 | int16 | int32 | float64 | int64
}

func Sum[T Number](v ...T) T {
	var sum T
	for _, v := range v {
		sum += v
	}
	return sum
}
