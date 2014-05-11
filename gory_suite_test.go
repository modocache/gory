package gory_test

import (
	"fmt"
	"github.com/modocache/gory"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os/exec"
	"unsafe"

	"testing"
)

type Builtin struct {
	Bool       bool
	String     string
	Int        int
	Int8       int8
	Int16      int16
	Int32      int32
	Int64      int64
	Uint       uint
	Uint8      uint8
	Uint16     uint16
	Uint32     uint32
	Uint64     uint64
	Uintptr    uintptr
	Byte       byte
	Rune       rune
	Float32    float32
	Float64    float64
	Complex64  complex64
	Complex128 complex128
}

type Unexported struct {
	str string
}

type Slice struct {
	Ints []int
}

type Map struct {
	Strings map[string]string
}

type Struct struct {
	Bin *Builtin
}

type Embedded struct {
	Builtin
}

type Lazily struct {
	Uuid string
}

var _ = BeforeSuite(func() {
	gory.Define("builtin", Builtin{}, func(factory gory.Factory) {
		factory["Bool"] = true
		factory["String"] = "string"
		factory["Int"] = 1
		factory["Int8"] = int8(8)
		factory["Int16"] = int16(16)
		factory["Int32"] = int32(32)
		factory["Int64"] = int64(64)
		factory["Uint"] = uint(1)
		factory["Uint8"] = uint8(8)
		factory["Uint16"] = uint16(16)
		factory["Uint32"] = uint32(32)
		factory["Uint64"] = uint64(64)
		factory["Uintptr"] = unsafe.Sizeof([]int{})
		factory["Byte"] = factory["String"].(string)[0]
		factory["Rune"] = '⌘'
		factory["Float32"] = float32(32.0)
		factory["Float64"] = float64(64.0)
		factory["Complex64"] = complex(float32(10.0), float32(-1.0))
		factory["Complex128"] = complex(float64(-10.0), float64(1.0))
	})

	gory.Define("unexported", Unexported{}, func(factory gory.Factory) {
		factory["str"] = "boom!"
	})

	gory.Define("slice", Slice{}, func(factory gory.Factory) {
		factory["Ints"] = []int{1, 2, 3}
	})

	gory.Define("map", Map{}, func(factory gory.Factory) {
		strings := make(map[string]string, 0)
		strings["key"] = "value"
		factory["Strings"] = strings
	})

	gory.Define("struct", Struct{}, func(factory gory.Factory) {
		factory["Bin"] = &Builtin{Int: 1}
	})

	gory.Define("embedded", Embedded{}, func(factory gory.Factory) {
		factory["Int"] = 1
	})

	gory.Define("sequenced", Builtin{}, func(factory gory.Factory) {
		factory["Int"] = gory.Sequence(gory.IntSequencer)
		factory["String"] = gory.Sequence(func(n int) interface{} {
			return fmt.Sprintf("string %d", n)
		})
	})

	gory.Define("lazily", Lazily{}, func(factory gory.Factory) {
		factory["Uuid"] = gory.Lazy(func() interface{} {
			out, _ := exec.Command("uuidgen").Output()
			return string(out)
		})
	})
})

func TestGory(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gory Suite")
}
