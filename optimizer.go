package gooptimizer

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"strings"
)

type Alignment struct {
	stdout  bool
	expect  uint8
	store   [256]uint8
	current uint8
	align   uint8
	size    uint8
	i18n    *i18n
	fields  []reflect.StructField
}

func (a *Alignment) Align() {
	for _, v := range a.fields {
		t := v.Type
		align := uint8(t.Align())
		size := uint8(t.Size())
		if a.stdout {
			fmt.Printf("%s: %s, %s: %s, %s: %d, %s: %d\n", a.i18n.Get("Field"), v.Name, a.i18n.Get("Type"), t.Name(), a.i18n.Get("Align"), align, a.i18n.Get("Size"), size)
		}

		a.expect += size
		for i := 0; i < int(math.Ceil(float64(size)/float64(align))); i++ {
			if a.align-a.store[a.current] < align {
				a.current++
			}
			a.store[a.current] += align
		}
	}
}

func (a *Alignment) sort() {
	sort.SliceStable(a.fields, func(i, j int) bool {
		if a.fields[i].Type.Align() == a.fields[j].Type.Align() {
			return a.fields[i].Type.Size() < a.fields[j].Type.Size()
		}
		return a.fields[i].Type.Align() < a.fields[j].Type.Align()
	})
}
func (a *Alignment) reset() {
	for i := 0; i < 256; i++ {
		a.store[i] = 0
	}
	a.expect = 0
	a.current = 0
}

func (a *Alignment) Optimize() (ok bool, err error) {
	if a.stdout {
		fmt.Printf("%s:\n", a.i18n.Get("Field alignment arrangement before"))
	}

	a.reset()
	a.Align()

	actual := uint8(0)
	for _, i := range a.store {
		if i == 0 {
			break
		}
		if a.stdout {
			fmt.Print(strings.Repeat("■", int(i)))
			fmt.Print(strings.Repeat("□", int(a.align-i)))
		}
		actual += a.align
	}
	if actual != a.size {
		fmt.Println("error")
		return
	}
	if a.stdout {
		fmt.Printf("\n%s: %d, %s: %d, %s: %d\n", a.i18n.Get("Align"), a.align, a.i18n.Get("Expect Size"), a.expect, a.i18n.Get("Actual Size"), actual)
	}

	oldActual := actual

	if a.stdout {
		fmt.Printf("\n%s:", a.i18n.Get("Field alignment arrangement after"))
	}

	a.reset()
	a.sort()
	a.Align()

	actual = uint8(0)
	for _, i := range a.store {
		if i == 0 {
			break
		}
		if a.stdout {
			fmt.Print(strings.Repeat("■", int(i)))
			fmt.Print(strings.Repeat("□", int(a.align-i)))
		}
		actual += a.align
	}
	if a.stdout {
		fmt.Printf("\n%s: %d, %s: %d, %s: %d\n", a.i18n.Get("Align"), a.align, a.i18n.Get("Expect Size"), a.expect, a.i18n.Get("Actual Size"), actual)
	}

	save := oldActual - actual
	ok = save == 0
	if !ok && a.stdout {
		fmt.Printf("%s: %d\n", a.i18n.Get("You should optimize the structure; there's potential to save"), save)
	}

	return
}

func StructAlign(v any) (ok bool) {
	if v == nil {
		fmt.Println("v nil")
		return
	}

	a := new(Alignment)
	a.i18n = DefaultI18n

	rv := reflect.TypeOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	a.align = uint8(rv.Align())
	a.size = uint8(rv.Size())

	l := rv.NumField()
	for i := 0; i < l; i++ {
		a.fields = append(a.fields, rv.Field(i))
	}

	ok, _ = a.Optimize()
	return
}

func StructAlignWithPrint(v any) {
	if v == nil {
		fmt.Println("v nil")
		return
	}

	a := new(Alignment)
	a.i18n = DefaultI18n
	a.stdout = true

	rv := reflect.TypeOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	a.align = uint8(rv.Align())
	a.size = uint8(rv.Size())

	l := rv.NumField()
	for i := 0; i < l; i++ {
		a.fields = append(a.fields, rv.Field(i))
	}

	_, _ = a.Optimize()
}

func StructAlignWithCNPrint(v any) {
	if v == nil {
		fmt.Println("v nil")
		return
	}

	a := new(Alignment)
	a.i18n = DefaultI18n
	a.i18n.lang = CN
	a.stdout = true

	rv := reflect.TypeOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	a.align = uint8(rv.Align())
	a.size = uint8(rv.Size())

	l := rv.NumField()
	for i := 0; i < l; i++ {
		a.fields = append(a.fields, rv.Field(i))
	}

	_, _ = a.Optimize()
}
