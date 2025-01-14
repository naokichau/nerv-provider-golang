// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package schema

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type RuntimeCallParams struct {
	_tab flatbuffers.Table
}

func GetRootAsRuntimeCallParams(buf []byte, offset flatbuffers.UOffsetT) *RuntimeCallParams {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &RuntimeCallParams{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *RuntimeCallParams) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *RuntimeCallParams) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *RuntimeCallParams) Input(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *RuntimeCallParams) InputLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *RuntimeCallParams) InputBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *RuntimeCallParams) MutateInput(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *RuntimeCallParams) Timeout() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RuntimeCallParams) MutateTimeout(n uint64) bool {
	return rcv._tab.MutateUint64Slot(6, n)
}

func RuntimeCallParamsStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func RuntimeCallParamsAddInput(builder *flatbuffers.Builder, Input flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(Input), 0)
}
func RuntimeCallParamsStartInputVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func RuntimeCallParamsAddTimeout(builder *flatbuffers.Builder, Timeout uint64) {
	builder.PrependUint64Slot(1, Timeout, 0)
}
func RuntimeCallParamsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
