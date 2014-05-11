/*
gory is a fixtures replacement with a straightforward definition syntax.
It's to Go what thoughtbot/factory_girl is to Ruby.
*/
package gory

import (
	"fmt"
	"reflect"
)

/*
A function executed when defining factories.
Use the factory parameter as a map to store values for
your struct fields.

Example:

	gory.Define("person", Person{}, func(factory gory.Factory) {
		factory["FirstName"] = "Jane"
		factory["LastName"] = "Doe"
		factory["Admin"] = false
	})

Later, when you build the object, the FirstName, LastName,
and Admin fields will be set to the parameters you specified.
If you attempt to set fields that don't exist or are not exported,
the Build() function will panic once called.
*/
type FactoryBuilder func(factory Factory)

var defined definitions

func initDefined() {
	if defined == nil {
		defined = make(definitions, 0)
	}
}

/*
Define a factory. Factories can be retrieved later by using the
definitionName parameter. You cannot define two factories with
identical names.

The instance parameter must be a struct literal of the type you'd
like the factory to return when building.

The builder parameter is a function executed when defining factories.
Use the factory parameter as a map to store values for your struct fields.
If this is nil, no fields on the returned struct are set.
*/
func Define(definitionName string, instance interface{}, builder FactoryBuilder) {
	initDefined()

	factory := make(Factory, 0)
	if builder != nil {
		builder(factory)
	}

	defined.set(definitionName, newDefinition(instance, factory))
}

/*
Returns an instance of the struct defined using the definitionName parameter
and the Define() function. If no matching definition exists, this function
panics. It also panics if an attempt to set an invalid field was made from
within the Define() function.

Example:

	person := gory.Build("person").(*Person)
	fmt.Println(person.FirstName) // "Jane"

*/
func Build(definitionName string) interface{} {
	definition := defined.get(definitionName)
	return build(definition.structType, definition.factory)
}

/*
Returns an instance of the struct defined by definitionName, but using the
specified parameters instead of the defined factory's values. If names of
fields that don't exist or are not exported are specified, the function panics.

Example:

	person := gory.BuildWithParams("person", Person{}, {
		"Name": "John"
	}).(*Person)
	fmt.Println(person.Name) // "John" instead of Factory value

*/
func BuildWithParams(definitionName string, params Factory) interface{} {
	definition := defined.get(definitionName)
	factory := definition.factory.copy()
	factory.merge(params)

	return build(definition.structType, factory)
}

func build(structType reflect.Type, factory Factory) interface{} {
	instance := reflect.New(structType)

	for name, value := range factory {
		field := instance.Elem().FieldByName(name)
		if !field.IsValid() {
			message := fmt.Sprintf("gory: '%s' is not a valid field on %s",
				name, structType.Name())
			panic(message)
		}
		if !field.CanSet() {
			message := fmt.Sprintf("gory: Field '%s' on %s is not an exported struct field; its value cannot be set",
				name, structType.Name())
			panic(message)
		}

		if isLazy(value) {
			value = value.(next)()
		}

		field.Set(reflect.ValueOf(value))
	}

	return instance.Interface()
}

func isLazy(value interface{}) bool {
	return reflect.TypeOf(value) == reflect.TypeOf(Lazy(func() interface{} {
		return nil
	}))
}
