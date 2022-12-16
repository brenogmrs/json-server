package utils

func Search(value int, ids []float64) int {

	lowerLimit := 0
	higherLimit := len(ids) - 1

	for lowerLimit <= higherLimit {

		mid := lowerLimit + (higherLimit-lowerLimit)/2
		midValue := int(ids[mid])

		if midValue == value {
			return mid
		} else if midValue > value {
			higherLimit = mid - 1
		} else {
			lowerLimit = mid + 1
		}

	}

	return -1
}

func RemoveIndex(data []interface{}, index int) []interface{} {
	return append(data[:index], data[index+1:]...)
}
