package graphs

func (m *Module) Clause() int {
	var cla int

	for _, x := range m.mpr.All() {
		for _, y := range x {
			cla += y.Clause()
		}
	}

	return cla
}
