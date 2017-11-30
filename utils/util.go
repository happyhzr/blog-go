package utils

func IsEqualInt64s(a []int64, b []int64) bool {
	for _, i := range a {
		found := false
		for _, j := range b {
			if i == j {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
}
