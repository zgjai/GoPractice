package mystrings

func FieldsFunc(s string, f func(rune) bool) []string {
	// 先遍历一遍求出结果slice的len
	n := 0
	inField := false
	for _, r := range s {
		wasInField := inField
		inField = !f(r)
		if !wasInField && inField {
			n++
		}
	}
	
	// make 大小为n的slice
	a := make([]string, n)
	
	// 分隔s
	// fieldStart用于保存一个符合的string的起始位置，为-1时表示尚未找到起始位置
	fieldStart := -1
	// na表示当前应填充的a的index位置
	na := 0
	for i, r := range s {
		inField = f(r)
		if inField && fieldStart != -1 {
			a[na] = s[fieldStart:i]
			na++
			fieldStart = -1
		}
		if !inField && fieldStart == -1 {
			fieldStart = i
		}
	}
	if fieldStart != -1 {
		a[na] = s[fieldStart:]
	}
	return a
}
