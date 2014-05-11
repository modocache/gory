package gory

/*
A function that returns a value to be set on a Factory.
n is incremented each time the Factory is built.
*/
type Sequencer func(n int) interface{}

/*
Used to set an sequenced value on a Factory.

Example:

    factory["Email"] = gory.Sequence(func(n int) interface{} {
        return fmt.Sprintf("john-doe-%d@example.com", n)
    })

Each time Build() is called, the sequenced value will use an
incremented value for n.
*/
func Sequence(sequencer Sequencer) next {
	n := 0
	return Lazy(func() interface{} {
		value := sequencer(n)
		n += 1
		return value
	})
}

/*
A Sequencer that simply returns the int value of n.
In other words, the first time the factory is built, it
will return 0, then 1, then 2, and so on.

Example:

    factory["NumberOfChildren"] = gory.Sequence(gory.IntSequencer)

*/
var IntSequencer Sequencer = func(n int) interface{} {
	return n
}
