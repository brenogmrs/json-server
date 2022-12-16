package utils

func ErrorChecker(err error) {
	if err != nil {
		panic(err)
	}
}
