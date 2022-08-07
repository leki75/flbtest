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
	PacketIdentifierRawQuote   = 3087845087638307936
	PacketIdentifierRawTrade   = 1533053534069801195
	PacketIdentifierQuoteBatch = 8603685289548553011
	PacketIdentifierTradeBatch = 14427887524916167865
)

type Quote struct {
	Symbol      string
	Conditions  []byte
	Timestamp   uint64
	ReceivedAt  uint64
	BidPrice    float64
	AskPrice    float64
	BidSize     uint32
	AskSize     uint32
	BidExchange byte
	AskExchange byte
	Tape        byte
	Nbbo        bool
}

func (x *Quote) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierQuote
}

func (x *Quote) Reset() {
	x.Symbol = x.Symbol[:0]
	x.Conditions = x.Conditions[:0]
	x.Timestamp = 0
	x.ReceivedAt = 0
	x.BidPrice = 0
	x.AskPrice = 0
	x.BidSize = 0
	x.AskSize = 0
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
	size := uint(80)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(72))
	__SymbolSize := uint(1 * len(x.Symbol))
	__SymbolOffset, err := writer.Alloc(__SymbolSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+4, uint32(__SymbolOffset))
	writer.Write4At(offset+4+4, uint32(__SymbolSize))
	writer.Write4At(offset+4+4+4, 1)
	__SymbolSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Symbol)), __SymbolSize, __SymbolSize}
	writer.WriteAt(__SymbolOffset, *(*[]byte)(unsafe.Pointer(&__SymbolSlice)))
	__ConditionsSize := uint(1 * len(x.Conditions))
	__ConditionsOffset, err := writer.Alloc(__ConditionsSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+16, uint32(__ConditionsOffset))
	writer.Write4At(offset+16+4, uint32(__ConditionsSize))
	writer.Write4At(offset+16+4+4, 1)
	__ConditionsSlice := *(*[3]uint)(unsafe.Pointer(&x.Conditions))
	__ConditionsSlice[1] = __ConditionsSize
	__ConditionsSlice[2] = __ConditionsSize
	writer.WriteAt(__ConditionsOffset, *(*[]byte)(unsafe.Pointer(&__ConditionsSlice)))
	__TimestampOffset := offset + 28
	writer.Write8At(__TimestampOffset, *(*uint64)(unsafe.Pointer(&x.Timestamp)))
	__ReceivedAtOffset := offset + 36
	writer.Write8At(__ReceivedAtOffset, *(*uint64)(unsafe.Pointer(&x.ReceivedAt)))
	__BidPriceOffset := offset + 44
	writer.Write8At(__BidPriceOffset, *(*uint64)(unsafe.Pointer(&x.BidPrice)))
	__AskPriceOffset := offset + 52
	writer.Write8At(__AskPriceOffset, *(*uint64)(unsafe.Pointer(&x.AskPrice)))
	__BidSizeOffset := offset + 60
	writer.Write4At(__BidSizeOffset, *(*uint32)(unsafe.Pointer(&x.BidSize)))
	__AskSizeOffset := offset + 64
	writer.Write4At(__AskSizeOffset, *(*uint32)(unsafe.Pointer(&x.AskSize)))
	__BidExchangeOffset := offset + 68
	writer.Write1At(__BidExchangeOffset, *(*uint8)(unsafe.Pointer(&x.BidExchange)))
	__AskExchangeOffset := offset + 69
	writer.Write1At(__AskExchangeOffset, *(*uint8)(unsafe.Pointer(&x.AskExchange)))
	__TapeOffset := offset + 70
	writer.Write1At(__TapeOffset, *(*uint8)(unsafe.Pointer(&x.Tape)))
	__NbboOffset := offset + 71
	writer.Write1At(__NbboOffset, *(*uint8)(unsafe.Pointer(&x.Nbbo)))

	return offset, nil
}

func (x *Quote) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewQuoteViewer(reader, 0), reader)
}

