package main

import (
	"fmt"
	"time"
	"unsafe"

	"leki75/flbtest/schema/proto"
)

var now = uint64(time.Now().UnixNano())

func main() {
	fmt.Println("ProtoQuoteBatch:", len(quoteBatch()), unsafe.Sizeof(proto.Quote{}))
	fmt.Println("ProtoTradeBatch:", len(tradeBatch()), unsafe.Sizeof(proto.Trade{}))
}

func quoteBatch() []byte {
	quote1 := proto.Quote{
		Symbol:      "AAPL",
		Conditions:  []byte{'C', 'D'},
		Timestamp:   now,
		ReceivedAt:  int64(now),
		BidPrice:    100.0,
		AskPrice:    200.0,
		BidSize:     100,
		AskSize:     200,
		BidExchange: 'A',
		AskExchange: 'B',
		Tape:        'E',
		Nbbo:        true,
	}

	quote2 := proto.Quote{
		Symbol:      "TSLA",
		Conditions:  []byte{'H', 'I'},
		Timestamp:   now,
		ReceivedAt:  int64(now),
		BidPrice:    300.0,
		AskPrice:    400.0,
		BidSize:     300,
		AskSize:     400,
		BidExchange: 'F',
		AskExchange: 'B',
		Tape:        'J',
		Nbbo:        true,
	}

	batch := proto.QuoteBatch{
		Quotes: []*proto.Quote{
			&quote1,
			&quote2,
		},
	}

	data, err := batch.MarshalVT()
	if err != nil {
		panic(err)
	}

	return data
}

func tradeBatch() []byte {
	trade1 := proto.Trade{
		Symbol:     "AAPL",
		Conditions: []byte{'A', 'B', 'C', 'D'},
		Id:         12345678,
		Timestamp:  now,
		ReceivedAt: int64(now),
		Price:      100.0,
		Volume:     100,
		Exchange:   'E',
		Tape:       'F',
	}

	trade2 := proto.Trade{
		Symbol:     "TSLA",
		Conditions: []byte{'G', 'H', 'I', 'J'},
		Id:         87654321,
		Timestamp:  now,
		ReceivedAt: int64(now),
		Price:      200.0,
		Volume:     200,
		Exchange:   'K',
		Tape:       'L',
	}

	batch := proto.TradeBatch{
		Trades: []*proto.Trade{
			&trade1,
			&trade2,
		},
	}

	data, err := batch.MarshalVT()
	if err != nil {
		panic(err)
	}

	return data
}
