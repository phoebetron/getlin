package graphs

func (m *Module) Automa() int {
	var aut int

	for _, x := range m.mod {
		for _, y := range x {
			aut += y.Automa()
		}
	}

	return aut
}
