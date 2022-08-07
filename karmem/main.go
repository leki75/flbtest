package main

import (
	"fmt"
	"time"

	karmemgo "karmem.org/golang"

	"leki75/flbtest/schema/karmem"
)

var now = uint64(time.Now().UnixNano())

func main() {
	fmt.Println("KarmemQuoteBatch:", len(quoteBatch()))
	fmt.Println("KarmemTradeBatch:", len(tradeBatch()))
}

func quoteBatch() []byte {
	quote1 := karmem.Quote{
		Symbol:      "AAPL",
		Conditions:  []byte{'C', 'D'},
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
		Symbol:      "TSLA",
		Conditions:  []byte{'H', 'I'},
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
		Quotes: []karmem.RawQuote{
			{Data: quote1},
			{Data: quote2},
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
		Symbol:     "AAPL",
		Conditions: []byte{'A', 'B', 'C', 'D'},
		ID:         12345678,
		Timestamp:  now,
		ReceivedAt: now,
		Price:      100.0,
		Volume:     100,
		Exchange:   'E',
		Tape:       'F',
	}

	trade2 := karmem.Trade{
		Symbol:     "AAPL",
		Conditions: []byte{'G', 'H', 'I', 'J'},
		ID:         87654321,
		Timestamp:  now,
		ReceivedAt: now,
		Price:      200.0,
		Volume:     200,
		Exchange:   'K',
		Tape:       'L',
	}

	batch := karmem.TradeBatch{
		Trades: []karmem.RawTrade{
			{Data: trade1},
			{Data: trade2},
		},
	}

	writer := karmemgo.NewWriter(256)
	if _, err := batch.WriteAsRoot(writer); err != nil {
		panic(err)
	}

	return writer.Bytes()
}
