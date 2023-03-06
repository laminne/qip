package utils

// NilFiller Stringのポインタを渡すとnilだった場合空文字列にして返却する
func NilFiller(in ...*string) []string {
	t := make([]string, 0)
	for _, v := range in {
		if v == nil {
			t = append(t, "")
		} else {
			t = append(t, *v)
		}
	}

	return t
}
