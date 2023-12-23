package piscine

func Split(s, sep string) []string {
	sLen := len(s)
	sepLen := len(sep)

	var temp string
	result := []string{}
	toAdd := false

	for i := 0; i < sLen; i++ {

		if (i+sepLen < sLen) && (s[i:i+sepLen] == sep) {
			result = append(result, temp)
			i += sepLen - 1
			temp = ""
			continue
		} else if i+sepLen > sLen {
			toAdd = true
		}

		temp += string(s[i])
	}

	if toAdd {
		result = append(result, temp)
	}

	return result
}
