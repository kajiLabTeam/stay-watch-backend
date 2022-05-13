package util

type Util struct{}

func (Util) SliceUniqueString(target []string) (unique []string) {
	m := map[string]bool{}

	for _, v := range target {
		if !m[v] {
			m[v] = true
			unique = append(unique, v)
		}
	}

	return unique
}

func (Util) SliceUniqueNumber(target []int64) (unique []int64) {
	m := map[int64]bool{}

	for _, v := range target {
		if !m[v] {
			m[v] = true
			unique = append(unique, v)
		}
	}

	return unique
}

//配列の中に特定の文字列が含まれているか
func (Util) ArrayStringContains(target []string, str string) bool {
	for _, v := range target {
		if v == str {
			return true
		}
	}
	return false
}
