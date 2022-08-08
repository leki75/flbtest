package karmem

import (
	karmem "karmem.org/golang"
	"unsafe"
)

var _ unsafe.Pointer

type (
	PacketIdentifier uint64
)

const (
	PacketIdentifierQuote      = 8596760148279994604
	PacketIdentifierTrade      = 6208342187296333386
	PacketIdentifierQuoteBatch = 8603685289548553011
	PacketIdentifierTradeBatch = 14427887524916167865
)

type Quote struct {
	Timestamp   uint64
	ReceivedAt  uint64
	BidPrice    float64
	AskPrice    float64
	BidSize     uint32
	AskSize     uint32
	Symbol      [11]byte
	Conditions  [2]byte
	BidExchange byte
	AskExchange byte
	Tape        byte
	Nbbo        bool
}

func (x *Quote) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierQuote
}

func (x *Quote) Reset() {
	x.Timestamp = 0
	x.ReceivedAt = 0
	x.BidPrice = 0
	x.AskPrice = 0
	x.BidSize = 0
	x.AskSize = 0
	x.Symbol = [11]byte{}
	x.Conditions = [2]byte{}
	x.BidExchange = 0
	x.AskExchange = 0
	x.Tape = 0
	x.Nbbo = false
}

func (x *Quote) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *Quote) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(64)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	__TimestampOffset := offset + 0
	writer.Write8At(__TimestampOffset, *(*uint64)(unsafe.Pointer(&x.Timestamp)))
	__ReceivedAtOffset := offset + 8
	writer.Write8At(__ReceivedAtOffset, *(*uint64)(unsafe.Pointer(&x.ReceivedAt)))
	__BidPriceOffset := offset + 16
	writer.Write8At(__BidPriceOffset, *(*uint64)(unsafe.Pointer(&x.BidPrice)))
	__AskPriceOffset := offset + 24
	writer.Write8At(__AskPriceOffset, *(*uint64)(unsafe.Pointer(&x.AskPrice)))
	__BidSizeOffset := offset + 32
	writer.Write4At(__BidSizeOffset, *(*uint32)(unsafe.Pointer(&x.BidSize)))
	__AskSizeOffset := offset + 36
	writer.Write4At(__AskSizeOffset, *(*uint32)(unsafe.Pointer(&x.AskSize)))
	__SymbolOffset := offset + 40
	writer.WriteAt(__SymbolOffset, (*[11]byte)(unsafe.Pointer(&x.Symbol))[:])
	__ConditionsOffset := offset + 51
	writer.WriteAt(__ConditionsOffset, (*[2]byte)(unsafe.Pointer(&x.Conditions))[:])
	__BidExchangeOffset := offset + 53
	writer.Write1At(__BidExchangeOffset, *(*uint8)(unsafe.Pointer(&x.BidExchange)))
	__AskExchangeOffset := offset + 54
	writer.Write1At(__AskExchangeOffset, *(*uint8)(unsafe.Pointer(&x.AskExchange)))
	__TapeOffset := offset + 55
	writer.Write1At(__TapeOffset, *(*uint8)(unsafe.Pointer(&x.Tape)))
	__NbboOffset := offset + 56
	writer.Write1At(__NbboOffset, *(*uint8)(unsafe.Pointer(&x.Nbbo)))

	return offset, nil
}

func (x *Quote) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewQuoteViewer(reader, 0), reader)
}

func (x *Quote) Read(viewer *QuoteViewer, reader *karmem.Reader) {
	x.Timestamp = viewer.Timestamp()
	x.ReceivedAt = viewer.ReceivedAt()
	x.BidPrice = viewer.BidPrice()
	x.AskPrice = viewer.AskPrice()
	x.BidSize = viewer.BidSize()
	x.AskSize = viewer.AskSize()
	__SymbolSlice := viewer.Symbol()
	__SymbolLen := len(__SymbolSlice)
	copy(x.Symbol[:], __SymbolSlice)
	for i := __SymbolLen; i < len(x.Symbol); i++ {
		x.Symbol[i] = 0
	}
	__ConditionsSlice := viewer.Conditions()
	__ConditionsLen := len(__ConditionsSlice)
	copy(x.Conditions[:], __ConditionsSlice)
	for i := __ConditionsLen; i < len(x.Conditions); i++ {
		x.Conditions[i] = 0
	}
	x.BidExchange = viewer.BidExchange()
	x.AskExchange = viewer.AskExchange()
	x.Tape = viewer.Tape()
	x.Nbbo = viewer.Nbbo()
}