func (x *Quote) Read(viewer *QuoteViewer, reader *karmem.Reader) {
	x.Symbol = string(viewer.Symbol(reader))
	__ConditionsSlice := viewer.Conditions(reader)
	__ConditionsLen := len(__ConditionsSlice)
	if __ConditionsLen > cap(x.Conditions) {
		x.Conditions = append(x.Conditions, make([]byte, __ConditionsLen-len(x.Conditions))...)
	} else if __ConditionsLen > len(x.Conditions) {
		x.Conditions = x.Conditions[:__ConditionsLen]
	}
	copy(x.Conditions, __ConditionsSlice)
	for i := __ConditionsLen; i < len(x.Conditions); i++ {
		x.Conditions[i] = 0
	}
	x.Conditions = x.Conditions[:__ConditionsLen]
	x.Timestamp = viewer.Timestamp()
	x.ReceivedAt = viewer.ReceivedAt()
	x.BidPrice = viewer.BidPrice()
	x.AskPrice = viewer.AskPrice()
	x.BidSize = viewer.BidSize()
	x.AskSize = viewer.AskSize()
	x.BidExchange = viewer.BidExchange()
	x.AskExchange = viewer.AskExchange()
	x.Tape = viewer.Tape()
	x.Nbbo = viewer.Nbbo()
}

type Trade struct {
	Symbol     string
	Conditions []byte
	ID         uint64
	Timestamp  uint64
	ReceivedAt uint64
	Price      float64
	Volume     uint32
	Exchange   byte
	Tape       byte
}

func (x *Trade) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierTrade
}

func (x *Trade) Reset() {
	x.Symbol = x.Symbol[:0]
	x.Conditions = x.Conditions[:0]
	x.ID = 0
	x.Timestamp = 0
	x.ReceivedAt = 0
	x.Price = 0
	x.Volume = 0
	x.Exchange = 0
	x.Tape = 0
}

func (x *Trade) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *Trade) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(72)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(66))
	__SymbolSize := uint(1 * len(x.Symbol))
	__SymbolOffset, err := writer.Alloc(__SymbolSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+4, uint32(__SymbolOffset))
	writer.Write4At(offset+4+4, uint32(__SymbolSize))
	writer.Write4At(offset+4+4+4, 1)
	__SymbolSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Symbol)), __SymbolSize, __SymbolSize}
	writer.WriteAt(__SymbolOffset, *(*[]byte)(unsafe.Pointer(&__SymbolSlice)))
	__ConditionsSize := uint(1 * len(x.Conditions))
	__ConditionsOffset, err := writer.Alloc(__ConditionsSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+16, uint32(__ConditionsOffset))
	writer.Write4At(offset+16+4, uint32(__ConditionsSize))
	writer.Write4At(offset+16+4+4, 1)
	__ConditionsSlice := *(*[3]uint)(unsafe.Pointer(&x.Conditions))
	__ConditionsSlice[1] = __ConditionsSize
	__ConditionsSlice[2] = __ConditionsSize
	writer.WriteAt(__ConditionsOffset, *(*[]byte)(unsafe.Pointer(&__ConditionsSlice)))
	__IDOffset := offset + 28
	writer.Write8At(__IDOffset, *(*uint64)(unsafe.Pointer(&x.ID)))
	__TimestampOffset := offset + 36
	writer.Write8At(__TimestampOffset, *(*uint64)(unsafe.Pointer(&x.Timestamp)))
	__ReceivedAtOffset := offset + 44
	writer.Write8At(__ReceivedAtOffset, *(*uint64)(unsafe.Pointer(&x.ReceivedAt)))
	__PriceOffset := offset + 52
	writer.Write8At(__PriceOffset, *(*uint64)(unsafe.Pointer(&x.Price)))
	__VolumeOffset := offset + 60
	writer.Write4At(__VolumeOffset, *(*uint32)(unsafe.Pointer(&x.Volume)))
	__ExchangeOffset := offset + 64
	writer.Write1At(__ExchangeOffset, *(*uint8)(unsafe.Pointer(&x.Exchange)))
	__TapeOffset := offset + 65
	writer.Write1At(__TapeOffset, *(*uint8)(unsafe.Pointer(&x.Tape)))

	return offset, nil
}

