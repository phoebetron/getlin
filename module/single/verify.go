package single

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/metric"
)

func (m *Module) Verify(vec [][2]getlin.Vector) getlin.Metric {
	var met getlin.Metric
	{
		met = metric.New(metric.Config{})
	}

	for _, x := range vec {
		for _, y := range x {
			var t uint8
			{
				t = y.Out().Raw()[0]
			}

			var p uint8
			{
				p = m.Search(y).Out().Raw()[0]
			}

			if t == 1 && p == 1 {
				met.Set().Mat(metric.TP, 1)
			}
			if t == 0 && p == 0 {
				met.Set().Mat(metric.TN, 1)
			}
			if t == 1 && p == 0 {
				met.Set().Mat(metric.FN, 1)
			}
			if t == 0 && p == 1 {
				met.Set().Mat(metric.FP, 1)
			}
		}
	}

	return met
}