package metric

import "github.com/phoebetron/getlin"

type Metric struct {
	get *getter
	set *setter
}

type shared struct {
	mat *matrix
	sta *states
}

func New(con Config) *Metric {
	{
		con.Verify()
	}

	var sha *shared
	{
		sha = &shared{
			mat: &matrix{},
			sta: &states{},
		}
	}

	return &Metric{
		get: &getter{sha: sha},
		set: &setter{sha: sha},
	}
}

func (m *Metric) Get() getlin.Getter {
	return m.get
}

func (m *Metric) Set() getlin.Setter {
	return m.set
}
