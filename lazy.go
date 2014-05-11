package gory

type next func() interface{}

/*
Used to set a lazily evaluated value on a Factory.

Example:

    factory["Created"] = gory.Lazy(func() interface{} {
        return time.Now()
    })

In the above example, time.Now() is not evaluated until the
Factory is built, using the Build() function.
*/
func Lazy(callback func() interface{}) next {
	return callback
}
