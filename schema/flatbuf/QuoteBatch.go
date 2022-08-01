// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flatbuf

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type QuoteBatch struct {
	_tab flatbuffers.Table
}

func GetRootAsQuoteBatch(buf []byte, offset flatbuffers.UOffsetT) *QuoteBatch {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &QuoteBatch{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsQuoteBatch(buf []byte, offset flatbuffers.UOffsetT) *QuoteBatch {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &QuoteBatch{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *QuoteBatch) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *QuoteBatch) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *QuoteBatch) Quotes(obj *Quote, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *QuoteBatch) QuotesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func QuoteBatchStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func QuoteBatchAddQuotes(builder *flatbuffers.Builder, quotes flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(quotes), 0)
}
func QuoteBatchStartQuotesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func QuoteBatchEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}