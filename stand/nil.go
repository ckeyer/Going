package main

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {

	var e *Error
	checkError(e)
}

type Error struct {
	errCode uint8
}

func (e *Error) Error() string {
	switch e.errCode {
	case 1:
		return "file not found"
	case 2:
		return "time out"
	case 3:
		return "permission denied"
	default:
		return "unknown error"
	}
}