func (x *Trade) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewTradeViewer(reader, 0), reader)
}

func (x *Trade) Read(viewer *TradeViewer, reader *karmem.Reader) {
	x.Symbol = string(viewer.Symbol(reader))
	__ConditionsSlice := viewer.Conditions(reader)
	__ConditionsLen := len(__ConditionsSlice)
	if __ConditionsLen > cap(x.Conditions) {
		x.Conditions = append(x.Conditions, make([]byte, __ConditionsLen-len(x.Conditions))...)
	} else if __ConditionsLen > len(x.Conditions) {
		x.Conditions = x.Conditions[:__ConditionsLen]
	}
	copy(x.Conditions, __ConditionsSlice)
	for i := __ConditionsLen; i < len(x.Conditions); i++ {
		x.Conditions[i] = 0
	}
	x.Conditions = x.Conditions[:__ConditionsLen]
	x.ID = viewer.ID()
	x.Timestamp = viewer.Timestamp()
	x.ReceivedAt = viewer.ReceivedAt()
	x.Price = viewer.Price()
	x.Volume = viewer.Volume()
	x.Exchange = viewer.Exchange()
	x.Tape = viewer.Tape()
}

type RawQuote struct {
	Data Quote
}

func (x *RawQuote) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierRawQuote
}

func (x *RawQuote) Reset() {
	x.Data.Reset()
}

func (x *RawQuote) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *RawQuote) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(8)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	__DataSize := uint(80)
	__DataOffset, err := writer.Alloc(__DataSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+0, uint32(__DataOffset))
	if _, err := x.Data.Write(writer, __DataOffset); err != nil {
		return offset, err
	}

	return offset, nil
}

func (x *RawQuote) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewRawQuoteViewer(reader, 0), reader)
}

func (x *RawQuote) Read(viewer *RawQuoteViewer, reader *karmem.Reader) {
	x.Data.Read(viewer.Data(reader), reader)
}

type RawTrade struct {
	Data Trade
}

func (x *RawTrade) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierRawTrade
}

func (x *RawTrade) Reset() {
	x.Data.Reset()
}

func (x *RawTrade) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *RawTrade) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(8)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	__DataSize := uint(72)
	__DataOffset, err := writer.Alloc(__DataSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+0, uint32(__DataOffset))
	if _, err := x.Data.Write(writer, __DataOffset); err != nil {
		return offset, err
	}

	return offset, nil
}

func (x *RawTrade) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewRawTradeViewer(reader, 0), reader)
}

func (x *RawTrade) Read(viewer *RawTradeViewer, reader *karmem.Reader) {
	x.Data.Read(viewer.Data(reader), reader)
}

