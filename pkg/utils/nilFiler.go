package utils

// NilFiller ポインタを渡すとnilだった場合空にして返却する
func NilFiller[T any](in ...*T) []T {
	t := make([]T, len(in))
	for i, v := range in {
		if v == nil {
			n := new(T)
			t[i] = *n
		} else {
			t[i] = *v
		}
	}

	return t
}