type Trade struct {
	ID         uint64
	Timestamp  uint64
	ReceivedAt uint64
	Price      float64
	Volume     uint32
	Conditions [4]byte
	Symbol     [11]byte
	Exchange   byte
	Tape       byte
}

func (x *Trade) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierTrade
}

func (x *Trade) Reset() {
	x.ID = 0
	x.Timestamp = 0
	x.ReceivedAt = 0
	x.Price = 0
	x.Volume = 0
	x.Conditions = [4]byte{}
	x.Symbol = [11]byte{}
	x.Exchange = 0
	x.Tape = 0
}

func (x *Trade) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *Trade) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(56)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	__IDOffset := offset + 0
	writer.Write8At(__IDOffset, *(*uint64)(unsafe.Pointer(&x.ID)))
	__TimestampOffset := offset + 8
	writer.Write8At(__TimestampOffset, *(*uint64)(unsafe.Pointer(&x.Timestamp)))
	__ReceivedAtOffset := offset + 16
	writer.Write8At(__ReceivedAtOffset, *(*uint64)(unsafe.Pointer(&x.ReceivedAt)))
	__PriceOffset := offset + 24
	writer.Write8At(__PriceOffset, *(*uint64)(unsafe.Pointer(&x.Price)))
	__VolumeOffset := offset + 32
	writer.Write4At(__VolumeOffset, *(*uint32)(unsafe.Pointer(&x.Volume)))
	__ConditionsOffset := offset + 36
	writer.WriteAt(__ConditionsOffset, (*[4]byte)(unsafe.Pointer(&x.Conditions))[:])
	__SymbolOffset := offset + 40
	writer.WriteAt(__SymbolOffset, (*[11]byte)(unsafe.Pointer(&x.Symbol))[:])
	__ExchangeOffset := offset + 51
	writer.Write1At(__ExchangeOffset, *(*uint8)(unsafe.Pointer(&x.Exchange)))
	__TapeOffset := offset + 52
	writer.Write1At(__TapeOffset, *(*uint8)(unsafe.Pointer(&x.Tape)))

	return offset, nil
}

func (x *Trade) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewTradeViewer(reader, 0), reader)
}

func (x *Trade) Read(viewer *TradeViewer, reader *karmem.Reader) {
	x.ID = viewer.ID()
	x.Timestamp = viewer.Timestamp()
	x.ReceivedAt = viewer.ReceivedAt()
	x.Price = viewer.Price()
	x.Volume = viewer.Volume()
	__ConditionsSlice := viewer.Conditions()
	__ConditionsLen := len(__ConditionsSlice)
	copy(x.Conditions[:], __ConditionsSlice)
	for i := __ConditionsLen; i < len(x.Conditions); i++ {
		x.Conditions[i] = 0
	}
	__SymbolSlice := viewer.Symbol()
	__SymbolLen := len(__SymbolSlice)
	copy(x.Symbol[:], __SymbolSlice)
	for i := __SymbolLen; i < len(x.Symbol); i++ {
		x.Symbol[i] = 0
	}
	x.Exchange = viewer.Exchange()
	x.Tape = viewer.Tape()
}

type QuoteBatch struct {
	Quotes []Quote
}

func (x *QuoteBatch) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierQuoteBatch
}

func (x *QuoteBatch) Reset() {
	for i := range x.Quotes {
		x.Quotes[i].Reset()
	}
	x.Quotes = x.Quotes[:0]
}

func (x *QuoteBatch) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *QuoteBatch) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(24)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(16))
	__QuotesSize := uint(64 * len(x.Quotes))
	__QuotesOffset, err := writer.Alloc(__QuotesSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+4, uint32(__QuotesOffset))
	writer.Write4At(offset+4+4, uint32(__QuotesSize))
	writer.Write4At(offset+4+4+4, 64)
	for i := range x.Quotes {
		if _, err := x.Quotes[i].Write(writer, __QuotesOffset); err != nil {
			return offset, err
		}
		__QuotesOffset += 64
	}

	return offset, nil
}

func (x *QuoteBatch) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewQuoteBatchViewer(reader, 0), reader)
}

