package ac_050

func firstUniqChar(s string) byte {
	if s == "" {
		return ' '
	}
	cnts := [26]int{}
	for i := range s{
		index := s[i] - 'a'
		cnts[index] ++
	}

	for i := range s{
		index := s[i] - 'a'
		if cnts[index] == 1 {
			return s[i]
		}
	}
	return ' '
}
