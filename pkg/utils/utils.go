package utils

func PInt(n int) *int {
	return &n
}

func PString(s string) *string {
	return &s
}

func PError(e error) *error {
	return &e
}