func (x *QuoteBatch) Read(viewer *QuoteBatchViewer, reader *karmem.Reader) {
	__QuotesSlice := viewer.Quotes(reader)
	__QuotesLen := len(__QuotesSlice)
	if __QuotesLen > cap(x.Quotes) {
		x.Quotes = append(x.Quotes, make([]Quote, __QuotesLen-len(x.Quotes))...)
	} else if __QuotesLen > len(x.Quotes) {
		x.Quotes = x.Quotes[:__QuotesLen]
	}
	for i := range x.Quotes {
		if i >= __QuotesLen {
			x.Quotes[i].Reset()
		} else {
			x.Quotes[i].Read(&__QuotesSlice[i], reader)
		}
	}
	x.Quotes = x.Quotes[:__QuotesLen]
}

type TradeBatch struct {
	Trades []Trade
}

func (x *TradeBatch) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierTradeBatch
}

func (x *TradeBatch) Reset() {
	for i := range x.Trades {
		x.Trades[i].Reset()
	}
	x.Trades = x.Trades[:0]
}

func (x *TradeBatch) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *TradeBatch) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(24)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(16))
	__TradesSize := uint(56 * len(x.Trades))
	__TradesOffset, err := writer.Alloc(__TradesSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+4, uint32(__TradesOffset))
	writer.Write4At(offset+4+4, uint32(__TradesSize))
	writer.Write4At(offset+4+4+4, 56)
	for i := range x.Trades {
		if _, err := x.Trades[i].Write(writer, __TradesOffset); err != nil {
			return offset, err
		}
		__TradesOffset += 56
	}

	return offset, nil
}

func (x *TradeBatch) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewTradeBatchViewer(reader, 0), reader)
}

func (x *TradeBatch) Read(viewer *TradeBatchViewer, reader *karmem.Reader) {
	__TradesSlice := viewer.Trades(reader)
	__TradesLen := len(__TradesSlice)
	if __TradesLen > cap(x.Trades) {
		x.Trades = append(x.Trades, make([]Trade, __TradesLen-len(x.Trades))...)
	} else if __TradesLen > len(x.Trades) {
		x.Trades = x.Trades[:__TradesLen]
	}
	for i := range x.Trades {
		if i >= __TradesLen {
			x.Trades[i].Reset()
		} else {
			x.Trades[i].Read(&__TradesSlice[i], reader)
		}
	}
	x.Trades = x.Trades[:__TradesLen]
}

type QuoteViewer struct {
	_data [64]byte
}

var _NullQuoteViewer = QuoteViewer{}

func NewQuoteViewer(reader *karmem.Reader, offset uint32) (v *QuoteViewer) {
	if !reader.IsValidOffset(offset, 64) {
		return &_NullQuoteViewer
	}
	v = (*QuoteViewer)(unsafe.Add(reader.Pointer, offset))
	return v
}

func (x *QuoteViewer) size() uint32 {
	return 64
}
func (x *QuoteViewer) Timestamp() (v uint64) {
	return *(*uint64)(unsafe.Add(unsafe.Pointer(&x._data), 0))
}
func (x *QuoteViewer) ReceivedAt() (v uint64) {
	return *(*uint64)(unsafe.Add(unsafe.Pointer(&x._data), 8))
}
func (x *QuoteViewer) BidPrice() (v float64) {
	return *(*float64)(unsafe.Add(unsafe.Pointer(&x._data), 16))
}
func (x *QuoteViewer) AskPrice() (v float64) {
	return *(*float64)(unsafe.Add(unsafe.Pointer(&x._data), 24))
}
func (x *QuoteViewer) BidSize() (v uint32) {
	return *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 32))
}
func (x *QuoteViewer) AskSize() (v uint32) {
	return *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 36))
}
func (x *QuoteViewer) Symbol() (v []byte) {
	slice := [3]uintptr{
		uintptr(unsafe.Add(unsafe.Pointer(&x._data), 40)), 11, 11,
	}
	return *(*[]byte)(unsafe.Pointer(&slice))
}
func (x *QuoteViewer) Conditions() (v []byte) {
	slice := [3]uintptr{
		uintptr(unsafe.Add(unsafe.Pointer(&x._data), 51)), 2, 2,
	}
	return *(*[]byte)(unsafe.Pointer(&slice))
}
func (x *QuoteViewer) BidExchange() (v byte) {
	return *(*byte)(unsafe.Add(unsafe.Pointer(&x._data), 53))
}
func (x *QuoteViewer) AskExchange() (v byte) {
	return *(*byte)(unsafe.Add(unsafe.Pointer(&x._data), 54))
}
func (x *QuoteViewer) Tape() (v byte) {
	return *(*byte)(unsafe.Add(unsafe.Pointer(&x._data), 55))
}
func (x *QuoteViewer) Nbbo() (v bool) {
	return *(*bool)(unsafe.Add(unsafe.Pointer(&x._data), 56))
}

