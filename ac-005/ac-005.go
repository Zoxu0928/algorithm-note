package ac_005

//遍历添加
func replaceSpace(s string) string {
	bytes := []byte(s)
	result := make([]byte, 0)
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == ' ' {
			result = append(result, []byte("%20")...)
		} else {
			result = append(result, bytes[i])
		}
	}
	return string(result)
}
