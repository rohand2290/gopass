package error_handling

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
