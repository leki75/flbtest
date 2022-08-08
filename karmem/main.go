package main

import (
	"fmt"
	"time"
	"unsafe"

	karmemgo "karmem.org/golang"

	"leki75/flbtest/schema/karmem"
)

var now = uint64(time.Now().UnixNano())

func main() {
	fmt.Println("KarmemQuoteBatch:", len(quoteBatch()), unsafe.Sizeof(karmem.Quote{}))
	fmt.Println("KarmemTradeBatch:", len(tradeBatch()), unsafe.Sizeof(karmem.Trade{}))
}

func quoteBatch() []byte {
	quote1 := karmem.Quote{
		Symbol:      [11]byte{'A', 'A', 'P', 'L'},
		Conditions:  [2]byte{'C', 'D'},
		Timestamp:   now,
		ReceivedAt:  now,
		BidPrice:    100.0,
		AskPrice:    200.0,
		BidSize:     100,
		AskSize:     200,
		BidExchange: 'A',
		AskExchange: 'B',
		Tape:        'E',
		Nbbo:        true,
	}

	quote2 := karmem.Quote{
		Symbol:      [11]byte{'T', 'S', 'L', 'A'},
		Conditions:  [2]byte{'H', 'I'},
		Timestamp:   now,
		ReceivedAt:  now,
		BidPrice:    300.0,
		AskPrice:    400.0,
		BidSize:     300,
		AskSize:     400,
		BidExchange: 'F',
		AskExchange: 'B',
		Tape:        'J',
		Nbbo:        true,
	}

	batch := karmem.QuoteBatch{
		Quotes: []karmem.Quote{
			quote1,
			quote2,
		},
	}

	writer := karmemgo.NewWriter(256)
	if _, err := batch.WriteAsRoot(writer); err != nil {
		panic(err)
	}

	return writer.Bytes()
}

func tradeBatch() []byte {
	trade1 := karmem.Trade{
		Symbol:     [11]byte{'A', 'A', 'P', 'L'},
		Conditions: [4]byte{'A', 'B', 'C', 'D'},
		ID:         12345678,
		Timestamp:  now,
		ReceivedAt: now,
		Price:      100.0,
		Volume:     100,
		Exchange:   'E',
		Tape:       'F',
	}

	trade2 := karmem.Trade{
		Symbol:     [11]byte{'T', 'S', 'L', 'A'},
		Conditions: [4]byte{'G', 'H', 'I', 'J'},
		ID:         87654321,
		Timestamp:  now,
		ReceivedAt: now,
		Price:      200.0,
		Volume:     200,
		Exchange:   'K',
		Tape:       'L',
	}

	batch := karmem.TradeBatch{
		Trades: []karmem.Trade{
			trade1,
			trade2,
		},
	}

	writer := karmemgo.NewWriter(256)
	if _, err := batch.WriteAsRoot(writer); err != nil {
		panic(err)
	}

	return writer.Bytes()
}
