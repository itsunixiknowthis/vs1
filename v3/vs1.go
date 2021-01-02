package vs1

type T1 struct {
	x int
	y int
}

func (t T1) F() string {
	return "T1.F"
}

func (t T1) G() string {
	return "T1.G"
}
