package gooptimizer

import (
	"fmt"
	"math"
	"reflect"
	"sort"
)

type Alignment struct {
	expect  uint8
	store   [256]uint8
	current uint8
	align   uint8
	size    uint8
	fields  []reflect.StructField
}

func (a *Alignment) Align() {
	for _, v := range a.fields {
		t := v.Type
		align := uint8(t.Align())
		size := uint8(t.Size())
		fmt.Printf("字段: %s, 类型: %s, 对齐: %d, 大小: %d\n", v.Name, t.Name(), align, size)

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

func (a *Alignment) Print() {
	fmt.Println("字段对齐排列前:")
	a.reset()
	a.Align()

	actual := uint8(0)
	for _, i := range a.store {
		if i == 0 {
			break
		}
		//fmt.Println(i)
		actual += uint8(a.align)
	}
	if actual != a.size {
		fmt.Println("error")
		return
	}
	fmt.Printf("对齐: %d, 期望大小: %d, 实际大小: %d\n", a.align, a.expect, actual)

	fmt.Println("\n字段对齐顺序排列后:")
	a.reset()
	a.sort()
	a.Align()

	actual = uint8(0)
	for _, i := range a.store {
		if i == 0 {
			break
		}
		//fmt.Println(i)
		actual += uint8(a.align)
	}
	fmt.Printf("对齐: %d, 期望大小: %d, 实际大小: %d\n", a.align, a.expect, actual)
}

func StructAlignment(v any) {
	if v == nil {
		fmt.Println("v nil")
		return
	}

	a := new(Alignment)

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

	a.Print()
}