type QuoteBatch struct {
	Quotes []RawQuote
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
	__QuotesSize := uint(8 * len(x.Quotes))
	__QuotesOffset, err := writer.Alloc(__QuotesSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+4, uint32(__QuotesOffset))
	writer.Write4At(offset+4+4, uint32(__QuotesSize))
	writer.Write4At(offset+4+4+4, 8)
	for i := range x.Quotes {
		if _, err := x.Quotes[i].Write(writer, __QuotesOffset); err != nil {
			return offset, err
		}
		__QuotesOffset += 8
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
		x.Quotes = append(x.Quotes, make([]RawQuote, __QuotesLen-len(x.Quotes))...)
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
	Trades []RawTrade
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
	__TradesSize := uint(8 * len(x.Trades))
	__TradesOffset, err := writer.Alloc(__TradesSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+4, uint32(__TradesOffset))
	writer.Write4At(offset+4+4, uint32(__TradesSize))
	writer.Write4At(offset+4+4+4, 8)
	for i := range x.Trades {
		if _, err := x.Trades[i].Write(writer, __TradesOffset); err != nil {
			return offset, err
		}
		__TradesOffset += 8
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
		x.Trades = append(x.Trades, make([]RawTrade, __TradesLen-len(x.Trades))...)
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
	_data [80]byte
}

var _NullQuoteViewer = QuoteViewer{}

func NewQuoteViewer(reader *karmem.Reader, offset uint32) (v *QuoteViewer) {
	if !reader.IsValidOffset(offset, 8) {
		return &_NullQuoteViewer
	}
	v = (*QuoteViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return &_NullQuoteViewer
	}
	return v
}

func (x *QuoteViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(&x._data))
}
func (x *QuoteViewer) Symbol(reader *karmem.Reader) (v []byte) {
	if 4+12 > x.size() {
		return []byte{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4+4))
	if !reader.IsValidOffset(offset, size) {
		return []byte{}
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]byte)(unsafe.Pointer(&slice))
}
func (x *QuoteViewer) Conditions(reader *karmem.Reader) (v []byte) {
	if 16+12 > x.size() {
		return []byte{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 16))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 16+4))
	if !reader.IsValidOffset(offset, size) {
		return []byte{}
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]byte)(unsafe.Pointer(&slice))
}
func (x *QuoteViewer) Timestamp() (v uint64) {
	if 28+8 > x.size() {
		return v
	}
	return *(*uint64)(unsafe.Add(unsafe.Pointer(&x._data), 28))
}
func (x *QuoteViewer) ReceivedAt() (v uint64) {
	if 36+8 > x.size() {
		return v
	}
	return *(*uint64)(unsafe.Add(unsafe.Pointer(&x._data), 36))
}
func (x *QuoteViewer) BidPrice() (v float64) {
	if 44+8 > x.size() {
		return v
	}
	return *(*float64)(unsafe.Add(unsafe.Pointer(&x._data), 44))
}
func (x *QuoteViewer) AskPrice() (v float64) {
	if 52+8 > x.size() {
		return v
	}
	return *(*float64)(unsafe.Add(unsafe.Pointer(&x._data), 52))
}
func (x *QuoteViewer) BidSize() (v uint32) {
	if 60+4 > x.size() {
		return v
	}
	return *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 60))
}
func (x *QuoteViewer) AskSize() (v uint32) {
	if 64+4 > x.size() {
		return v
	}
	return *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 64))
}
func (x *QuoteViewer) BidExchange() (v byte) {
	if 68+1 > x.size() {
		return v
	}
	return *(*byte)(unsafe.Add(unsafe.Pointer(&x._data), 68))
}
func (x *QuoteViewer) AskExchange() (v byte) {
	if 69+1 > x.size() {
		return v
	}
	return *(*byte)(unsafe.Add(unsafe.Pointer(&x._data), 69))
}
func (x *QuoteViewer) Tape() (v byte) {
	if 70+1 > x.size() {
		return v
	}
	return *(*byte)(unsafe.Add(unsafe.Pointer(&x._data), 70))
}
func (x *QuoteViewer) Nbbo() (v bool) {
	if 71+1 > x.size() {
		return v
	}
	return *(*bool)(unsafe.Add(unsafe.Pointer(&x._data), 71))
}

type TradeViewer struct {
	_data [72]byte
}

var _NullTradeViewer = TradeViewer{}

func NewTradeViewer(reader *karmem.Reader, offset uint32) (v *TradeViewer) {
	if !reader.IsValidOffset(offset, 8) {
		return &_NullTradeViewer
	}
	v = (*TradeViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return &_NullTradeViewer
	}
	return v
}

