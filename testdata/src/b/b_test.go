package b

type T1 struct {
	Pub string
	pri string
}

func okT1() T1 {
	return T1{Pub: "Pub", pri: "pri"}
}

func ng1T1() T1 {
	return T1{
		"Pub",
		"pri",
	}
}
