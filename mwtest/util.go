package mwtest

type vec2 struct {
	x, y uint
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
