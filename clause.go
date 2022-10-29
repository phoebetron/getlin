package getlin

type Clause interface {
	// Automa returns the number of TAs this Clause is made of.
	Automa() int
	// Metric returns the Clause's instance of a telemetric like object
	// maintaining state of certain runtime behaviour.
	Metric() Metric
	// Search produces the Clause's predicted label as an output vector, e.g.
	// [1] for predicting a true result. Search simply collects all literals
	// included by its TAs and returns the normalized vector as computed by the
	// configurable normalization method, e.g. ANDing.
	Search(Vector) bool
	// Update evolves the Clause's TAs towards more desirable patterns. Update
	// applies stochastic feedback activation, for which different activation
	// methods are configurable. Different feedback mechanism may be
	// configurable in the future or may even be learnable at some point.
	Update(Vector)
}
