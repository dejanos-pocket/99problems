package assert

func Empty(err error) {
	if err != nil {
		panic(err)
	}
}
