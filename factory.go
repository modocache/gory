package gory

/*
A Factory maps field names to values.
*/
type Factory map[string]interface{}

func (this Factory) copy() Factory {
	factory := make(Factory, 0)
	factory.merge(this)
	return factory
}

func (this Factory) merge(factory Factory) {
	for name, value := range factory {
		this[name] = value
	}
}
