package utils

func ErrCheck(e error) {
	if e != nil {
		panic(e)
	}
}
