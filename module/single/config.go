package single

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/random/native"
)

type Config struct {
	// Cla is the number of Clauses to be managed by this module, where half of
	// them will be assigned positive voting power, that is, voting rights for
	// 1, and half of them negative voting power, that is, voting rights for 0.
	//
	//           +               -
	//
	//     C   C   C   C   C   C   C   C
	//
	Cla int
	// Fre is the probability used to calculate the feedback frequency. See
	// clause.Config for more information.
	Fre float32
	// Inp is the length of the input vector received. This number is equivalent
	// to the number of Tsetlin Automata (TAs) managed by each and every Clause
	// within this module.
	//
	//     I    [0 1 1 0 1 0 1 1 1]
	//
	Inp int
	// Ran provides randomization primitives for applying stochastic feedback.
	Ran getlin.Random
}

func (c Config) Ensure() Config {
	if c.Ran == nil {
		c.Ran = native.New(native.Config{})
	}

	return c
}

func (c Config) Verify() {
	if c.Cla == 0 {
		panic("Config.Cla must not be empty")
	}
	if c.Fre == 0 {
		panic("Config.Fre must not be empty")
	}
	if c.Inp == 0 {
		panic("Config.Inp must not be empty")
	}
	if c.Ran == nil {
		panic("Config.Ran must not be empty")
	}
}
