package main

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/heartchord/jxonline/gamestruct"
)

type A struct {
	// should be exported member when read back from buffer
	There byte
	One   int32
	Two   int32
}

func main() {
	var aa A
	var a A

	a.One = int32(1)
	a.Two = int32(2)
	a.There = byte(1)

	buf := new(bytes.Buffer)
	fmt.Println("a‘s size is ", binary.Size(a))
	binary.Write(buf, binary.LittleEndian, a)
	fmt.Println("after write ，buf is:", buf.Bytes())

	binary.Read(buf, binary.LittleEndian, &aa)
	fmt.Println("after aa is ", aa)

	var r gamestruct.RoleData
	fmt.Println(binary.Size(r))
}
