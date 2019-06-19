package template

type A struct {
	Value int
	Name string
}

func (*A) Parse() {
	panic("implement me")
}

type B struct {
	K int
	V string
}

func (*B) Parse() {
	panic("implement me")
}


