package main

import (
	"fmt"
	"github.com/lizongying/gooptimizer"
)

type T1 struct {
	u8   uint8
	u32  uint32
	u64  uint64
	u16  uint16
	i32  int32
	iu64 int64
	i8   int8
	i16  int16
	uptr uintptr
	s    string
}

func main() {
	should := gooptimizer.StructAlign(new(gooptimizer.Alignment))
	fmt.Println(should)

	// print
	//gooptimizer.StructAlignWithPrint(new(gooptimizer.Alignment))

	// cn print
	//gooptimizer.StructAlignWithCNPrint(new(T1))
}