type TradeViewer struct {
	_data [56]byte
}

var _NullTradeViewer = TradeViewer{}

func NewTradeViewer(reader *karmem.Reader, offset uint32) (v *TradeViewer) {
	if !reader.IsValidOffset(offset, 56) {
		return &_NullTradeViewer
	}
	v = (*TradeViewer)(unsafe.Add(reader.Pointer, offset))
	return v
}

func (x *TradeViewer) size() uint32 {
	return 56
}
func (x *TradeViewer) ID() (v uint64) {
	return *(*uint64)(unsafe.Add(unsafe.Pointer(&x._data), 0))
}
func (x *TradeViewer) Timestamp() (v uint64) {
	return *(*uint64)(unsafe.Add(unsafe.Pointer(&x._data), 8))
}
func (x *TradeViewer) ReceivedAt() (v uint64) {
	return *(*uint64)(unsafe.Add(unsafe.Pointer(&x._data), 16))
}
func (x *TradeViewer) Price() (v float64) {
	return *(*float64)(unsafe.Add(unsafe.Pointer(&x._data), 24))
}
func (x *TradeViewer) Volume() (v uint32) {
	return *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 32))
}
func (x *TradeViewer) Conditions() (v []byte) {
	slice := [3]uintptr{
		uintptr(unsafe.Add(unsafe.Pointer(&x._data), 36)), 4, 4,
	}
	return *(*[]byte)(unsafe.Pointer(&slice))
}
func (x *TradeViewer) Symbol() (v []byte) {
	slice := [3]uintptr{
		uintptr(unsafe.Add(unsafe.Pointer(&x._data), 40)), 11, 11,
	}
	return *(*[]byte)(unsafe.Pointer(&slice))
}
func (x *TradeViewer) Exchange() (v byte) {
	return *(*byte)(unsafe.Add(unsafe.Pointer(&x._data), 51))
}
func (x *TradeViewer) Tape() (v byte) {
	return *(*byte)(unsafe.Add(unsafe.Pointer(&x._data), 52))
}

type QuoteBatchViewer struct {
	_data [24]byte
}

var _NullQuoteBatchViewer = QuoteBatchViewer{}

func NewQuoteBatchViewer(reader *karmem.Reader, offset uint32) (v *QuoteBatchViewer) {
	if !reader.IsValidOffset(offset, 8) {
		return &_NullQuoteBatchViewer
	}
	v = (*QuoteBatchViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return &_NullQuoteBatchViewer
	}
	return v
}

func (x *QuoteBatchViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(&x._data))
}
func (x *QuoteBatchViewer) Quotes(reader *karmem.Reader) (v []QuoteViewer) {
	if 4+12 > x.size() {
		return []QuoteViewer{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4+4))
	if !reader.IsValidOffset(offset, size) {
		return []QuoteViewer{}
	}
	length := uintptr(size / 64)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]QuoteViewer)(unsafe.Pointer(&slice))
}

type TradeBatchViewer struct {
	_data [24]byte
}

var _NullTradeBatchViewer = TradeBatchViewer{}

func NewTradeBatchViewer(reader *karmem.Reader, offset uint32) (v *TradeBatchViewer) {
	if !reader.IsValidOffset(offset, 8) {
		return &_NullTradeBatchViewer
	}
	v = (*TradeBatchViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return &_NullTradeBatchViewer
	}
	return v
}

func (x *TradeBatchViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(&x._data))
}
func (x *TradeBatchViewer) Trades(reader *karmem.Reader) (v []TradeViewer) {
	if 4+12 > x.size() {
		return []TradeViewer{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4+4))
	if !reader.IsValidOffset(offset, size) {
		return []TradeViewer{}
	}
	length := uintptr(size / 56)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]TradeViewer)(unsafe.Pointer(&slice))
}
