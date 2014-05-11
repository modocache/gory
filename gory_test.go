package gory_test

import (
	"github.com/modocache/gory"
	"unsafe"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gory", func() {
	Describe("Define()", func() {
		Context("when the factory function is nil", func() {
			It("creates a definition that doesn't set any values", func() {
				gory.Define("empty", Builtin{}, nil)
				empty := gory.Build("empty").(*Builtin)
				Expect(empty.Int).To(Equal(0))
			})
		})

		Context("when attempting to create a definition with the same name", func() {
			var name string
			BeforeEach(func() {
				name = "dupe"
				gory.Define(name, Builtin{}, func(factory gory.Factory) {})
			})

			It("panics", func() {
				define := func() { gory.Define(name, Builtin{}, func(factory gory.Factory) {}) }
				Expect(define).To(Panic())
			})
		})
	})

	Describe("Build()", func() {
		Context("when there is no corresponding definition", func() {
			It("panics", func() {
				Expect(func() { gory.Build("undefined") }).To(Panic())
			})
		})

		Context("when the definition defines an unexported value", func() {
			It("panics", func() {
				Expect(func() { gory.Build("unexported") }).To(Panic())
			})
		})

		Context("when the definition defines values for builtin types", func() {
			var builtin *Builtin
			BeforeEach(func() {
				builtin = gory.Build("builtin").(*Builtin)
			})

			It("sets the specified bool value", func() {
				Expect(builtin.Bool).To(BeTrue())
			})
			It("sets the specified string value", func() {
				Expect(builtin.String).To(Equal("string"))
			})
			It("sets the specified int value", func() {
				Expect(builtin.Int).To(Equal(1))
			})
			It("sets the specified int8 value", func() {
				Expect(builtin.Int8).To(Equal(int8(8)))
			})
			It("sets the specified int16 value", func() {
				Expect(builtin.Int16).To(Equal(int16(16)))
			})
			It("sets the specified int32 value", func() {
				Expect(builtin.Int32).To(Equal(int32(32)))
			})
			It("sets the specified int64 value", func() {
				Expect(builtin.Int64).To(Equal(int64(64)))
			})
			It("sets the specified uint value", func() {
				Expect(builtin.Uint).To(Equal(uint(1)))
			})
			It("sets the specified uint8 value", func() {
				Expect(builtin.Uint8).To(Equal(uint8(8)))
			})
			It("sets the specified uint16 value", func() {
				Expect(builtin.Uint16).To(Equal(uint16(16)))
			})
			It("sets the specified uint32 value", func() {
				Expect(builtin.Uint32).To(Equal(uint32(32)))
			})
			It("sets the specified uint64 value", func() {
				Expect(builtin.Uint64).To(Equal(uint64(64)))
			})
			It("sets the specified uintptr value", func() {
				Expect(builtin.Uintptr).To(Equal(unsafe.Sizeof([]int{})))
			})
			It("sets the specified byte value", func() {
				Expect(builtin.Byte).To(Equal(builtin.String[0]))
			})
			It("sets the specified rune value", func() {
				Expect(builtin.Rune).To(Equal('⌘'))
			})
			It("sets the specified float32 value", func() {
				Expect(builtin.Float32).To(Equal(float32(32.0)))
			})
			It("sets the specified float64 value", func() {
				Expect(builtin.Float64).To(Equal(float64(64.0)))
			})
			It("sets the specified complex64 value", func() {
				Expect(builtin.Complex64).To(Equal(complex(float32(10.0), float32(-1.0))))
			})
			It("sets the specified complex128 value", func() {
				Expect(builtin.Complex128).To(Equal(complex(float64(-10.0), float64(1.0))))
			})
		})

		Context("when the definition defines values for slices", func() {
			It("sets the slice", func() {
				slice := gory.Build("slice").(*Slice)
				Expect(slice.Ints).To(Equal([]int{1, 2, 3}))
			})
		})

		Context("when the definition defines values for maps", func() {
			It("sets the map", func() {
				m := gory.Build("map").(*Map)
				Expect(m.Strings["key"]).To(Equal("value"))
			})
		})

		Context("when the definition defines values for struct pointers", func() {
			It("sets the struct pointers", func() {
				strct := gory.Build("struct").(*Struct)
				Expect(strct.Bin.Int).To(Equal(1))
			})
		})

		Context("when the definition defines values for embedded fields", func() {
			It("sets those values", func() {
				embedded := gory.Build("embedded").(*Embedded)
				Expect(embedded.Int).To(Equal(1))
			})
		})

		Describe("lazy values", func() {
			var name string
			BeforeEach(func() { name = "lazily" })
			It("evaluates the values when they're built", func() {
				left := gory.Build(name).(*Lazily)
				right := gory.Build(name).(*Lazily)
				Expect(left.Uuid).ToNot((Equal(right.Uuid)))
			})
		})

		Describe("sequenced values", func() {
			var name string
			BeforeEach(func() { name = "sequenced" })
			It("sequences the values", func() {
				first := gory.Build(name).(*Builtin)
				second := gory.Build(name).(*Builtin)

				Expect(first.Int).To(Equal(0))
				Expect(first.String).To(Equal("string 0"))
				Expect(second.Int).To(Equal(1))
				Expect(second.String).To(Equal("string 1"))
			})
		})
	})

	Describe("BuildWithParams()", func() {
		Context("when given an invalid field name", func() {
			It("panics", func() {
				build := func() {
					gory.BuildWithParams("builtin", gory.Factory{"Name": "Bob"})
				}
				Expect(build).To(Panic())
			})
		})

		Context("when given values valid field names", func() {
			It("uses those values instead of the Factory values", func() {
				builtin := gory.BuildWithParams("builtin", gory.Factory{
					"Int":    21,
					"String": "hello",
				}).(*Builtin)
				Expect(builtin.Int).To(Equal(21))
				Expect(builtin.String).To(Equal("hello"))
				Expect(builtin.Int8).To(Equal(int8(8)))
			})
		})
	})
})