func (x *TradeViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(&x._data))
}
func (x *TradeViewer) Symbol(reader *karmem.Reader) (v []byte) {
	if 4+12 > x.size() {
		return []byte{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4+4))
	if !reader.IsValidOffset(offset, size) {
		return []byte{}
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]byte)(unsafe.Pointer(&slice))
}
func (x *TradeViewer) Conditions(reader *karmem.Reader) (v []byte) {
	if 16+12 > x.size() {
		return []byte{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 16))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 16+4))
	if !reader.IsValidOffset(offset, size) {
		return []byte{}
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]byte)(unsafe.Pointer(&slice))
}
func (x *TradeViewer) ID() (v uint64) {
	if 28+8 > x.size() {
		return v
	}
	return *(*uint64)(unsafe.Add(unsafe.Pointer(&x._data), 28))
}
func (x *TradeViewer) Timestamp() (v uint64) {
	if 36+8 > x.size() {
		return v
	}
	return *(*uint64)(unsafe.Add(unsafe.Pointer(&x._data), 36))
}
func (x *TradeViewer) ReceivedAt() (v uint64) {
	if 44+8 > x.size() {
		return v
	}
	return *(*uint64)(unsafe.Add(unsafe.Pointer(&x._data), 44))
}
func (x *TradeViewer) Price() (v float64) {
	if 52+8 > x.size() {
		return v
	}
	return *(*float64)(unsafe.Add(unsafe.Pointer(&x._data), 52))
}
func (x *TradeViewer) Volume() (v uint32) {
	if 60+4 > x.size() {
		return v
	}
	return *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 60))
}
func (x *TradeViewer) Exchange() (v byte) {
	if 64+1 > x.size() {
		return v
	}
	return *(*byte)(unsafe.Add(unsafe.Pointer(&x._data), 64))
}
func (x *TradeViewer) Tape() (v byte) {
	if 65+1 > x.size() {
		return v
	}
	return *(*byte)(unsafe.Add(unsafe.Pointer(&x._data), 65))
}

type RawQuoteViewer struct {
	_data [8]byte
}

var _NullRawQuoteViewer = RawQuoteViewer{}

func NewRawQuoteViewer(reader *karmem.Reader, offset uint32) (v *RawQuoteViewer) {
	if !reader.IsValidOffset(offset, 8) {
		return &_NullRawQuoteViewer
	}
	v = (*RawQuoteViewer)(unsafe.Add(reader.Pointer, offset))
	return v
}

func (x *RawQuoteViewer) size() uint32 {
	return 8
}
func (x *RawQuoteViewer) Data(reader *karmem.Reader) (v *QuoteViewer) {
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 0))
	if !reader.IsValidOffset(offset, 80) {
		return &_NullQuoteViewer
	}
	v = (*QuoteViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return &_NullQuoteViewer
	}
	return v
}

type RawTradeViewer struct {
	_data [8]byte
}

var _NullRawTradeViewer = RawTradeViewer{}

func NewRawTradeViewer(reader *karmem.Reader, offset uint32) (v *RawTradeViewer) {
	if !reader.IsValidOffset(offset, 8) {
		return &_NullRawTradeViewer
	}
	v = (*RawTradeViewer)(unsafe.Add(reader.Pointer, offset))
	return v
}

func (x *RawTradeViewer) size() uint32 {
	return 8
}
func (x *RawTradeViewer) Data(reader *karmem.Reader) (v *TradeViewer) {
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 0))
	if !reader.IsValidOffset(offset, 72) {
		return &_NullTradeViewer
	}
	v = (*TradeViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return &_NullTradeViewer
	}
	return v
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
func (x *QuoteBatchViewer) Quotes(reader *karmem.Reader) (v []RawQuoteViewer) {
	if 4+12 > x.size() {
		return []RawQuoteViewer{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4+4))
	if !reader.IsValidOffset(offset, size) {
		return []RawQuoteViewer{}
	}
	length := uintptr(size / 8)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]RawQuoteViewer)(unsafe.Pointer(&slice))
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
func (x *TradeBatchViewer) Trades(reader *karmem.Reader) (v []RawTradeViewer) {
	if 4+12 > x.size() {
		return []RawTradeViewer{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4+4))
	if !reader.IsValidOffset(offset, size) {
		return []RawTradeViewer{}
	}
	length := uintptr(size / 8)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]RawTradeViewer)(unsafe.Pointer(&slice))
}
