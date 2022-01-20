package goPackage

func Fun1(s string, slice []string) bool {
	for _, sliceItem := range slice {
		if sliceItem == s {
			return true
		}
	}
	return false
}

func Fun2(slice1 []string, slice2 []string) []string {
	var my_map = make(map[string]int)
	var res []string
	for _, val := range slice2 {
		my_map[val] = my_map[val] + 1
	}
	for _, val := range slice1 {
		if my_map[val] == 0 {
			res = append(res, val)
		}
	}
	return res
}

func Fun3(slice []string) []string {
	var my_map = make(map[string]int)
	var res []string
	for _, val := range slice {
		if my_map[val] == 0 {
			my_map[val] = 1
		}
	}
	for id, _ := range my_map {
		res = append(res, id)
	}
	return res
}
