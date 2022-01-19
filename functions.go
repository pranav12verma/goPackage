package goPackage

func Fun1(s string, slice []string) bool {
	for _, sliceItem := range slice {
		if sliceItem == s {
			return true
		}
	}
	return false
}
