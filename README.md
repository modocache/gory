# gory

[![Build Status](https://drone.io/github.com/modocache/gory/status.png?refresh=1)](https://drone.io/github.com/modocache/gory/latest)
[![Coverage Status](https://coveralls.io/repos/modocache/gory/badge.png?branch=master)](https://coveralls.io/r/modocache/gory?branch=master)
[![GoDoc](https://godoc.org/github.com/modocache/gory?status.png)](https://godoc.org/github.com/modocache/gory)

Factories for your Go structs. Think [factory_girl](https://github.com/thoughtbot/factory_girl).

## Usage

To install, just `go get` it:

```sh
go get github.com/modocache/gory
```

gory is [fully documented](https://godoc.org/github.com/modocache/gory), but
below are some examples to get you started.

### Defining Factories

Define factories that may be used during any test. Works great in
a [global setup hook](http://onsi.github.io/ginkgo/#global_setup_and_teardown__and_).

```go
gory.Define("user", User{}, func(factory gory.Factory) {
    factory["FirstName"] = "John"
    factory["LastName"] = "Doe"
    factory["Admin"] = false

    // 'n' in email is incremented each time the factory is built
    factory["Email"] = gory.Sequence(func(n int) interface{} {
        return fmt.Sprintf("john-doe-%d@example.com", n)
    })

    // time.Now() is evaluated when the factory is built
    factory["Created"] = gory.Lazy(func() interface{} {
        return time.Now()
    })
})
```

See `gory_suite_test.go` for more examples of defining factories.

### Using Factories

```go
john := gory.Build("user").(*User)
fmt.Println(john.FirstName) // "John"

jane := gory.BuildWithParams("user", gory.Factory{
    "FirstName": "Jane"
}).(*User)
fmt.Println(jane.FirstName) // "Jane"
```

See `gory_test.go` for more examples of using factories.

## Coming Soon

- Aliases
- Dependent attributes
- Transient attributes
- Associations
- Inheritance
- Traits
- Callbacks
- ...and pretty much [anything else factory_girl can do](https://github.com/thoughtbot/factory_girl/blob/master/GETTING_STARTED.md).
