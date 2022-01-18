package ac_058

func ReverseLeftWords(s string, n int) string {
	b := []byte(s)
	tempBytes := make([]byte, 0)
	result := make([]byte, 0)
	for i := 0; i < len(b); i++ {
		if i < n {
			tempBytes = append(tempBytes, b[i])
		} else {
			result = append(result, b[i])
		}
	}
	for i := 0; i < len(tempBytes); i++ {
		result = append(result, tempBytes[i])
	}
	return string(result)
}
