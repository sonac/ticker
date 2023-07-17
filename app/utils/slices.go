package utils

import "sort"

func RemoveDuplicates(s *[]string) {
	if len(*s) == 0 {
		return
	}

	sort.Strings(*s)
	i := 0
	for j := 1; j < len(*s); j++ {
		if (*s)[j] != (*s)[i] {
			i++
			(*s)[i] = (*s)[j]
		}
	}
	*s = (*s)[:i+1]
}
