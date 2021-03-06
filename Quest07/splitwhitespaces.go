package main

func SplitWhiteSpaces(str string) []string {
	len := 0
	len2 := 0
	for _, i := range str {
		if i == ' ' || i == '\t' || i == '\n' {
			len++
		}
	}
	len++
	if str[0] == ' ' {
		len--
	}
	if str[len] == ' ' {
		len--
	}
	for i := range str {
		len2 = i
	}

	res := make([]string, len)
	check := true
	k := 0
	for i := 0; i <= len2; i++ {
		if str[i] == ' ' || str[i] == '\t' || str[i] == '\n' {
			if !check {
				k++
			}
			check = true
			continue
		}
		res[k] += string(str[i])
		check = false
	}
	return res
}
