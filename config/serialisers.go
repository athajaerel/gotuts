package main

import "fmt"
import "strconv"

type Serialiser interface {
	ioi32(*int32)
	ios(*string)
}

type Serialisable interface {
	serialise(Serialiser)
}

type Sizer struct {
	size *uint16
}

func (re Sizer) ioi32(i *int32) {
	*(re.size) += 4
}

func (re Sizer) ios(s *string) {
	*(re.size) += uint16(len(*s))
	*(re.size) += 4
}

type Saver struct {
	buffer *[]byte
	ptr uint16
}

func (re Saver) ioi32(i *int32) {
	fmt.Println("Writing: " + strconv.Itoa(int(*i)))
	a := byte(*i >> 24)
	b := byte(*i >> 16)
	c := byte(*i >> 8)
	d := byte(*i >> 0)
	fmt.Println("a: " + strconv.Itoa(int(a)))
	fmt.Println("b: " + strconv.Itoa(int(b)))
	fmt.Println("c: " + strconv.Itoa(int(c)))
	fmt.Println("d: " + strconv.Itoa(int(d)))
	(*(re.buffer))[re.ptr + 0] = byte(a)
	(*(re.buffer))[re.ptr + 1] = byte(b)
	(*(re.buffer))[re.ptr + 2] = byte(c)
	(*(re.buffer))[re.ptr + 3] = byte(d)
	re.ptr += 4
	fmt.Println("New index: " + strconv.Itoa(int(re.ptr))) // hmm, writes a copy only??
}

func (re Saver) ios(s *string) {
	var val int32 = int32(len(*s))
	re.ioi32(&val)
	//
	re.ptr += 4
	re.ptr += uint16(val)
}

type Loader struct {
	buffer *[]byte
	ptr uint16
}

func (re Loader) ioi32(i *int32) {
	var val int32 = 0
	val += int32((*(re.buffer))[re.ptr + 0] << 24)
	val += int32((*(re.buffer))[re.ptr + 1] << 16)
	val += int32((*(re.buffer))[re.ptr + 2] << 8)
	val += int32((*(re.buffer))[re.ptr + 3] << 0)
	*i = val
	re.ptr += 4
}

func (re Loader) ios(s *string) {
	var val int32 = 0
	re.ioi32(&val)
	//
}

type Printer struct {
}

func (re Printer) ioi32(i *int32) {
	fmt.Println(strconv.Itoa(int(*i)))
}

func (re Printer) ios(s *string) {
	fmt.Println(*s)
}
